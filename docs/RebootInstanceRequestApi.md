# \RebootInstanceRequestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRebootInstanceRequest**](RebootInstanceRequestApi.md#AddRebootInstanceRequest) | **Post** /rebootinstancerequest/ | Create a Reboot Instance Request
[**DeleteRebootInstanceRequest**](RebootInstanceRequestApi.md#DeleteRebootInstanceRequest) | **Delete** /rebootinstancerequest/{name} | Delete a Reboot Instance Request
[**DiscoverRebootInstanceRequest**](RebootInstanceRequestApi.md#DiscoverRebootInstanceRequest) | **Get** /rebootinstancerequest/{container} | Retrieve Names of all Reboot Instance Requests in a Container
[**DiscoverRootRebootInstanceRequest**](RebootInstanceRequestApi.md#DiscoverRootRebootInstanceRequest) | **Get** /rebootinstancerequest/ | Retrieve Names of Containers
[**GetRebootInstanceRequest**](RebootInstanceRequestApi.md#GetRebootInstanceRequest) | **Get** /rebootinstancerequest/{name} | Retrieve Details of a Reboot Instance Request
[**ListRebootInstanceRequest**](RebootInstanceRequestApi.md#ListRebootInstanceRequest) | **Get** /rebootinstancerequest/{container}/ | Retrieve Details of all Reboot Instance Requests in a Container


# **AddRebootInstanceRequest**
> RebootInstanceRequestResponse AddRebootInstanceRequest($body)

Create a Reboot Instance Request

If your instance hangs after it starts running, you can use this request to reboot your instance. After creating this request, use <a class=\"xref\" href=\"op-rebootinstancerequest-{name}-get.html\">GET /rebootinstancerequest/{name}</a> to retrieve the status of the request. When the status of the <code>rebootinstancerequest</code> changes to <code>complete</code>, you know that the instance has been rebooted. <p>To reboot the instance, you need to know its name, which you can view with the <a class=\"xref\" href=\"op-instance-{container}--get.html\">GET /instance/{container}/</a> request.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RebootInstanceRequestPostRequest**](RebootInstanceRequestPostRequest.md)|  | 

### Return type

[**RebootInstanceRequestResponse**](RebootInstanceRequest-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRebootInstanceRequest**
> DeleteRebootInstanceRequest($name)

Delete a Reboot Instance Request

Deletes a reboot instance request. No response is returned for the delete action.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the rebootinstancerequest object that you want to delete. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRebootInstanceRequest**
> RebootInstanceRequestDiscoverResponse DiscoverRebootInstanceRequest($container)

Retrieve Names of all Reboot Instance Requests in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 

### Return type

[**RebootInstanceRequestDiscoverResponse**](RebootInstanceRequest-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootRebootInstanceRequest**
> RebootInstanceRequestDiscoverResponse DiscoverRootRebootInstanceRequest()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**RebootInstanceRequestDiscoverResponse**](RebootInstanceRequest-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRebootInstanceRequest**
> RebootInstanceRequestResponse GetRebootInstanceRequest($name)

Retrieve Details of a Reboot Instance Request

Retrieves details of the specified reboot instance request. You can use this request when you want to find out the status of a reboot instance request.<p>When you create a reboot instance request, the status of the request changes from <code>queued</code> to <code>active</code>, and then to <code>complete</code>. When status is <code>active</code>, the instance starts getting rebooted. When the reboot of the instance is complete, the status changes to <code>complete</code>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the object. | 

### Return type

[**RebootInstanceRequestResponse**](RebootInstanceRequest-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRebootInstanceRequest**
> RebootInstanceRequestListResponse ListRebootInstanceRequest($container, $instance, $hard, $name)

Retrieve Details of all Reboot Instance Requests in a Container

Retrieves details of the reboot instance requests that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the reboot instance requests in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify parent path of the object name, such as &lt;code&gt;/Compute-acme&lt;/code&gt;. | 
 **instance** | **string**| Name of the instance that is rebooted. | [optional] 
 **hard** | **bool**| When set to &lt;code&gt;true&lt;/code&gt;, a hard reset is performed. | [optional] 
 **name** | **string**| Unique identifier generated by the server. | [optional] 

### Return type

[**RebootInstanceRequestListResponse**](RebootInstanceRequest-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

