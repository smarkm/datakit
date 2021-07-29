package kubernetes

import (
	"fmt"
	"time"

	batchv1 "k8s.io/api/batch/v1"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

const kubernetesJobName = "kubernetes_jobs"

type job struct {
	client interface {
		getJobs() (*batchv1.JobList, error)
	}
}

func (j job) Gather() {
	list, err := j.client.getJobs()
	if err != nil {
		l.Errorf("failed of get nodes resource: %s", err)
		return
	}

	for _, obj := range list.Items {
		tags := map[string]string{
			"name":         fmt.Sprintf("%v", obj.UID),
			"job_name":     obj.Name,
			"cluster_name": obj.ClusterName,
			"pod_name":     obj.ClusterName,
			"namespace":    obj.Namespace,
		}
		fields := map[string]interface{}{
			"age": int64(time.Now().Sub(obj.CreationTimestamp.Time).Seconds()),

			"active":    obj.Status.Active,
			"succeeded": obj.Status.Succeeded,
			"failed":    obj.Status.Failed,
		}

		if obj.Spec.Parallelism != nil {
			fields["parallelism"] = *obj.Spec.Parallelism
		}
		if obj.Spec.Completions != nil {
			fields["completions"] = *obj.Spec.Completions
		}
		if obj.Spec.ActiveDeadlineSeconds != nil {
			fields["active_deadline"] = *obj.Spec.ActiveDeadlineSeconds
		}
		if obj.Spec.BackoffLimit != nil {
			fields["backoff_limit"] = *obj.Spec.BackoffLimit
		}

		addJSONStringToMap("kubernetes_labels", obj.Labels, fields)
		addJSONStringToMap("kubernetes_annotations", obj.Annotations, fields)
		addMessageToFields(tags, fields)

		pt, err := io.MakePoint(kubernetesJobName, tags, fields, time.Now())
		if err != nil {
			l.Error(err)
		} else {
			if err := io.Feed(inputName, datakit.Object, []*io.Point{pt}, nil); err != nil {
				l.Error(err)
			}
		}
	}
}

func (*job) LineProto() (*io.Point, error) {
	return nil, nil
}

func (*job) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: kubernetesJobName,
		Desc: kubernetesJobName,
		Tags: map[string]interface{}{
			"name":         inputs.NewTagInfo(""),
			"job_name":     inputs.NewTagInfo(""),
			"cluster_name": inputs.NewTagInfo(""),
			// "status":       inputs.NewTagInfo(""),
			"namespace": inputs.NewTagInfo(""),
		},
		Fields: map[string]interface{}{
			"age":       &inputs.FieldInfo{DataType: inputs.Int, Unit: inputs.UnknownUnit, Desc: ""},
			"active":    &inputs.FieldInfo{DataType: inputs.Int, Unit: inputs.NCount, Desc: ""},
			"succeeded": &inputs.FieldInfo{DataType: inputs.Int, Unit: inputs.NCount, Desc: ""},
			"failed":    &inputs.FieldInfo{DataType: inputs.Int, Unit: inputs.NCount, Desc: ""},
			// "pod_statuses":           &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"completions":     &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"parallelism":     &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"backoff_limit":   &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"active_deadline": &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			//"duration":               &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"kubernetes_labels":      &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"kubernetes_annotations": &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
			"message":                &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: ""},
		},
	}
}
