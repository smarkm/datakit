package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/man"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type enabledInput struct {
	Input     string   `json:"input"`
	Instances int      `json:"instances"`
	Cfgs      []string `json:"configs"`
	Panics    int      `json:"panic"`
}

type datakitStats struct {
	InputsStats     map[string]*io.InputsStat `json:"inputs_status"`
	EnabledInputs   []*enabledInput           `json:"enabled_inputs"`
	AvailableInputs []string                  `json:"available_inputs"`

	Version      string    `json:"version"`
	BuildAt      string    `json:"build_at"`
	Branch       string    `json:"branch"`
	Uptime       string    `json:"uptime"`
	OSArch       string    `json:"os_arch"`
	Reload       time.Time `json:"reload,omitempty"`
	ReloadCnt    int       `json:"reload_cnt"`
	WithinDocker bool      `json:"docker"`
	IOChanStat   string    `json:"io_chan_stats"`

	CSS string `json:"-"`
}

const (
	monitorTmpl = `
{{.CSS}}

# DataKit 运行展示

## 基本信息

- 版本         : {{.Version}}
- 运行时间     : {{.Uptime}}
- 发布日期     : {{.BuildAt}}
- 分支         : {{.Branch}}
- 系统类型     : {{.OSArch}}
- 是否容器运行 : {{.WithinDocker}}
- Reload 情况  : {{.ReloadCnt}} 次/{{.Reload}}
- IO 消耗统计  : {{.IOChanStat}}

## 采集器运行情况

{{.InputsStatsTable}}
`
)

func (x *datakitStats) InputsStatsTable() string {
	/*
			"category": "metric",
			"frequency": "20.49/min",
			"avg_size": 1,
			"total": 29,
			"count": 28,
			"first": "2021-05-07T19:41:22.943257+08:00",
			"last": "2021-05-07T19:42:44.252574+08:00",
			"last_error": "mocked error from demo input",
			"last_error_ts": "2021-05-07T19:42:43.923569+08:00",
			"max_collect_cost": 1001554420,
			"avg_collect_cost": 1000723212

		Category  string    `json:"category"`
		Frequency string    `json:"frequency,omitempty"`
		AvgSize   int64     `json:"avg_size"`
		Total     int64     `json:"total"`
		Count     int64     `json:"count"`
		First     time.Time `json:"first"`
		Last      time.Time `json:"last"`

		LastErr   string    `json:"last_error,omitempty"`
		LastErrTS time.Time `json:"last_error_ts,omitempty"`

		MaxCollectCost time.Duration `json:"max_collect_cost"`
		AvgCollectCost time.Duration `json:"avg_collect_cost"`
	*/

	const (
		tblHeader = `
| 采集器 | 实例个数 | 数据类型 | 频率 | 平均 IO 大小 | 总次/点数 | 首/末采集时间（相对当前时间） | 当前错误/时间（相对当前时间） | 平均/最大采集消耗 | 奔溃次数 |
| ----   |:----:    |----      | ---- | :----:       | :----:    | :----:                        | :----:                        | :----:            | :----:   |
`
		rowFmt = "|%s|%d|%s|%s|%d|%d/%d|%s/%s|%s/%s|%s/%s|%d|"
	)

	if len(x.EnabledInputs) == 0 {
		return "没有开启任何采集器"
	}

	now := time.Now()

	rows := []string{}
	for _, v := range x.EnabledInputs {
		if s, ok := x.InputsStats[v.Input]; !ok {
			rows = append(rows, fmt.Sprintf(rowFmt,
				v.Input, v.Instances,
				"-", "-", 0, 0, 0, "-", "-", "-", "-", "-", "-", 0))
			continue
		} else {
			firstIO := now.Sub(s.First)
			lastIO := now.Sub(s.Last)

			lastErr := "-"
			if s.LastErr != "" {
				lastErr = s.LastErr
			}

			lastErrTime := "-"
			if s.LastErr != "" {
				lastErrTime = fmt.Sprintf("%s", now.Sub(s.LastErrTS))
			}

			freq := "-"
			if s.Frequency != "" {
				freq = s.Frequency
			}

			rows = append(rows, fmt.Sprintf(rowFmt,
				v.Input, v.Instances,
				s.Category, freq, s.AvgSize, s.Count, s.Total, firstIO,
				lastIO, lastErr, lastErrTime, s.AvgCollectCost, s.MaxCollectCost, v.Panics))
		}
	}

	sort.Strings(rows)
	return tblHeader + strings.Join(rows, "\n")
}

func getStats() (*datakitStats, error) {

	stats := &datakitStats{
		Version:      git.Version,
		BuildAt:      git.BuildAt,
		Branch:       git.Branch,
		Uptime:       fmt.Sprintf("%v", time.Since(uptime)),
		OSArch:       fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		ReloadCnt:    reloadCnt,
		WithinDocker: datakit.Docker,
		IOChanStat:   io.ChanStat(),
	}

	if reloadCnt > 0 {
		stats.Reload = reload
	}

	var err error

	stats.InputsStats, err = io.GetStats(time.Second * 5) // get all inputs stats
	if err != nil {
		return nil, err
	}

	for k := range inputs.Inputs {
		if !datakit.Enabled(k) {
			continue
		}

		n, cfgs := inputs.InputEnabled(k)
		npanic := inputs.GetPanicCnt(k)
		if n > 0 {
			stats.EnabledInputs = append(stats.EnabledInputs, &enabledInput{Input: k, Instances: n, Cfgs: cfgs, Panics: npanic})
		}
	}

	for k := range inputs.Inputs {
		if !datakit.Enabled(k) {
			continue
		}
		stats.AvailableInputs = append(stats.AvailableInputs, fmt.Sprintf("[D] %s", k))
	}

	// add available inputs(datakit) stats
	stats.AvailableInputs = append(stats.AvailableInputs, fmt.Sprintf("tatal %d, datakit %d",
		len(stats.AvailableInputs), len(inputs.Inputs)))

	sort.Strings(stats.AvailableInputs)
	return stats, nil
}

func apiGetDatakitMonitor(c *gin.Context) {
	s, err := getStats()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	temp, err := template.New("").Parse(monitorTmpl)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	s.CSS = man.MarkdownCSS

	var buf bytes.Buffer
	if err := temp.Execute(&buf, s); err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	mdext := parser.CommonExtensions
	psr := parser.NewWithExtensions(mdext)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags}
	//opts := html.RendererOptions{Flags: htmlFlags, Head: headerScript}
	renderer := html.NewRenderer(opts)

	out := markdown.ToHTML(buf.Bytes(), psr, renderer)

	c.Data(http.StatusOK, "text/html; charset=UTF-8", out)
}

func apiGetDatakitStats(c *gin.Context) {

	s, err := getStats()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	body, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}