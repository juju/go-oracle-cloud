# \OrchestrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrchestration**](OrchestrationsApi.md#AddOrchestration) | **Post** /orchestration/ | Add an Orchestration
[**DeleteOrchestration**](OrchestrationsApi.md#DeleteOrchestration) | **Delete** /orchestration/{name} | Delete an Orchestration
[**DiscoverOrchestration**](OrchestrationsApi.md#DiscoverOrchestration) | **Get** /orchestration/{container} | Retrieve Names of all Orchestrations in a Container
[**DiscoverRootOrchestration**](OrchestrationsApi.md#DiscoverRootOrchestration) | **Get** /orchestration/ | Retrieve Names of Containers
[**GetOrchestration**](OrchestrationsApi.md#GetOrchestration) | **Get** /orchestration/{name} | Retrieve Details of an Orchestration
[**ListOrchestration**](OrchestrationsApi.md#ListOrchestration) | **Get** /orchestration/{container}/ | Retrieve Details of all Orchestrations in a Container
[**UpdateOrchestration**](OrchestrationsApi.md#UpdateOrchestration) | **Put** /orchestration/{name} | Update an Orchestration


# **AddOrchestration**
> OrchestrationResponse AddOrchestration($body)

Add an Orchestration

Adds an orchestration to Oracle Compute Cloud Service.<p>After adding an orchestration, you can start the orchestration by using the HTTP request /orchestration/{name}?action=START to create all the objects you have defined in the orchestration JSON file. See <a class=\"xref\" href=\"op-orchestration-{name}-put.html\">Update an Orchestration</a>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrchestrationPostRequest**](OrchestrationPostRequest.md)|  | 

### Return type

[**OrchestrationResponse**](Orchestration-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrchestration**
> DeleteOrchestration($name)

Delete an Orchestration

Deletes an orchestration from the system. The orchestration must be stopped to be deleted. All the objects created by the orchestration are deleted when you stop the orchestration. No response is returned for the delete action.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverOrchestration**
> OrchestrationDiscoverResponse DiscoverOrchestration($container)

Retrieve Names of all Orchestrations in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 

### Return type

[**OrchestrationDiscoverResponse**](Orchestration-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootOrchestration**
> OrchestrationDiscoverResponse DiscoverRootOrchestration()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**OrchestrationDiscoverResponse**](Orchestration-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestration**
> OrchestrationResponse GetOrchestration($name)

Retrieve Details of an Orchestration

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

[**OrchestrationResponse**](Orchestration-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrchestration**
> OrchestrationListResponse ListOrchestration($container, $status)

Retrieve Details of all Orchestrations in a Container

Retrieves details of the orchestrations that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the orchestrations in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **status** | **string**| &lt;p&gt;Specify one of the following values:&lt;p&gt;starting&lt;p&gt;scheduled&lt;p&gt;ready&lt;p&gt;updating&lt;p&gt;error&lt;p&gt;stopping&lt;p&gt;stopped | [optional] 

### Return type

[**OrchestrationListResponse**](Orchestration-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrchestration**
> OrchestrationResponse UpdateOrchestration($name, $action, $body)

Update an Orchestration

Updates an orchestration. You can update an orchestration by modifying the orchestration JSON file. Ensure that you specify the name of the orchestration that you want to update in the updated orchestration JSON file.<p>You can update orchestrations when they are in one of the following states: <code>stopped</code> and <code>running</code>.<p>Depending on the state of an orchestration, you can make the following updates:<ul><li>When an orchestration is in the <code>stopped</code> state, you can update all the oplan attributes, add oplans, or remove oplans.</li><li>When an orchestration is running, you can only modify the HA policy of the existing oplans. You cannot modify any other attributes of the existing oplans. You can add or remove oplans at runtime.</li></ul><p>When you update an orchestration, the name of the orchestration must be specified in the PUT body.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **action** | **string**| Use the query argument action&#x3D;START/STOP to start/stop an orchestration immediately.&lt;ul&gt;&lt;li&gt;Specify the &lt;code&gt;action&#x3D;START&lt;/code&gt; query argument to immediately start an orchestration that has already been added to the system.&lt;p&gt;Starting an orchestration creates all the objects defined in the orchestration. The status of the orchestration changes over time. You can view the details of the orchestration to see the status as the orchestration progresses.&lt;/p&gt;&lt;p&gt;&lt;b&gt;Note: &lt;/b&gt;If the orchestration is in a scheduled state, you must stop the orchestration to cancel the schedule before you can start the orchestration. See Stop an Orchestration. After stopping the orchestration, you can either start it immediately, or update the orchestration to specify a new start time.&lt;/li&gt;&lt;li&gt;Specify the &lt;code&gt;action&#x3D;STOP&lt;/code&gt; query argument to immediately stop a running orchestration. Stopping an orchestration deletes all objects that were created by the orchestration. The status of the orchestration changes over time. You can get the details of the orchestration to view the status as the orchestration stops.&lt;p&gt;&lt;b&gt;Note: &lt;/b&gt;If you want to stop an orchestration at some future time, you should update the orchestration object and specify the &lt;code&gt;stop_time&lt;/code&gt; in the request body.&lt;/p&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
 **body** | [**OrchestrationPutRequest**](OrchestrationPutRequest.md)|  | [optional] 

### Return type

[**OrchestrationResponse**](Orchestration-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

