package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	inputName   = "redis"
	catalogName = "db"
	l           = logger.DefaultSLogger("redis")
)

type Input struct {
	Host              string               `toml:"host"`
	Port              int                  `toml:"port"`
	UnixSocketPath    string               `toml:"unix_socket_path"`
	DB                int                  `toml:"db"`
	Password          string               `toml:"password"`
	Service           string               `toml:"service"`
	SocketTimeout     int                  `toml:"socket_timeout"`
	Interval          string               `toml:"interval"`
	IntervalDuration  time.Duration        `toml:"-"`
	Keys              []string             `toml:"keys"`
	WarnOnMissingKeys bool                 `toml:"warn_on_missing_keys"`
	CommandStats      bool                 `toml:"command_stats"`
	Slowlog           bool                 `toml:"slow_log"`
	SlowlogMaxLen     int                  `toml:"slowlog-max-len"`
	Tags              map[string]string    `toml:"tags"`
	client            *redis.Client        `toml:"-"`
	collectCache      []inputs.Measurement `toml:"-"`
	Addr              string               `toml:"-"`
}

func (i *Input) initCfg() {
	// 采集频度
	i.IntervalDuration = 1 * time.Minute

	if i.Interval != "" {
		du, err := time.ParseDuration(i.Interval)
		if err != nil {
			l.Errorf("bad interval %s: %s, use default: 1m", i.Interval, err.Error())
		} else {
			i.IntervalDuration = du
		}
	}

	i.Addr = fmt.Sprintf("%s:%d", i.Host, i.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     i.Addr,
		Password: i.Password, // no password set
		DB:       i.DB,       // use default DB
	})

	i.client = client

	i.globalTag()
}

func (i *Input) globalTag() {
	i.Tags["server"] = i.Addr
	i.Tags["service_name"] = i.Service
}

func (i *Input) Collect() error {
	i.collectCache = []inputs.Measurement{}

	i.collectInfoMeasurement()

	// 获取客户端信息
	i.collectClientMeasurement()

	// db command
	if i.CommandStats {
		i.collectCommandMeasurement()
	}

	// slowlog
	if i.Slowlog {
		slowlogMeasurement := CollectSlowlogMeasurement(i)
		i.collectCache = append(i.collectCache, slowlogMeasurement)
	}

	// bigkey
	if len(i.Keys) > 0 {
		bigKeyMeasurement := CollectBigKeyMeasurement(i)
		i.collectCache = append(i.collectCache, bigKeyMeasurement)
	}

	return nil
}

func (i *Input) collectInfoMeasurement() {
	m := &infoMeasurement{
		client:  i.client,
		resData: make(map[string]interface{}),
		tags:    make(map[string]string),
		fields:  make(map[string]interface{}),
	}

	m.name = "redis_info"
	for key, value := range i.Tags {
		m.tags[key] = value
	}

	m.getData()
	m.submit()

	i.collectCache = append(i.collectCache, m)
}

func (i *Input) collectClientMeasurement() {
	i.getClientData()
}

func (i *Input) collectCommandMeasurement() {
	i.getCommandData()
}

func (i *Input) Run() {
	l = logger.SLogger("redis")
	tick := time.NewTicker(i.IntervalDuration)
	defer tick.Stop()

	i.initCfg()

	n := 0

	for {
		n++
		select {
		case <-tick.C:
			l.Debugf("redis input gathering...")
			start := time.Now()
			if err := i.Collect(); err != nil {
				l.Error(err)
			} else {
				inputs.FeedMeasurement(inputName, io.Metric, i.collectCache,
					&io.Option{CollectCost: time.Since(start), HighFreq: (n%2 == 0)})

				i.collectCache = i.collectCache[:] // NOTE: do not forget to clean cache
			}

		case <-datakit.Exit.Wait():
			return
		}
	}
}

func (i *Input) Catalog() string { return catalogName }

func (i *Input) SampleConfig() string { return configSample }

func (i *Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&infoMeasurement{},
		&clientMeasurement{},
		&commandMeasurement{},
		&slowlogMeasurement{},
		&bigKeyMeasurement{},
	}
}

func init() {
	inputs.Add(inputName, func() inputs.Input {
		return &Input{}
	})
}
