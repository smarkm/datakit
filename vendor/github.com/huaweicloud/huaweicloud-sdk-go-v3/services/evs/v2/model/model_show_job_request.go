/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
