package domain

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

// QueryArtExtension invokes the domain.QueryArtExtension API synchronously
// api document: https://help.aliyun.com/api/domain/queryartextension.html
func (client *Client) QueryArtExtension(request *QueryArtExtensionRequest) (response *QueryArtExtensionResponse, err error) {
	response = CreateQueryArtExtensionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryArtExtensionWithChan invokes the domain.QueryArtExtension API asynchronously
// api document: https://help.aliyun.com/api/domain/queryartextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryArtExtensionWithChan(request *QueryArtExtensionRequest) (<-chan *QueryArtExtensionResponse, <-chan error) {
	responseChan := make(chan *QueryArtExtensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryArtExtension(request)
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

// QueryArtExtensionWithCallback invokes the domain.QueryArtExtension API asynchronously
// api document: https://help.aliyun.com/api/domain/queryartextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryArtExtensionWithCallback(request *QueryArtExtensionRequest, callback func(response *QueryArtExtensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryArtExtensionResponse
		var err error
		defer close(result)
		response, err = client.QueryArtExtension(request)
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

// QueryArtExtensionRequest is the request struct for api QueryArtExtension
type QueryArtExtensionRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryArtExtensionResponse is the response struct for api QueryArtExtension
type QueryArtExtensionResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	ObjectType              string `json:"ObjectType" xml:"ObjectType"`
	MaterialsAndTechniques  string `json:"MaterialsAndTechniques" xml:"MaterialsAndTechniques"`
	Dimensions              string `json:"Dimensions" xml:"Dimensions"`
	Title                   string `json:"Title" xml:"Title"`
	DateOrPeriod            string `json:"DateOrPeriod" xml:"DateOrPeriod"`
	Maker                   string `json:"Maker" xml:"Maker"`
	InscriptionsAndMarkings string `json:"InscriptionsAndMarkings" xml:"InscriptionsAndMarkings"`
	Subject                 string `json:"Subject" xml:"Subject"`
	Features                string `json:"Features" xml:"Features"`
	Reference               string `json:"Reference" xml:"Reference"`
}

// CreateQueryArtExtensionRequest creates a request to invoke QueryArtExtension API
func CreateQueryArtExtensionRequest() (request *QueryArtExtensionRequest) {
	request = &QueryArtExtensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryArtExtension", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryArtExtensionResponse creates a response to parse from QueryArtExtension response
func CreateQueryArtExtensionResponse() (response *QueryArtExtensionResponse) {
	response = &QueryArtExtensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
