// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package worker

import (
	"testing"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/parser"
)

func TestResult_checkFieldValLen(t *testing.T) {
	type fields struct {
		output      *parser.Output
		measurement string
		ts          time.Time
		err         string
	}
	type args struct {
		messageLen int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case1",
			fields: fields{
				output: &parser.Output{
					Dropped: false,
					Error:   "",
					Cost:    nil,
					Tags:    nil,
					Fields: map[string]interface{}{
						"msg":         "0123456789",
						"message":     "0123456789",
						"other_field": "0123456789",
					},
				},
				measurement: "logging",
				ts:          time.Now(),
				err:         "",
			},
			args: args{messageLen: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &pipeline.Result{
				Output: tt.fields.output,
				TS:     tt.fields.ts,
				Err:    tt.fields.err,
			}
			r.CheckFieldValLen(tt.args.messageLen)

			for key := range r.Output.Fields {
				if i, err := r.GetField(key); err == nil {
					if mass, isString := i.(string); isString {
						if len(mass) > tt.args.messageLen {
							t.Errorf("key=%s val=%s over massageLen:%d", key, mass, tt.args.messageLen)
						}
					}
				}
			}
		})
	}
}
