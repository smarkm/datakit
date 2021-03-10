/*
 * IAM
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
type ShowProjectDetailsAndStatusRequest struct {
	ProjectId string `json:"project_id"`
}

func (o ShowProjectDetailsAndStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectDetailsAndStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectDetailsAndStatusRequest", string(data)}, " ")
}
