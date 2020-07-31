package slb

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

// DescribeLoadBalancerHTTPListenerAttribute invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerhttplistenerattribute.html
func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerHTTPListenerAttributeWithChan invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerhttplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithChan(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (<-chan *DescribeLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeWithCallback invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerhttplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithCallback(request *DescribeLoadBalancerHTTPListenerAttributeRequest, callback func(response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeRequest is the request struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerHTTPListenerAttributeResponse is the response struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                  string                                            `json:"RequestId" xml:"RequestId"`
	ListenerPort               int                                               `json:"ListenerPort" xml:"ListenerPort"`
	BackendServerPort          int                                               `json:"BackendServerPort" xml:"BackendServerPort"`
	Bandwidth                  int                                               `json:"Bandwidth" xml:"Bandwidth"`
	Status                     string                                            `json:"Status" xml:"Status"`
	SecurityStatus             string                                            `json:"SecurityStatus" xml:"SecurityStatus"`
	XForwardedFor              string                                            `json:"XForwardedFor" xml:"XForwardedFor"`
	Scheduler                  string                                            `json:"Scheduler" xml:"Scheduler"`
	StickySession              string                                            `json:"StickySession" xml:"StickySession"`
	StickySessionType          string                                            `json:"StickySessionType" xml:"StickySessionType"`
	CookieTimeout              int                                               `json:"CookieTimeout" xml:"CookieTimeout"`
	Cookie                     string                                            `json:"Cookie" xml:"Cookie"`
	HealthCheck                string                                            `json:"HealthCheck" xml:"HealthCheck"`
	HealthCheckType            string                                            `json:"HealthCheckType" xml:"HealthCheckType"`
	HealthCheckDomain          string                                            `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckURI             string                                            `json:"HealthCheckURI" xml:"HealthCheckURI"`
	HealthyThreshold           int                                               `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold         int                                               `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckTimeout         int                                               `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	HealthCheckInterval        int                                               `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckConnectPort     int                                               `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckHttpCode        string                                            `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	HealthCheckMethod          string                                            `json:"HealthCheckMethod" xml:"HealthCheckMethod"`
	HealthCheckHttpVersion     string                                            `json:"HealthCheckHttpVersion" xml:"HealthCheckHttpVersion"`
	MaxConnection              int                                               `json:"MaxConnection" xml:"MaxConnection"`
	VServerGroupId             string                                            `json:"VServerGroupId" xml:"VServerGroupId"`
	Gzip                       string                                            `json:"Gzip" xml:"Gzip"`
	XForwardedForSLBIP         string                                            `json:"XForwardedFor_SLBIP" xml:"XForwardedFor_SLBIP"`
	XForwardedForSLBID         string                                            `json:"XForwardedFor_SLBID" xml:"XForwardedFor_SLBID"`
	XForwardedForProto         string                                            `json:"XForwardedFor_proto" xml:"XForwardedFor_proto"`
	AclId                      string                                            `json:"AclId" xml:"AclId"`
	AclType                    string                                            `json:"AclType" xml:"AclType"`
	AclStatus                  string                                            `json:"AclStatus" xml:"AclStatus"`
	VpcIds                     string                                            `json:"VpcIds" xml:"VpcIds"`
	ListenerForward            string                                            `json:"ListenerForward" xml:"ListenerForward"`
	ForwardPort                int                                               `json:"ForwardPort" xml:"ForwardPort"`
	RequestTimeout             int                                               `json:"RequestTimeout" xml:"RequestTimeout"`
	IdleTimeout                int                                               `json:"IdleTimeout" xml:"IdleTimeout"`
	Description                string                                            `json:"Description" xml:"Description"`
	XForwardedForSLBPORT       string                                            `json:"XForwardedFor_SLBPORT" xml:"XForwardedFor_SLBPORT"`
	XForwardedForClientSrcPort string                                            `json:"XForwardedFor_ClientSrcPort" xml:"XForwardedFor_ClientSrcPort"`
	ForwardCode                int                                               `json:"ForwardCode" xml:"ForwardCode"`
	AclIds                     AclIdsInDescribeLoadBalancerHTTPListenerAttribute `json:"AclIds" xml:"AclIds"`
	Rules                      RulesInDescribeLoadBalancerHTTPListenerAttribute  `json:"Rules" xml:"Rules"`
}

// CreateDescribeLoadBalancerHTTPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerHTTPListenerAttribute API
func CreateDescribeLoadBalancerHTTPListenerAttributeRequest() (request *DescribeLoadBalancerHTTPListenerAttributeRequest) {
	request = &DescribeLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerHTTPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerHTTPListenerAttribute response
func CreateDescribeLoadBalancerHTTPListenerAttributeResponse() (response *DescribeLoadBalancerHTTPListenerAttributeResponse) {
	response = &DescribeLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
