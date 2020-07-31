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

// DescribePrice invokes the ecs.DescribePrice API synchronously
// api document: https://help.aliyun.com/api/ecs/describeprice.html
func (client *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	response = CreateDescribePriceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePriceWithChan invokes the ecs.DescribePrice API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePriceWithChan(request *DescribePriceRequest) (<-chan *DescribePriceResponse, <-chan error) {
	responseChan := make(chan *DescribePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePrice(request)
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

// DescribePriceWithCallback invokes the ecs.DescribePrice API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePriceWithCallback(request *DescribePriceRequest, callback func(response *DescribePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePriceResponse
		var err error
		defer close(result)
		response, err = client.DescribePrice(request)
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

// DescribePriceRequest is the request struct for api DescribePrice
type DescribePriceRequest struct {
	*requests.RpcRequest
	DataDisk3Size              requests.Integer `position:"Query" name:"DataDisk.3.Size"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DataDisk3Category          string           `position:"Query" name:"DataDisk.3.Category"`
	DataDisk4Size              requests.Integer `position:"Query" name:"DataDisk.4.Size"`
	PriceUnit                  string           `position:"Query" name:"PriceUnit"`
	Period                     requests.Integer `position:"Query" name:"Period"`
	DataDisk1PerformanceLevel  string           `position:"Query" name:"DataDisk.1.PerformanceLevel"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	InternetChargeType         string           `position:"Query" name:"InternetChargeType"`
	InstanceNetworkType        string           `position:"Query" name:"InstanceNetworkType"`
	InstanceAmount             requests.Integer `position:"Query" name:"InstanceAmount"`
	DataDisk3PerformanceLevel  string           `position:"Query" name:"DataDisk.3.PerformanceLevel"`
	ImageId                    string           `position:"Query" name:"ImageId"`
	IoOptimized                string           `position:"Query" name:"IoOptimized"`
	InternetMaxBandwidthOut    requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskCategory         string           `position:"Query" name:"SystemDisk.Category"`
	Platform                   string           `position:"Query" name:"Platform"`
	Capacity                   requests.Integer `position:"Query" name:"Capacity"`
	SystemDiskPerformanceLevel string           `position:"Query" name:"SystemDisk.PerformanceLevel"`
	DataDisk4Category          string           `position:"Query" name:"DataDisk.4.Category"`
	DataDisk4PerformanceLevel  string           `position:"Query" name:"DataDisk.4.PerformanceLevel"`
	Scope                      string           `position:"Query" name:"Scope"`
	InstanceType               string           `position:"Query" name:"InstanceType"`
	DataDisk2Category          string           `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size              requests.Integer `position:"Query" name:"DataDisk.1.Size"`
	Amount                     requests.Integer `position:"Query" name:"Amount"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	DataDisk2Size              requests.Integer `position:"Query" name:"DataDisk.2.Size"`
	ResourceType               string           `position:"Query" name:"ResourceType"`
	DataDisk1Category          string           `position:"Query" name:"DataDisk.1.Category"`
	DataDisk2PerformanceLevel  string           `position:"Query" name:"DataDisk.2.PerformanceLevel"`
	SystemDiskSize             requests.Integer `position:"Query" name:"SystemDisk.Size"`
	OfferingType               string           `position:"Query" name:"OfferingType"`
}

// DescribePriceResponse is the response struct for api DescribePrice
type DescribePriceResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	PriceInfo PriceInfo `json:"PriceInfo" xml:"PriceInfo"`
}

// CreateDescribePriceRequest creates a request to invoke DescribePrice API
func CreateDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribePrice", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePriceResponse creates a response to parse from DescribePrice response
func CreateDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
