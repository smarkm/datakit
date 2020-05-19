package rds

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

// ModifyDBProxyInstance invokes the rds.ModifyDBProxyInstance API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbproxyinstance.html
func (client *Client) ModifyDBProxyInstance(request *ModifyDBProxyInstanceRequest) (response *ModifyDBProxyInstanceResponse, err error) {
	response = CreateModifyDBProxyInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBProxyInstanceWithChan invokes the rds.ModifyDBProxyInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbproxyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBProxyInstanceWithChan(request *ModifyDBProxyInstanceRequest) (<-chan *ModifyDBProxyInstanceResponse, <-chan error) {
	responseChan := make(chan *ModifyDBProxyInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBProxyInstance(request)
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

// ModifyDBProxyInstanceWithCallback invokes the rds.ModifyDBProxyInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbproxyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBProxyInstanceWithCallback(request *ModifyDBProxyInstanceRequest, callback func(response *ModifyDBProxyInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBProxyInstanceResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBProxyInstance(request)
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

// ModifyDBProxyInstanceRequest is the request struct for api ModifyDBProxyInstance
type ModifyDBProxyInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EffectiveTime         string           `position:"Query" name:"EffectiveTime"`
	EffectiveSpecificTime string           `position:"Query" name:"EffectiveSpecificTime"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	DBProxyInstanceNum    string           `position:"Query" name:"DBProxyInstanceNum"`
	DBProxyInstanceType   string           `position:"Query" name:"DBProxyInstanceType"`
}

// ModifyDBProxyInstanceResponse is the response struct for api ModifyDBProxyInstance
type ModifyDBProxyInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBProxyInstanceRequest creates a request to invoke ModifyDBProxyInstance API
func CreateModifyDBProxyInstanceRequest() (request *ModifyDBProxyInstanceRequest) {
	request = &ModifyDBProxyInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBProxyInstance", "rds", "openAPI")
	return
}

// CreateModifyDBProxyInstanceResponse creates a response to parse from ModifyDBProxyInstance response
func CreateModifyDBProxyInstanceResponse() (response *ModifyDBProxyInstanceResponse) {
	response = &ModifyDBProxyInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
