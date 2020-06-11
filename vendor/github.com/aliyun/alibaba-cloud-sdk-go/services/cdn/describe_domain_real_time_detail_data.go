package cdn

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

// DescribeDomainRealTimeDetailData invokes the cdn.DescribeDomainRealTimeDetailData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimedetaildata.html
func (client *Client) DescribeDomainRealTimeDetailData(request *DescribeDomainRealTimeDetailDataRequest) (response *DescribeDomainRealTimeDetailDataResponse, err error) {
	response = CreateDescribeDomainRealTimeDetailDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeDetailDataWithChan invokes the cdn.DescribeDomainRealTimeDetailData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimedetaildata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeDetailDataWithChan(request *DescribeDomainRealTimeDetailDataRequest) (<-chan *DescribeDomainRealTimeDetailDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeDetailDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeDetailData(request)
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

// DescribeDomainRealTimeDetailDataWithCallback invokes the cdn.DescribeDomainRealTimeDetailData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimedetaildata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeDetailDataWithCallback(request *DescribeDomainRealTimeDetailDataRequest, callback func(response *DescribeDomainRealTimeDetailDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeDetailDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeDetailData(request)
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

// DescribeDomainRealTimeDetailDataRequest is the request struct for api DescribeDomainRealTimeDetailData
type DescribeDomainRealTimeDetailDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	Merge          string           `position:"Query" name:"Merge"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	MergeLocIsp    string           `position:"Query" name:"MergeLocIsp"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Field          string           `position:"Query" name:"Field"`
}

// DescribeDomainRealTimeDetailDataResponse is the response struct for api DescribeDomainRealTimeDetailData
type DescribeDomainRealTimeDetailDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDescribeDomainRealTimeDetailDataRequest creates a request to invoke DescribeDomainRealTimeDetailData API
func CreateDescribeDomainRealTimeDetailDataRequest() (request *DescribeDomainRealTimeDetailDataRequest) {
	request = &DescribeDomainRealTimeDetailDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeDetailData", "", "")
	return
}

// CreateDescribeDomainRealTimeDetailDataResponse creates a response to parse from DescribeDomainRealTimeDetailData response
func CreateDescribeDomainRealTimeDetailDataResponse() (response *DescribeDomainRealTimeDetailDataResponse) {
	response = &DescribeDomainRealTimeDetailDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}