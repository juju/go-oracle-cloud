# \IPNetworkExchangeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIpNetworkExchange**](IPNetworkExchangeApi.md#AddIpNetworkExchange) | **Post** /network/v1/ipnetworkexchange/ | Create an IP Network Exchange
[**DeleteIpNetworkExchange**](IPNetworkExchangeApi.md#DeleteIpNetworkExchange) | **Delete** /network/v1/ipnetworkexchange/{name} | Delete an IP Network Exchange
[**GetIpNetworkExchange**](IPNetworkExchangeApi.md#GetIpNetworkExchange) | **Get** /network/v1/ipnetworkexchange/{name} | Retrieve Details of an IP Network Exchange
[**ListIpNetworkExchange**](IPNetworkExchangeApi.md#ListIpNetworkExchange) | **Get** /network/v1/ipnetworkexchange/{container}/ | Retrieve Details of all IP Network Exchanges in a Container


# **AddIpNetworkExchange**
> IpNetworkExchangeResponse AddIpNetworkExchange($body)

Create an IP Network Exchange

Create an IP network exchange to enable access between IP networks that have non-overlapping addresses, so that instances on these networks can exchange packets with each other without NAT.<p>After creating an IP network exchange, you can add IP networks to the same IP network exchange to enable access between instances on these IP networks. Use <a target=\"_blank\" href=\"http://ohcrest.doceng.oraclecorp.com/client/compute/compute/html/op-network-v1-ipnetwork-%7Bname%7D-put.html\">PUT /network/v1/ipnetwork/{name}</a> request to add an IP network to an IP network exchange. An IP network exchange can include multiple IP networks, but an IP network can be added to only one IP network exchange.<p><b>Required Role: </b> To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpNetworkExchangePostRequest**](IpNetworkExchangePostRequest.md)|  | 

### Return type

[**IpNetworkExchangeResponse**](IpNetworkExchange-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpNetworkExchange**
> DeleteIpNetworkExchange($name)

Delete an IP Network Exchange

<b>Required Role: </b> To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network exchange. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpNetworkExchange**
> IpNetworkExchangeResponse GetIpNetworkExchange($name)

Retrieve Details of an IP Network Exchange

Retrieves details of a specific IP network exchange.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network exchange. | 

### Return type

[**IpNetworkExchangeResponse**](IpNetworkExchange-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpNetworkExchange**
> IpNetworkExchangeListResponse ListIpNetworkExchange($container)

Retrieve Details of all IP Network Exchanges in a Container

Retrieves details of all the IP network exchanges that are available in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 

### Return type

[**IpNetworkExchangeListResponse**](IpNetworkExchange-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

