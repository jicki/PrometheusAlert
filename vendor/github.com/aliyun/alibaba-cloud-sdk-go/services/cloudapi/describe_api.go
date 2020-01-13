package cloudapi

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

// DescribeApi invokes the cloudapi.DescribeApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapi.html
func (client *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
	response = CreateDescribeApiResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiWithChan invokes the cloudapi.DescribeApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiWithChan(request *DescribeApiRequest) (<-chan *DescribeApiResponse, <-chan error) {
	responseChan := make(chan *DescribeApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApi(request)
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

// DescribeApiWithCallback invokes the cloudapi.DescribeApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiWithCallback(request *DescribeApiRequest, callback func(response *DescribeApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiResponse
		var err error
		defer close(result)
		response, err = client.DescribeApi(request)
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

// DescribeApiRequest is the request struct for api DescribeApi
type DescribeApiRequest struct {
	*requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeApiResponse is the response struct for api DescribeApi
type DescribeApiResponse struct {
	*responses.BaseResponse
	RequestId              string                              `json:"RequestId" xml:"RequestId"`
	RegionId               string                              `json:"RegionId" xml:"RegionId"`
	ApiId                  string                              `json:"ApiId" xml:"ApiId"`
	ApiName                string                              `json:"ApiName" xml:"ApiName"`
	GroupId                string                              `json:"GroupId" xml:"GroupId"`
	GroupName              string                              `json:"GroupName" xml:"GroupName"`
	Visibility             string                              `json:"Visibility" xml:"Visibility"`
	AuthType               string                              `json:"AuthType" xml:"AuthType"`
	ResultType             string                              `json:"ResultType" xml:"ResultType"`
	ResultSample           string                              `json:"ResultSample" xml:"ResultSample"`
	FailResultSample       string                              `json:"FailResultSample" xml:"FailResultSample"`
	CreatedTime            string                              `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime           string                              `json:"ModifiedTime" xml:"ModifiedTime"`
	Description            string                              `json:"Description" xml:"Description"`
	Mock                   string                              `json:"Mock" xml:"Mock"`
	MockResult             string                              `json:"MockResult" xml:"MockResult"`
	AllowSignatureMethod   string                              `json:"AllowSignatureMethod" xml:"AllowSignatureMethod"`
	WebSocketApiType       string                              `json:"WebSocketApiType" xml:"WebSocketApiType"`
	ResultBodyModel        string                              `json:"ResultBodyModel" xml:"ResultBodyModel"`
	ForceNonceCheck        bool                                `json:"ForceNonceCheck" xml:"ForceNonceCheck"`
	DisableInternet        bool                                `json:"DisableInternet" xml:"DisableInternet"`
	RequestConfig          RequestConfig                       `json:"RequestConfig" xml:"RequestConfig"`
	ServiceConfig          ServiceConfig                       `json:"ServiceConfig" xml:"ServiceConfig"`
	OpenIdConnectConfig    OpenIdConnectConfig                 `json:"OpenIdConnectConfig" xml:"OpenIdConnectConfig"`
	ErrorCodeSamples       ErrorCodeSamplesInDescribeApi       `json:"ErrorCodeSamples" xml:"ErrorCodeSamples"`
	ResultDescriptions     ResultDescriptionsInDescribeApi     `json:"ResultDescriptions" xml:"ResultDescriptions"`
	SystemParameters       SystemParametersInDescribeApi       `json:"SystemParameters" xml:"SystemParameters"`
	CustomSystemParameters CustomSystemParametersInDescribeApi `json:"CustomSystemParameters" xml:"CustomSystemParameters"`
	ConstantParameters     ConstantParametersInDescribeApi     `json:"ConstantParameters" xml:"ConstantParameters"`
	RequestParameters      RequestParametersInDescribeApi      `json:"RequestParameters" xml:"RequestParameters"`
	ServiceParameters      ServiceParametersInDescribeApi      `json:"ServiceParameters" xml:"ServiceParameters"`
	ServiceParametersMap   ServiceParametersMapInDescribeApi   `json:"ServiceParametersMap" xml:"ServiceParametersMap"`
	DeployedInfos          DeployedInfos                       `json:"DeployedInfos" xml:"DeployedInfos"`
}

// CreateDescribeApiRequest creates a request to invoke DescribeApi API
func CreateDescribeApiRequest() (request *DescribeApiRequest) {
	request = &DescribeApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApi", "apigateway", "openAPI")
	return
}

// CreateDescribeApiResponse creates a response to parse from DescribeApi response
func CreateDescribeApiResponse() (response *DescribeApiResponse) {
	response = &DescribeApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}