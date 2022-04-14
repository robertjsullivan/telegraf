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

// CreateMonitorGroup invokes the cms.CreateMonitorGroup API synchronously
func (client *Client) CreateMonitorGroup(request *CreateMonitorGroupRequest) (response *CreateMonitorGroupResponse, err error) {
	response = CreateCreateMonitorGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMonitorGroupWithChan invokes the cms.CreateMonitorGroup API asynchronously
func (client *Client) CreateMonitorGroupWithChan(request *CreateMonitorGroupRequest) (<-chan *CreateMonitorGroupResponse, <-chan error) {
	responseChan := make(chan *CreateMonitorGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMonitorGroup(request)
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

// CreateMonitorGroupWithCallback invokes the cms.CreateMonitorGroup API asynchronously
func (client *Client) CreateMonitorGroupWithCallback(request *CreateMonitorGroupRequest, callback func(response *CreateMonitorGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMonitorGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateMonitorGroup(request)
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

// CreateMonitorGroupRequest is the request struct for api CreateMonitorGroup
type CreateMonitorGroupRequest struct {
	*requests.RpcRequest
	ContactGroups string           `position:"Query" name:"ContactGroups"`
	Type          string           `position:"Query" name:"Type"`
	GroupName     string           `position:"Query" name:"GroupName"`
	Options       string           `position:"Query" name:"Options"`
	ServiceId     requests.Integer `position:"Query" name:"ServiceId"`
	BindUrl       string           `position:"Query" name:"BindUrl"`
}

// CreateMonitorGroupResponse is the response struct for api CreateMonitorGroup
type CreateMonitorGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   int64  `json:"GroupId" xml:"GroupId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateMonitorGroupRequest creates a request to invoke CreateMonitorGroup API
func CreateCreateMonitorGroupRequest() (request *CreateMonitorGroupRequest) {
	request = &CreateMonitorGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateMonitorGroup", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMonitorGroupResponse creates a response to parse from CreateMonitorGroup response
func CreateCreateMonitorGroupResponse() (response *CreateMonitorGroupResponse) {
	response = &CreateMonitorGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}