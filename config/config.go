package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/influxdata/toml"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
)

var (
	Cfg           Config
	CfgPath       string
	ExecutableDir string

	ErrNoTelegrafConf = errors.New("no telegraf config")
)

const (
	telegrafConfTemplate = `
[agent]
  interval = "10s"
  round_interval = true

  metric_batch_size = 1000
  metric_buffer_limit = 10000
  collection_jitter = "0s"
  flush_interval = "10s"
  flush_jitter = "0s"
  precision = ""
  logfile="{{.LogFile}}"
  debug = false
  quiet = false
  hostname = ""
  omit_hostname = false

[[outputs.http]]
  url = "{{.FtGateway}}"
  method = "POST"
  data_format = "influx"
  content_encoding = "gzip"

  ## Additional HTTP headers
  [outputs.http.headers]
    ## Should be set manually to "application/json" for json data_format
	X-Datakit-UUID = "{{.DKUUID}}"
	X-Datakit-Version = "{{.DKVERSION}}"
	User-Agent = "forethought datakit"
`
)

var (
	DKVersion = "datakit-v" + git.Version
)

type Config struct {
	UUID      string `toml:"uuid"`
	FtGateway string `toml:"ftdataway"`

	Log      string `toml:"log"`
	LogLevel string `toml:"log_level"`

	ConfigDir string `toml:"config_dir,omitempty"`
}

func LoadConfig(f string) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	if err := toml.Unmarshal(data, &Cfg); err != nil {
		return err
	}

	return nil
}

type Configuration interface {
	SampleConfig() string
	FilePath(string) string
	ToTelegraf() (string, error)
	Load(string) error
}

var SubConfigs = map[string]Configuration{}

func AddConfig(name string, c Configuration) {
	SubConfigs[name] = c
}

func InitializeConfigs() error {

	out, err := toml.Marshal(&Cfg)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(CfgPath, out, 0664); err != nil {
		return err
	}

	for _, c := range SubConfigs {
		sample := c.SampleConfig()
		f := c.FilePath(Cfg.ConfigDir)
		os.MkdirAll(filepath.Dir(f), 0775)
		if err := ioutil.WriteFile(f, []byte(sample), 0644); err != nil {
			return err
		}
	}
	return nil
}

func LoadSubConfigs(root string) error {

	for _, c := range SubConfigs {
		f := c.FilePath(root)
		_, err := os.Stat(f)
		if err != nil && os.IsNotExist(err) {
			continue
		}
		if err := c.Load(f); err != nil {
			return fmt.Errorf("load config \"%s\" failed: %s", f, err.Error())
		}
	}

	return nil
}

func GenerateTelegrafConfig() (string, error) {

	type AgentCfg struct {
		LogFile   string
		FtGateway string
		DKUUID    string
		DKVERSION string
	}

	agentcfg := AgentCfg{
		LogFile:   filepath.Join(ExecutableDir, "agent.log"),
		FtGateway: Cfg.FtGateway,
		DKUUID:    Cfg.UUID,
		DKVERSION: DKVersion,
	}

	var err error
	tm := template.New("")
	tm, err = tm.Parse(telegrafConfTemplate)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})
	if err = tm.Execute(buf, &agentcfg); err != nil {
		return "", err
	}

	cfg := string(buf.Bytes())

	telcfgs := ""

	for _, c := range SubConfigs {
		telcfg, err := c.ToTelegraf()
		if err != nil {
			return "", err
		}
		telcfgs += telcfg
	}

	if telcfgs == "" {
		return "", ErrNoTelegrafConf
	}

	cfg += telcfgs

	return cfg, err
}
