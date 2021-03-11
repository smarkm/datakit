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

// Response Object
type ShowProjectQuotaResponse struct {
	Quotas         *QuotaResult `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowProjectQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectQuotaResponse", string(data)}, " ")
}