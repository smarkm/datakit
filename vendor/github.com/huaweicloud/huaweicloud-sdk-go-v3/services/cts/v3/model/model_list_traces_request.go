/*
 * CTS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListTracesRequest struct {
	TraceType    ListTracesRequestTraceType    `json:"trace_type"`
	Limit        *int32                        `json:"limit,omitempty"`
	From         *int64                        `json:"from,omitempty"`
	Next         *string                       `json:"next,omitempty"`
	To           *int64                        `json:"to,omitempty"`
	TrackerName  *string                       `json:"tracker_name,omitempty"`
	ServiceType  *string                       `json:"service_type,omitempty"`
	User         *string                       `json:"user,omitempty"`
	ResourceId   *string                       `json:"resource_id,omitempty"`
	ResourceName *string                       `json:"resource_name,omitempty"`
	ResourceType *string                       `json:"resource_type,omitempty"`
	TraceId      *string                       `json:"trace_id,omitempty"`
	TraceName    *string                       `json:"trace_name,omitempty"`
	TraceRating  *ListTracesRequestTraceRating `json:"trace_rating,omitempty"`
}

func (o ListTracesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTracesRequest struct{}"
	}

	return strings.Join([]string{"ListTracesRequest", string(data)}, " ")
}

type ListTracesRequestTraceType struct {
	value string
}

type ListTracesRequestTraceTypeEnum struct {
	SYSTEM ListTracesRequestTraceType
	DATA   ListTracesRequestTraceType
}

func GetListTracesRequestTraceTypeEnum() ListTracesRequestTraceTypeEnum {
	return ListTracesRequestTraceTypeEnum{
		SYSTEM: ListTracesRequestTraceType{
			value: "system",
		},
		DATA: ListTracesRequestTraceType{
			value: "data",
		},
	}
}

func (c ListTracesRequestTraceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTracesRequestTraceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListTracesRequestTraceRating struct {
	value string
}

type ListTracesRequestTraceRatingEnum struct {
	NORMAL   ListTracesRequestTraceRating
	WARNING  ListTracesRequestTraceRating
	INCIDENT ListTracesRequestTraceRating
}

func GetListTracesRequestTraceRatingEnum() ListTracesRequestTraceRatingEnum {
	return ListTracesRequestTraceRatingEnum{
		NORMAL: ListTracesRequestTraceRating{
			value: "normal",
		},
		WARNING: ListTracesRequestTraceRating{
			value: "warning",
		},
		INCIDENT: ListTracesRequestTraceRating{
			value: "incident",
		},
	}
}

func (c ListTracesRequestTraceRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTracesRequestTraceRating) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
