package smartag

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

// DescribeUserOnlineClients invokes the smartag.DescribeUserOnlineClients API synchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclients.html
func (client *Client) DescribeUserOnlineClients(request *DescribeUserOnlineClientsRequest) (response *DescribeUserOnlineClientsResponse, err error) {
	response = CreateDescribeUserOnlineClientsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserOnlineClientsWithChan invokes the smartag.DescribeUserOnlineClients API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclients.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserOnlineClientsWithChan(request *DescribeUserOnlineClientsRequest) (<-chan *DescribeUserOnlineClientsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserOnlineClientsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserOnlineClients(request)
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

// DescribeUserOnlineClientsWithCallback invokes the smartag.DescribeUserOnlineClients API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclients.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserOnlineClientsWithCallback(request *DescribeUserOnlineClientsRequest, callback func(response *DescribeUserOnlineClientsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserOnlineClientsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserOnlineClients(request)
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

// DescribeUserOnlineClientsRequest is the request struct for api DescribeUserOnlineClients
type DescribeUserOnlineClientsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// DescribeUserOnlineClientsResponse is the response struct for api DescribeUserOnlineClients
type DescribeUserOnlineClientsResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Users     UsersInDescribeUserOnlineClients `json:"Users" xml:"Users"`
}

// CreateDescribeUserOnlineClientsRequest creates a request to invoke DescribeUserOnlineClients API
func CreateDescribeUserOnlineClientsRequest() (request *DescribeUserOnlineClientsRequest) {
	request = &DescribeUserOnlineClientsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeUserOnlineClients", "smartag", "openAPI")
	return
}

// CreateDescribeUserOnlineClientsResponse creates a response to parse from DescribeUserOnlineClients response
func CreateDescribeUserOnlineClientsResponse() (response *DescribeUserOnlineClientsResponse) {
	response = &DescribeUserOnlineClientsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}