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

// DeleteUserUsageDataExportTask invokes the cdn.DeleteUserUsageDataExportTask API synchronously
// api document: https://help.aliyun.com/api/cdn/deleteuserusagedataexporttask.html
func (client *Client) DeleteUserUsageDataExportTask(request *DeleteUserUsageDataExportTaskRequest) (response *DeleteUserUsageDataExportTaskResponse, err error) {
	response = CreateDeleteUserUsageDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserUsageDataExportTaskWithChan invokes the cdn.DeleteUserUsageDataExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/deleteuserusagedataexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserUsageDataExportTaskWithChan(request *DeleteUserUsageDataExportTaskRequest) (<-chan *DeleteUserUsageDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteUserUsageDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserUsageDataExportTask(request)
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

// DeleteUserUsageDataExportTaskWithCallback invokes the cdn.DeleteUserUsageDataExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/deleteuserusagedataexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserUsageDataExportTaskWithCallback(request *DeleteUserUsageDataExportTaskRequest, callback func(response *DeleteUserUsageDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserUsageDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserUsageDataExportTask(request)
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

// DeleteUserUsageDataExportTaskRequest is the request struct for api DeleteUserUsageDataExportTask
type DeleteUserUsageDataExportTaskRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	TaskId  string           `position:"Query" name:"TaskId"`
}

// DeleteUserUsageDataExportTaskResponse is the response struct for api DeleteUserUsageDataExportTask
type DeleteUserUsageDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUserUsageDataExportTaskRequest creates a request to invoke DeleteUserUsageDataExportTask API
func CreateDeleteUserUsageDataExportTaskRequest() (request *DeleteUserUsageDataExportTaskRequest) {
	request = &DeleteUserUsageDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DeleteUserUsageDataExportTask", "", "")
	return
}

// CreateDeleteUserUsageDataExportTaskResponse creates a response to parse from DeleteUserUsageDataExportTask response
func CreateDeleteUserUsageDataExportTaskResponse() (response *DeleteUserUsageDataExportTaskResponse) {
	response = &DeleteUserUsageDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
