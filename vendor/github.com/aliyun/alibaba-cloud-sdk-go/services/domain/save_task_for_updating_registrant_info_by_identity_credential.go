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

// SaveTaskForUpdatingRegistrantInfoByIdentityCredential invokes the domain.SaveTaskForUpdatingRegistrantInfoByIdentityCredential API synchronously
// api document: https://help.aliyun.com/api/domain/savetaskforupdatingregistrantinfobyidentitycredential.html
func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredential(request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) (response *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, err error) {
	response = CreateSaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse()
	err = client.DoAction(request, response)
	return
}

// SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithChan invokes the domain.SaveTaskForUpdatingRegistrantInfoByIdentityCredential API asynchronously
// api document: https://help.aliyun.com/api/domain/savetaskforupdatingregistrantinfobyidentitycredential.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithChan(request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) (<-chan *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, <-chan error) {
	responseChan := make(chan *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveTaskForUpdatingRegistrantInfoByIdentityCredential(request)
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

// SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithCallback invokes the domain.SaveTaskForUpdatingRegistrantInfoByIdentityCredential API asynchronously
// api document: https://help.aliyun.com/api/domain/savetaskforupdatingregistrantinfobyidentitycredential.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithCallback(request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest, callback func(response *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse
		var err error
		defer close(result)
		response, err = client.SaveTaskForUpdatingRegistrantInfoByIdentityCredential(request)
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

// SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest is the request struct for api SaveTaskForUpdatingRegistrantInfoByIdentityCredential
type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
	*requests.RpcRequest
	Country                  string           `position:"Query" name:"Country"`
	IdentityCredentialType   string           `position:"Query" name:"IdentityCredentialType"`
	City                     string           `position:"Query" name:"City"`
	IdentityCredential       string           `position:"Body" name:"IdentityCredential"`
	TransferOutProhibited    requests.Boolean `position:"Query" name:"TransferOutProhibited"`
	ZhCity                   string           `position:"Query" name:"ZhCity"`
	TelExt                   string           `position:"Query" name:"TelExt"`
	Province                 string           `position:"Query" name:"Province"`
	ZhRegistrantName         string           `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string           `position:"Query" name:"PostalCode"`
	Lang                     string           `position:"Query" name:"Lang"`
	Email                    string           `position:"Query" name:"Email"`
	ZhRegistrantOrganization string           `position:"Query" name:"ZhRegistrantOrganization"`
	Address                  string           `position:"Query" name:"Address"`
	TelArea                  string           `position:"Query" name:"TelArea"`
	ZhAddress                string           `position:"Query" name:"ZhAddress"`
	RegistrantType           string           `position:"Query" name:"RegistrantType"`
	DomainName               *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	Telephone                string           `position:"Query" name:"Telephone"`
	ZhProvince               string           `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string           `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string           `position:"Query" name:"UserClientIp"`
	IdentityCredentialNo     string           `position:"Query" name:"IdentityCredentialNo"`
	RegistrantName           string           `position:"Query" name:"RegistrantName"`
}

// SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse is the response struct for api SaveTaskForUpdatingRegistrantInfoByIdentityCredential
type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest creates a request to invoke SaveTaskForUpdatingRegistrantInfoByIdentityCredential API
func CreateSaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest() (request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) {
	request = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForUpdatingRegistrantInfoByIdentityCredential", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse creates a response to parse from SaveTaskForUpdatingRegistrantInfoByIdentityCredential response
func CreateSaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse() (response *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) {
	response = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
