package tailf

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/hpcloud/tail"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
)

type tailer struct {
	tf *Tailf

	filename string
	source   string
	tags     map[string]string

	tail *tail.Tail
	pipe *pipeline.Pipeline

	textLine    bytes.Buffer
	tailerOpen  bool
	channelOpen bool
	count       int64
}

func newTailer(tl *Tailf, filename string) *tailer {
	t := tailer{tf: tl, filename: filename, source: tl.Source}

	t.tags = func() map[string]string {
		var m = make(map[string]string)

		for k, v := range tl.Tags {
			m[k] = v
		}

		if _, ok := m["filename"]; !ok {
			m["filename"] = filename
		}
		return m
	}()

	t.tailerOpen = true
	t.channelOpen = true

	return &t
}

func (t *tailer) run() {
	var err error

	t.tail, err = tail.TailFile(t.filename, t.tf.tailerConf)
	if err != nil {
		l.Error("failed of build tailer, err:%s", err)
		return
	}
	defer t.tail.Cleanup()

	if t.tf.PipelinePath != "" {
		t.pipe, err = pipeline.NewPipelineFromFile(t.tf.PipelinePath)
		if err != nil {
			l.Error("failed of pipeline, err:%s", err)
			return
		}
	}

	t.receiver()
}

func (t *tailer) receiver() {
	ticker := time.NewTicker(defaultDruation)
	defer ticker.Stop()

	var (
		buffer bytes.Buffer
		line   *tail.Line
	)

	for {
		line = nil

		select {
		case <-datakit.Exit.Wait():
			l.Debugf("Tailing source:%s, file %s is ending", t.source, t.filename)
			return

		case line, t.tailerOpen = <-t.tail.Lines:
			if !t.tailerOpen {
				t.channelOpen = false
			}

		case <-ticker.C:
			if t.count > 0 {
				if err := io.NamedFeed(buffer.Bytes(), io.Logging, t.source); err != nil {
					l.Error(err)
				}
				buffer = bytes.Buffer{}
				t.count = 0
			}
			_, statErr := os.Lstat(t.filename)
			if os.IsNotExist(statErr) {
				l.Warnf("check file %s is not exist", t.filename)
				return
			}
		}

		text, status := t.multiline(line)
		switch status {
		case _return:
			return
		case _continue:
			continue
		case _next:
			//pass
		}

		text, err := t.decode(text)
		if err != nil {
			continue
		}

		data, err := t.pipeline(text)
		if err != nil {
			continue
		}

		buffer.Write(data)
		buffer.WriteString("\n")
		t.count++

		if t.count >= metricFeedCount {
			if err := io.NamedFeed(buffer.Bytes(), io.Logging, t.source); err != nil {
				l.Error(err)
			}
			buffer = bytes.Buffer{}
			t.count = 0
		}
	}
}

type multilineStatus int

const (
	// tail channel 关闭，执行 return
	_return multilineStatus = iota
	// multiline 判断数据为多行，将数据存入缓存，继续读取下一行
	_continue
	// multiline 判断多行数据结束，将缓存中的数据放出，继续执行后续处理
	_next
)

func (t *tailer) multiline(line *tail.Line) (text string, status multilineStatus) {
	if line != nil {
		text = strings.TrimRight(line.Text, "\r")

		if t.tf.multiline.IsEnabled() {
			if text = t.tf.multiline.ProcessLine(text, &t.textLine); text == "" {
				status = _continue
				return
			}
		}
	}

	if line == nil || !t.channelOpen || !t.tailerOpen {
		if text += t.tf.multiline.Flush(&t.textLine); text == "" {
			if !t.channelOpen {
				status = _return
				l.Warnf("Tailing %s data channel is closed", t.filename)
				return
			}

			status = _continue
			return
		}
	}

	if line != nil && line.Err != nil {
		l.Errorf("Tailing %q: %s", t.filename, line.Err.Error())
		status = _continue
		return
	}

	status = _next
	return
}

func (t *tailer) decode(text string) (str string, err error) {
	str, err = t.tf.decoder.String(text)
	if err != nil {
		l.Errorf("decode error, %s", err) // only print err
	}
	return
}

func (t *tailer) pipeline(text string) (data []byte, err error) {
	var fields = make(map[string]interface{})

	if t.pipe != nil {
		fields, err = t.pipe.Run(text).Result()
		if err != nil {
			l.Errorf("run pipeline error, %s", err)
			return
		}
	} else {
		fields["message"] = text
	}

	var ts time.Time

	if v, ok := fields["time"]; ok { // time should be nano-second
		nanots, ok := v.(int64)
		if !ok {
			l.Warn("filed `time' should be nano-second, but got `%s'", reflect.TypeOf(v).String())
			err = fmt.Errorf("invalid fileds time")
			return
		}

		ts = time.Unix(nanots/int64(time.Second), nanots%int64(time.Second))
		delete(fields, "time")
	} else {
		ts = time.Now()
	}

	data, err = io.MakeMetric(t.source, t.tags, fields, ts)
	if err != nil {
		l.Error(err)
	}

	return
}