package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyHostAvailability invokes the cms.ModifyHostAvailability API synchronously
// api document: https://help.aliyun.com/api/cms/modifyhostavailability.html
func (client *Client) ModifyHostAvailability(request *ModifyHostAvailabilityRequest) (response *ModifyHostAvailabilityResponse, err error) {
	response = CreateModifyHostAvailabilityResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyHostAvailabilityWithChan invokes the cms.ModifyHostAvailability API asynchronously
// api document: https://help.aliyun.com/api/cms/modifyhostavailability.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHostAvailabilityWithChan(request *ModifyHostAvailabilityRequest) (<-chan *ModifyHostAvailabilityResponse, <-chan error) {
	responseChan := make(chan *ModifyHostAvailabilityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHostAvailability(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyHostAvailabilityWithCallback invokes the cms.ModifyHostAvailability API asynchronously
// api document: https://help.aliyun.com/api/cms/modifyhostavailability.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHostAvailabilityWithCallback(request *ModifyHostAvailabilityRequest, callback func(response *ModifyHostAvailabilityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHostAvailabilityResponse
		var err error
		defer close(result)
		response, err = client.ModifyHostAvailability(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyHostAvailabilityRequest is the request struct for api ModifyHostAvailability
type ModifyHostAvailabilityRequest struct {
	*requests.RpcRequest
	TaskOptionHttpMethod               string                                             `position:"Query" name:"TaskOption.HttpMethod"`
	AlertConfigEscalationList          *[]ModifyHostAvailabilityAlertConfigEscalationList `position:"Query" name:"AlertConfigEscalationList"  type:"Repeated"`
	TaskName                           string                                             `position:"Query" name:"TaskName"`
	AlertConfigSilenceTime             requests.Integer                                   `position:"Query" name:"AlertConfig.SilenceTime"`
	TaskOptionHttpResponseCharset      string                                             `position:"Query" name:"TaskOption.HttpResponseCharset"`
	TaskOptionHttpNegative             requests.Boolean                                   `position:"Query" name:"TaskOption.HttpNegative"`
	AlertConfigNotifyType              requests.Integer                                   `position:"Query" name:"AlertConfig.NotifyType"`
	TaskOptionTelnetOrPingHost         string                                             `position:"Query" name:"TaskOption.TelnetOrPingHost"`
	TaskOptionHttpResponseMatchContent string                                             `position:"Query" name:"TaskOption.HttpResponseMatchContent"`
	Id                                 requests.Integer                                   `position:"Query" name:"Id"`
	InstanceList                       *[]string                                          `position:"Query" name:"InstanceList"  type:"Repeated"`
	GroupId                            requests.Integer                                   `position:"Query" name:"GroupId"`
	AlertConfigEndTime                 requests.Integer                                   `position:"Query" name:"AlertConfig.EndTime"`
	TaskOptionHttpURI                  string                                             `position:"Query" name:"TaskOption.HttpURI"`
	TaskScope                          string                                             `position:"Query" name:"TaskScope"`
	AlertConfigStartTime               requests.Integer                                   `position:"Query" name:"AlertConfig.StartTime"`
	AlertConfigWebHook                 string                                             `position:"Query" name:"AlertConfig.WebHook"`
}

// ModifyHostAvailabilityAlertConfigEscalationList is a repeated param struct in ModifyHostAvailabilityRequest
type ModifyHostAvailabilityAlertConfigEscalationList struct {
	Times      string `name:"Times"`
	MetricName string `name:"MetricName"`
	Value      string `name:"Value"`
	Operator   string `name:"Operator"`
	Aggregate  string `name:"Aggregate"`
}

// ModifyHostAvailabilityResponse is the response struct for api ModifyHostAvailability
type ModifyHostAvailabilityResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyHostAvailabilityRequest creates a request to invoke ModifyHostAvailability API
func CreateModifyHostAvailabilityRequest() (request *ModifyHostAvailabilityRequest) {
	request = &ModifyHostAvailabilityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "ModifyHostAvailability", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyHostAvailabilityResponse creates a response to parse from ModifyHostAvailability response
func CreateModifyHostAvailabilityResponse() (response *ModifyHostAvailabilityResponse) {
	response = &ModifyHostAvailabilityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
