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

// DescribeEventRuleAttribute invokes the cms.DescribeEventRuleAttribute API synchronously
func (client *Client) DescribeEventRuleAttribute(request *DescribeEventRuleAttributeRequest) (response *DescribeEventRuleAttributeResponse, err error) {
	response = CreateDescribeEventRuleAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventRuleAttributeWithChan invokes the cms.DescribeEventRuleAttribute API asynchronously
func (client *Client) DescribeEventRuleAttributeWithChan(request *DescribeEventRuleAttributeRequest) (<-chan *DescribeEventRuleAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeEventRuleAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEventRuleAttribute(request)
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

// DescribeEventRuleAttributeWithCallback invokes the cms.DescribeEventRuleAttribute API asynchronously
func (client *Client) DescribeEventRuleAttributeWithCallback(request *DescribeEventRuleAttributeRequest, callback func(response *DescribeEventRuleAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventRuleAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeEventRuleAttribute(request)
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

// DescribeEventRuleAttributeRequest is the request struct for api DescribeEventRuleAttribute
type DescribeEventRuleAttributeRequest struct {
	*requests.RpcRequest
	RuleName string `position:"Query" name:"RuleName"`
}

// DescribeEventRuleAttributeResponse is the response struct for api DescribeEventRuleAttribute
type DescribeEventRuleAttributeResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeEventRuleAttributeRequest creates a request to invoke DescribeEventRuleAttribute API
func CreateDescribeEventRuleAttributeRequest() (request *DescribeEventRuleAttributeRequest) {
	request = &DescribeEventRuleAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeEventRuleAttribute", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEventRuleAttributeResponse creates a response to parse from DescribeEventRuleAttribute response
func CreateDescribeEventRuleAttributeResponse() (response *DescribeEventRuleAttributeResponse) {
	response = &DescribeEventRuleAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
