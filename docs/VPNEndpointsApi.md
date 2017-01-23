# \VPNEndpointsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVPNEndpoint**](VPNEndpointsApi.md#AddVPNEndpoint) | **Post** /vpnendpoint/ | Create a VPN Endpoint
[**DeleteVPNEndpoint**](VPNEndpointsApi.md#DeleteVPNEndpoint) | **Delete** /vpnendpoint/{name} | Delete a VPN Endpoint
[**GetVPNEndpoint**](VPNEndpointsApi.md#GetVPNEndpoint) | **Get** /vpnendpoint/{name} | Retrieve Details of a VPN Endpoint
[**ListVPNEndpoint**](VPNEndpointsApi.md#ListVPNEndpoint) | **Get** /vpnendpoint/{container}/ | Retrieve Details of all VPN Endpoints in a Container
[**UpdateVPNEndpoint**](VPNEndpointsApi.md#UpdateVPNEndpoint) | **Put** /vpnendpoint/{name} | Update a VPN Endpoint


# **AddVPNEndpoint**
> VpnEndpointResponse AddVPNEndpoint($body)

Create a VPN Endpoint

Creates a VPN tunnel between your data center and your Oracle Compute Cloud Service site. You can create up to 20 VPN tunnels to your Oracle Compute Cloud Service site.<p>Before you create a VPN tunnel, you must complete the following tasks:<ol type='1'><li>Request the Oracle Network Cloud Service - VPN for Dedicated Compute service. For more information, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-8F89288E-1A6F-436E-A422-1F7792291001\">Requesting Oracle Network Cloud Service - VPN for Dedicated Compute</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.</li><li>Configure your VPN gateway to connect to the Oracle Cloud VPN gateway after the Oracle Network Cloud Service - VPN for Dedicated Compute service is provisioned. For more information, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-BDC73B91-C337-45E3-8BB1-096BCFB9B1EC\">Configuring Your VPN Gateway</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.</ol><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VpnEndpointPostRequest**](VpnEndpointPostRequest.md)|  | 

### Return type

[**VpnEndpointResponse**](VPNEndpoint-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVPNEndpoint**
> DeleteVPNEndpoint($name)

Delete a VPN Endpoint

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Two-part name of the object (&lt;code&gt;&lt;em&gt;/Compute-acme/object&lt;/em&gt;&lt;/code&gt;) | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVPNEndpoint**
> VpnEndpointResponse GetVPNEndpoint($name)

Retrieve Details of a VPN Endpoint

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Two-part name of the object (&lt;code&gt;&lt;em&gt;/Compute-acme/object&lt;/em&gt;&lt;/code&gt;) | 

### Return type

[**VpnEndpointResponse**](VPNEndpoint-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVPNEndpoint**
> VpnEndpointListResponse ListVPNEndpoint($container)

Retrieve Details of all VPN Endpoints in a Container

Retrieves details of all the VPN endpoints that are available in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify parent path of the object name, such as &lt;code&gt;/Compute-acme&lt;/code&gt;. | 

### Return type

[**VpnEndpointListResponse**](VPNEndpoint-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVPNEndpoint**
> VpnEndpointResponse UpdateVPNEndpoint($name, $body)

Update a VPN Endpoint

After you've configured your VPN connection, you can update the connection to enable or disable the VPN tunnel, or to change other connection details.<P><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Two-part name of the object (&lt;code&gt;&lt;em&gt;/Compute-acme/object&lt;/em&gt;&lt;/code&gt;) | 
 **body** | [**VpnEndpointPutRequest**](VpnEndpointPutRequest.md)|  | 

### Return type

[**VpnEndpointResponse**](VPNEndpoint-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

