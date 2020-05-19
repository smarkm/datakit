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

// DisableSiteMonitors invokes the cms.DisableSiteMonitors API synchronously
// api document: https://help.aliyun.com/api/cms/disablesitemonitors.html
func (client *Client) DisableSiteMonitors(request *DisableSiteMonitorsRequest) (response *DisableSiteMonitorsResponse, err error) {
	response = CreateDisableSiteMonitorsResponse()
	err = client.DoAction(request, response)
	return
}

// DisableSiteMonitorsWithChan invokes the cms.DisableSiteMonitors API asynchronously
// api document: https://help.aliyun.com/api/cms/disablesitemonitors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableSiteMonitorsWithChan(request *DisableSiteMonitorsRequest) (<-chan *DisableSiteMonitorsResponse, <-chan error) {
	responseChan := make(chan *DisableSiteMonitorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableSiteMonitors(request)
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

// DisableSiteMonitorsWithCallback invokes the cms.DisableSiteMonitors API asynchronously
// api document: https://help.aliyun.com/api/cms/disablesitemonitors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableSiteMonitorsWithCallback(request *DisableSiteMonitorsRequest, callback func(response *DisableSiteMonitorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableSiteMonitorsResponse
		var err error
		defer close(result)
		response, err = client.DisableSiteMonitors(request)
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

// DisableSiteMonitorsRequest is the request struct for api DisableSiteMonitors
type DisableSiteMonitorsRequest struct {
	*requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

// DisableSiteMonitorsResponse is the response struct for api DisableSiteMonitors
type DisableSiteMonitorsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDisableSiteMonitorsRequest creates a request to invoke DisableSiteMonitors API
func CreateDisableSiteMonitorsRequest() (request *DisableSiteMonitorsRequest) {
	request = &DisableSiteMonitorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DisableSiteMonitors", "cms", "openAPI")
	return
}

// CreateDisableSiteMonitorsResponse creates a response to parse from DisableSiteMonitors response
func CreateDisableSiteMonitorsResponse() (response *DisableSiteMonitorsResponse) {
	response = &DisableSiteMonitorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
