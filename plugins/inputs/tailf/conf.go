package tailf

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/hpcloud/tail"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
)

func (t *Tailf) loadcfg() bool {
	var err error

	t.composePipeline()
	t.composeTailerConf()
	t.composeService()
	t.composeTags()

	multilineConfig := &MultilineConfig{
		Pattern:        t.Match,
		InvertMatch:    true,
		MatchWhichLine: "previous",
	}

	for {
		select {
		case <-datakit.Exit.Wait():
			t.log.Info("exit")
			return true
		default:
			// nil
		}

		if t.Source == "" {
			err = fmt.Errorf("tailf source cannot be empty")
			goto label
		}

		// FIXME: add t.log.Debuf("check xxx") ?
		if t.decoder, err = NewDecoder(t.CharacterEncoding); err != nil {
			goto label
		}

		if t.multiline, err = multilineConfig.NewMultiline(); err != nil {
			goto label
		}

		if t.watcher, err = fsnotify.NewWatcher(); err != nil {
			goto label
		}

		if err = checkPipeLine(t.Pipeline); err != nil {
			goto label
		} else {
			break
		}

	label:
		t.log.Error(err)
		time.Sleep(time.Second)
	}

	return false
}

func (t *Tailf) composePipeline() {
	// 兼容旧版配置 pipeline_path
	if t.Pipeline == "" && t.DeprecatedPipeline != "" {
		t.Pipeline = t.DeprecatedPipeline
	}

	if t.Pipeline == "" {
		t.Pipeline = filepath.Join(datakit.PipelineDir, t.Source+".p")
	} else {
		t.Pipeline = filepath.Join(datakit.PipelineDir, t.Pipeline)
	}

	if isExist(t.Pipeline) {
		t.log.Debugf("use pipeline %s", t.Pipeline)
	} else {
		t.Pipeline = ""
		t.log.Warn("no pipeline applied")
	}
}

func (t *Tailf) composeTailerConf() {
	var seek *tail.SeekInfo
	if !t.FromBeginning {
		seek = &tail.SeekInfo{
			Whence: 2, // seek is 2
			Offset: 0,
		}
	}

	t.tailerConf = tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  seek,
		MustExist: true,
		Poll:      false, // default watch method is "inotify"
		Pipe:      false,
		Logger:    tail.DiscardingLogger,
	}
}

func (t *Tailf) composeService() {
	if t.Service == "" {
		t.Service = t.Source
	}
}

func (t *Tailf) composeTags() {
	if t.Tags == nil {
		t.Tags = make(map[string]string)
	}
	// 覆盖自定义tags
	t.Tags["service"] = t.Service
}
