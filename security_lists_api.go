/* 
 * REST API for Oracle Compute Cloud Service (IaaS)
 *
 * Use the Oracle Compute Cloud Service REST API to provision and manage instances and the associated resources
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"net/url"
	"encoding/json"
	"fmt"
	"strings"
)

type SecurityListsApi struct {
	Configuration Configuration
}

func NewSecurityListsApi() *SecurityListsApi {
	configuration := NewConfiguration()
	return &SecurityListsApi{
		Configuration: *configuration,
	}
}

func NewSecurityListsApiWithBasePath(basePath string) *SecurityListsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SecurityListsApi{
		Configuration: *configuration,
	}
}

/**
 * Create a Security List
 * Creates a security list. After creating security lists, you can add instances to them by using the HTTP request, POST /secassociation/ &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-secassociation--post.html\&quot;&gt;(Create a Security Association)&lt;/a&gt;.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param body 
 * @return *SecListResponse
 */
func (a SecurityListsApi) AddSecList(body SecListPostRequest) (*SecListResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(SecListResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete a Security List
 * Deletes the specified security list. No response is returned.&lt;p&gt;Note that when a security list is deleted, all the security associations involving that security list are also deleted.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @return void
 */
func (a SecurityListsApi) DeleteSecList(name string) (*APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Retrieve Names of Containers
 * Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @return *SecListDiscoverResponse
 */
func (a SecurityListsApi) DiscoverRootSecList() (*SecListDiscoverResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+directory+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SecListDiscoverResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Retrieve Names of all Security Lists in a Container
 * Retrieves the names of objects and subcontainers that you can access in the specified container.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identity_domain&lt;/i&gt;/default&lt;/code&gt; to retrieve name of the predefined default security list.
 * @return *SecListDiscoverResponse
 */
func (a SecurityListsApi) DiscoverSecList(container string) (*SecListDiscoverResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/{container}"
	path = strings.Replace(path, "{"+"container"+"}", fmt.Sprintf("%v", container), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+directory+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SecListDiscoverResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Retrieve Details of a Security List
 * Retrieves information about the specified security list.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @return *SecListResponse
 */
func (a SecurityListsApi) GetSecList(name string) (*SecListResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SecListResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Retrieve Details of all Security Lists in a Container
 * Retrieves details of the security lists that are in the specified container and match the specified query criteria. If you don&#39;t specify any query criteria, then details of all the security lists in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:&lt;p&gt;&lt;code&gt;?parameter1&#x3D;value1&amp;ampparameter2&#x3D;value2&amp;ampparameterN&#x3D;valueN&lt;/code&gt;&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Monitor&lt;/code&gt; or &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param container &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; and &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;/code&gt;for user-defined security lists and &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/default&lt;/code&gt; for the predefined default security list.
 * @param name The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/default/default&lt;/code&gt;).
 * @return *SecListListResponse
 */
func (a SecurityListsApi) ListSecList(container string, name string) (*SecListListResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/{container}/"
	path = strings.Replace(path, "{"+"container"+"}", fmt.Sprintf("%v", container), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("name", a.Configuration.APIClient.ParameterToString(name, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SecListListResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update a Security List
 * Updates inbound policy, outbound policy, and description for the specified security list.&lt;p&gt;&lt;b&gt;Required Role: &lt;/b&gt;To complete this task, you must have the &lt;code&gt;Compute_Operations&lt;/code&gt; role. If this role isn&#39;t assigned to you or you&#39;re not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\&quot;&gt;Modifying User Roles&lt;/a&gt; in &lt;em&gt;Managing and Monitoring Oracle Cloud&lt;/em&gt;.
 *
 * @param name &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).
 * @param body 
 * @return *SecListResponse
 */
func (a SecurityListsApi) UpdateSecList(name string, body SecListPutRequest) (*SecListResponse, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/seclist/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(SecListResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

