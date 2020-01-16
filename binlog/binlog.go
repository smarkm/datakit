package binlog

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
	//"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/log"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/service"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/uploader"
)

func init() {
	config.AddConfig("binlog", &Cfg)

	service.Add("binlog", func(logger log.Logger) service.Service {
		if len(Cfg.Datasources) == 0 {
			return nil
		}
		return &BinlogSvr{
			logger: logger,
		}
	})
}

type Binloger struct {
	m sync.Mutex

	cfg *BinlogDatasource

	parser *parser.Parser
	master *masterInfo

	syncer *replication.BinlogSyncer

	eventHandler EventHandler

	connLock sync.Mutex
	conn     *client.Conn

	tableLock          sync.RWMutex
	tables             map[string]*schema.Table
	errorTablesGetTime map[string]time.Time

	tableMatchCache map[string]bool
	// includeTableRegex []*regexp.Regexp
	// excludeTableRegex []*regexp.Regexp

	delay *uint32

	//ctx    context.Context
	//cancel context.CancelFunc

	logger log.Logger

	storage uploader.IUploader
}

type BinlogSvr struct {
	binlogs []*Binloger
	logger  log.Logger
}

var (
	UnknownTableRetryPeriod = time.Second * time.Duration(10)
	ErrExcludedTable        = errors.New("excluded table meta")

	HeartbeatPeriod = 60 * time.Second
	ReadTimeout     = 90 * time.Second

	binlogs []*Binloger
)

func (s *BinlogSvr) Start(ctx context.Context, up uploader.IUploader) error {

	if len(Cfg.Datasources) == 0 {
		return nil
	}

	s.logger.Info("Starting Binlog...")

	var wg sync.WaitGroup

	s.binlogs = []*Binloger{}

	for _, dt := range Cfg.Datasources {
		dt.ServerID = uint32(rand.New(rand.NewSource(time.Now().Unix())).Intn(1000)) + 1001
		dt.Charset = mysql.DEFAULT_CHARSET
		dt.Flavor = mysql.MySQLFlavor
		dt.HeartbeatPeriod = HeartbeatPeriod
		dt.DiscardNoMetaRowEvent = true
		dt.ReadTimeout = ReadTimeout
		dt.UseDecimal = true
		dt.ParseTime = true
		dt.SemiSyncEnabled = false

		binloger := NewBinloger(dt, s.logger)
		binloger.storage = up

		s.binlogs = append(s.binlogs, binloger)

		wg.Add(1)
		go func(b *Binloger) {
			defer wg.Done()
			if err := b.Run(ctx); err != nil && err != context.Canceled {
				s.logger.Errorf("%s", err.Error())
			}
		}(binloger)

	}

	wg.Wait()

	s.logger.Info("Binlog done")

	return nil
}

func (b *Binloger) Run(ctx context.Context) error {

	b.tables = make(map[string]*schema.Table)
	if b.cfg.DiscardNoMetaRowEvent {
		b.errorTablesGetTime = make(map[string]time.Time)
	}
	b.master = &masterInfo{}

	// if c.includeTableRegex != nil || c.excludeTableRegex != nil {
	b.tableMatchCache = make(map[string]bool)
	// }

	var err error

	if err = b.prepareSyncer(); err != nil {
		return err
	}

	if err = b.checkMysqlVersion(); err != nil {
		return err
	}

	if err := b.checkBinlogRowFormat(); err != nil {
		return err
	}

	if err := b.CheckBinlogRowImage("FULL"); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return context.Canceled
	default:
	}

	b.logger.Info("check requirments ok")

	b.master.UpdateTimestamp(uint32(time.Now().Unix()))

	if err = b.getMasterStatus(b.master); err != nil {
		return err
	}

	if err := b.runSyncBinlog(ctx); err != nil {
		if err != context.Canceled && errors.Cause(err) != context.Canceled {
			b.logger.Errorf("start sync binlog err: %v", err)
			return err
		}
	}

	return nil
}

// func (b *Binloger) Stop() {
// 	b.m.Lock()
// 	defer b.m.Unlock()

// 	b.syncer.Close()
// 	b.connLock.Lock()
// 	if b.conn != nil {
// 		b.conn.Close()
// 		b.conn = nil
// 	}
// 	b.connLock.Unlock()

// 	b.eventHandler.OnPosSynced(b.master.Position(), b.master.GTIDSet(), true)
// }

func NewBinloger(cfg *BinlogDatasource, logger log.Logger) *Binloger {
	b := &Binloger{
		cfg:    cfg,
		logger: logger,
		parser: parser.New(),
		delay:  new(uint32),
	}

	b.eventHandler = &MainEventHandler{
		binloger: b,
	}

	return b
}

func (b *Binloger) prepareSyncer() error {
	cfg := replication.BinlogSyncerConfig{
		ServerID:                b.cfg.ServerID,
		Flavor:                  "mysql",
		User:                    b.cfg.User,
		Password:                b.cfg.Password,
		Charset:                 b.cfg.Charset,
		HeartbeatPeriod:         b.cfg.HeartbeatPeriod,
		ReadTimeout:             b.cfg.ReadTimeout,
		UseDecimal:              b.cfg.UseDecimal,
		ParseTime:               b.cfg.ParseTime,
		SemiSyncEnabled:         b.cfg.SemiSyncEnabled,
		MaxReconnectAttempts:    b.cfg.MaxReconnectAttempts,
		TimestampStringLocation: b.cfg.TimestampStringLocation,
	}

	if strings.Contains(b.cfg.Addr, "/") {
		cfg.Host = b.cfg.Addr
	} else {
		seps := strings.Split(b.cfg.Addr, ":")
		if len(seps) != 2 {
			return errors.Errorf("invalid mysql addr format %s, must host:port", b.cfg.Addr)
		}

		port, err := strconv.ParseUint(seps[1], 10, 16)
		if err != nil {
			return errors.Trace(err)
		}

		cfg.Host = seps[0]
		cfg.Port = uint16(port)
	}

	b.syncer = replication.NewBinlogSyncer(cfg)

	return nil
}

func (c *Binloger) GetTable(db string, table string) (*schema.Table, *BinlogDatabase, error) {
	key := fmt.Sprintf("%s.%s", db, table)
	// if table is excluded, return error and skip parsing event or dump
	//fmt.Println(key)
	target := c.checkTableMatch(db, table)
	if target == nil {
		return nil, nil, ErrExcludedTable
	}

	c.tableLock.RLock()
	t, ok := c.tables[key]
	c.tableLock.RUnlock()

	if ok {
		return t, target, nil
	}

	if c.cfg.DiscardNoMetaRowEvent {
		c.tableLock.RLock()
		lastTime, ok := c.errorTablesGetTime[key]
		c.tableLock.RUnlock()
		if ok && time.Now().Sub(lastTime) < UnknownTableRetryPeriod {
			return nil, nil, schema.ErrMissingTableMeta
		}
	}

	t, err := schema.NewTable(c, db, table)
	if err != nil {
		// check table not exists
		if ok, err1 := schema.IsTableExist(c, db, table); err1 == nil && !ok {
			return nil, nil, schema.ErrTableNotExist
		}
		// work around : RDS HAHeartBeat
		// ref : https://github.com/alibaba/canal/blob/master/parse/src/main/java/com/alibaba/otter/canal/parse/inbound/mysql/dbsync/LogEventConvert.java#L385
		// issue : https://github.com/alibaba/canal/issues/222
		// This is a common error in RDS that canal can't get HAHealthCheckSchema's meta, so we mock a table meta.
		// If canal just skip and log error, as RDS HA heartbeat interval is very short, so too many HAHeartBeat errors will be logged.
		if key == schema.HAHealthCheckSchema {
			// mock ha_health_check meta
			ta := &schema.Table{
				Schema:  db,
				Name:    table,
				Columns: make([]schema.TableColumn, 0, 2),
				Indexes: make([]*schema.Index, 0),
			}
			ta.AddColumn("id", "bigint(20)", "", "")
			ta.AddColumn("type", "char(1)", "", "")
			c.tableLock.Lock()
			c.tables[key] = ta
			c.tableLock.Unlock()
			return ta, target, nil
		}
		// if DiscardNoMetaRowEvent is true, we just log this error
		if c.cfg.DiscardNoMetaRowEvent {
			c.tableLock.Lock()
			c.errorTablesGetTime[key] = time.Now()
			c.tableLock.Unlock()
			// log error and return ErrMissingTableMeta
			c.logger.Errorf("get table meta err: %v", errors.Trace(err))
			return nil, nil, schema.ErrMissingTableMeta
		}
		return nil, nil, err
	}

	c.tableLock.Lock()
	c.tables[key] = t
	if c.cfg.DiscardNoMetaRowEvent {
		// if get table info success, delete this key from errorTablesGetTime
		delete(c.errorTablesGetTime, key)
	}
	c.tableLock.Unlock()

	return t, target, nil
}

// ClearTableCache clear table cache
func (c *Binloger) ClearTableCache(db []byte, table []byte) {
	key := fmt.Sprintf("%s.%s", db, table)
	c.tableLock.Lock()
	delete(c.tables, key)
	if c.cfg.DiscardNoMetaRowEvent {
		delete(c.errorTablesGetTime, key)
	}
	c.tableLock.Unlock()
}

// CheckBinlogRowImage checks MySQL binlog row image, must be in FULL, MINIMAL, NOBLOB
func (c *Binloger) CheckBinlogRowImage(image string) error {
	// need to check MySQL binlog row image? full, minimal or noblob?
	// now only log
	if c.cfg.Flavor == mysql.MySQLFlavor {
		if res, err := c.Execute(`SHOW GLOBAL VARIABLES LIKE "binlog_row_image"`); err != nil {
			return errors.Trace(err)
		} else {
			// MySQL has binlog row image from 5.6, so older will return empty
			rowImage, _ := res.GetString(0, 1)
			if rowImage != "" && !strings.EqualFold(rowImage, image) {
				return errors.Errorf("MySQL uses %s binlog row image, but we want %s", rowImage, image)
			}
		}
	}

	return nil
}

func (c *Binloger) checkBinlogRowFormat() error {
	res, err := c.Execute(`SHOW GLOBAL VARIABLES LIKE "binlog_format";`)
	if err != nil {
		return errors.Trace(err)
	} else if f, _ := res.GetString(0, 1); f != "ROW" {
		return errors.Errorf("binlog must ROW format, but %s now", f)
	}

	return nil
}

func (c *Binloger) getMasterStatus(m *masterInfo) error {
	res, err := c.Execute(`show master status;`)
	if err != nil {
		return errors.Trace(err)
	}

	filename, err := res.GetString(0, 0)
	if err != nil {
		return errors.Trace(err)
	}
	pos, err := res.GetUint(0, 1)
	if err != nil {
		return errors.Trace(err)
	}

	m.Update(mysql.Position{
		Name: filename,
		Pos:  uint32(pos),
	})

	dodb, err := res.GetString(0, 2)
	if err == nil {
		m.binlogDoDB = dodb
	}

	ignoredb, err := res.GetString(0, 3)
	if err == nil {
		m.binlogIngoreDB = ignoredb
	}

	return nil
}

func (b *Binloger) checkMysqlVersion() error {
	es, err := b.Execute(`SELECT version();`)
	if err != nil {
		return errors.Trace(err)
	}

	ver, _ := es.GetString(0, 0)
	b.logger.Infof("server version: %s", ver)
	if strings.Contains(strings.ToLower(ver), "maria") {
		b.cfg.Flavor = "mariadb"
	}

	return nil
}

// Execute a SQL
func (c *Binloger) Execute(cmd string, args ...interface{}) (rr *mysql.Result, err error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	retryNum := 3
	for i := 0; i < retryNum; i++ {
		if c.conn == nil {
			c.logger.Infof("user is %s", c.cfg.User)
			c.conn, err = client.Connect(c.cfg.Addr, c.cfg.User, c.cfg.Password, "")
			if err != nil {
				return nil, errors.Trace(err)
			}
		}

		rr, err = c.conn.Execute(cmd, args...)
		if err != nil && !mysql.ErrorEqual(err, mysql.ErrBadConn) {
			return
		} else if mysql.ErrorEqual(err, mysql.ErrBadConn) {
			c.conn.Close()
			c.conn = nil
			continue
		} else {
			return
		}
	}
	return
}

func (c *Binloger) checkTableMatch(db, table string) *BinlogDatabase {

	var destdb *BinlogDatabase
	for _, t := range c.cfg.Databases {
		if t.Database == db {
			destdb = t
			break
		}
	}

	if destdb == nil {
		return nil
	}

	bmatch := false

	if len(destdb.Tables) > 0 {
		for _, tbl := range destdb.Tables {
			//fmt.Printf("tbname: %s\n", tbl.Name)
			if tbl.Name == table {
				bmatch = true
				break
			}
		}
	} else {
		bmatch = true
	}

	// if len(destdb.ExcludeTables) > 0 {
	// 	for _, tblname := range destdb.ExcludeTables {
	// 		if table == tblname {
	// 			bmatch = false
	// 			break
	// 		}
	// 	}
	// }

	if !bmatch {
		return nil
	}

	return destdb
}
