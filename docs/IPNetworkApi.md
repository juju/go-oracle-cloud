# \IPNetworkApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIpNetwork**](IPNetworkApi.md#AddIpNetwork) | **Post** /network/v1/ipnetwork/ | Create an IP Network
[**DeleteIpNetwork**](IPNetworkApi.md#DeleteIpNetwork) | **Delete** /network/v1/ipnetwork/{name} | Delete an IP Network
[**GetIpNetwork**](IPNetworkApi.md#GetIpNetwork) | **Get** /network/v1/ipnetwork/{name} | Retrieve Details of an IP Network
[**ListIpNetwork**](IPNetworkApi.md#ListIpNetwork) | **Get** /network/v1/ipnetwork/{container}/ | Retrieve Details of all IP Networks in a Container
[**UpdateIpNetwork**](IPNetworkApi.md#UpdateIpNetwork) | **Put** /network/v1/ipnetwork/{name} | Update an IP Network


# **AddIpNetwork**
> IpNetworkResponse AddIpNetwork($body)

Create an IP Network

Creates an IP network. An IP network allows you to define an IP subnet in your account. With an IP network you can isolate instances by creating separate IP networks and adding instances to specific networks. Traffic can flow between instances within the same IP network, but by default each network is isolated from other networks and from the public Internet.<p><b>Required Role: </b> To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpNetworkPostRequest**](IpNetworkPostRequest.md)|  | 

### Return type

[**IpNetworkResponse**](IpNetwork-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpNetwork**
> DeleteIpNetwork($name)

Delete an IP Network

<b>Required Role: </b> To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpNetwork**
> IpNetworkResponse GetIpNetwork($name)

Retrieve Details of an IP Network

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network. | 

### Return type

[**IpNetworkResponse**](IpNetwork-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpNetwork**
> IpNetworkListResponse ListIpNetwork($container)

Retrieve Details of all IP Networks in a Container

Retrieves details of all the IP networks that are available in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 

### Return type

[**IpNetworkListResponse**](IpNetwork-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpNetwork**
> IpNetworkResponse UpdateIpNetwork($name, $body)

Update an IP Network

You can update an IP network and change the specified IP address prefix for the network after you've created the network and attached instances to it. However, when you change an IP address prefix, it could cause the IP addresses currently assigned to existing instances to fall outside the specified IP network. If this happens, all traffic to and from those vNICs will be dropped.<p>If the IP address of an instance is dynamically allocated, stopping the instance orchestration and restarting it will reassign a valid IP address from the IP network to the instance.<p>However, if the IP address of an instance is static - that is, if the IP address is specified in the instance orchestration while creating the instance - then the IP address can't be updated by stopping the instance orchestration and restarting it. You would have to manually update the orchestration to assign a valid IP address to the vNIC attached to that IP network.<p>It is therefore recommended that if you update an IP network, you only expand the network by specifying the same IP address prefix but with a shorter prefix length. For example, you can expand 192.168.1.0/24 to 192.168.1.0/20. Don't, however, change the IP address. This ensures that all IP addresses that have been currently allocated to instances remain valid in the updated IP network.<p><b>Required Role: </b> To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network. | 
 **body** | [**IpNetworkPutRequest**](IpNetworkPutRequest.md)|  | 

### Return type

[**IpNetworkResponse**](IpNetwork-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

