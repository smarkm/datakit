package ecs

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

// ModifyInstanceChargeType invokes the ecs.ModifyInstanceChargeType API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancechargetype.html
func (client *Client) ModifyInstanceChargeType(request *ModifyInstanceChargeTypeRequest) (response *ModifyInstanceChargeTypeResponse, err error) {
	response = CreateModifyInstanceChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceChargeTypeWithChan invokes the ecs.ModifyInstanceChargeType API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancechargetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceChargeTypeWithChan(request *ModifyInstanceChargeTypeRequest) (<-chan *ModifyInstanceChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceChargeType(request)
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

// ModifyInstanceChargeTypeWithCallback invokes the ecs.ModifyInstanceChargeType API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancechargetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceChargeTypeWithCallback(request *ModifyInstanceChargeTypeRequest, callback func(response *ModifyInstanceChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceChargeType(request)
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

// ModifyInstanceChargeTypeRequest is the request struct for api ModifyInstanceChargeType
type ModifyInstanceChargeTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	IsDetailFee          requests.Boolean `position:"Query" name:"IsDetailFee"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	Period               requests.Integer `position:"Query" name:"Period"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	IncludeDataDisks     requests.Boolean `position:"Query" name:"IncludeDataDisks"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
}

// ModifyInstanceChargeTypeResponse is the response struct for api ModifyInstanceChargeType
type ModifyInstanceChargeTypeResponse struct {
	*responses.BaseResponse
	RequestId      string                                   `json:"RequestId" xml:"RequestId"`
	OrderId        string                                   `json:"OrderId" xml:"OrderId"`
	FeeOfInstances FeeOfInstancesInModifyInstanceChargeType `json:"FeeOfInstances" xml:"FeeOfInstances"`
}

// CreateModifyInstanceChargeTypeRequest creates a request to invoke ModifyInstanceChargeType API
func CreateModifyInstanceChargeTypeRequest() (request *ModifyInstanceChargeTypeRequest) {
	request = &ModifyInstanceChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceChargeType", "ecs", "openAPI")
	return
}

// CreateModifyInstanceChargeTypeResponse creates a response to parse from ModifyInstanceChargeType response
func CreateModifyInstanceChargeTypeResponse() (response *ModifyInstanceChargeTypeResponse) {
	response = &ModifyInstanceChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}