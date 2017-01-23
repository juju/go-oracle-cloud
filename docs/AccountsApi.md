# \AccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverAccount**](AccountsApi.md#DiscoverAccount) | **Get** /account/{container} | Retrieve Names of all Accounts in a Container
[**DiscoverRootAccount**](AccountsApi.md#DiscoverRootAccount) | **Get** /account/ | Retrieve Names of Containers
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /account/{name} | Retrieve Details of an Account
[**ListAccount**](AccountsApi.md#ListAccount) | **Get** /account/{container}/ | Retrieve Details of all Accounts in a Container


# **DiscoverAccount**
> AccountDiscoverResponse DiscoverAccount($container)

Retrieve Names of all Accounts in a Container

Retrieves names of all the accounts in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;&lt;/code&gt; to retrieve the names of objects that you can access. | 

### Return type

[**AccountDiscoverResponse**](Account-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootAccount**
> AccountDiscoverResponse DiscoverRootAccount()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**AccountDiscoverResponse**](Account-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccount**
> AccountResponse GetAccount($name)

Retrieve Details of an Account

Retrieves details of the specified account.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Two-part name of the account: &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/default&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/cloud_storage&lt;/code&gt; | 

### Return type

[**AccountResponse**](Account-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccount**
> AccountListResponse ListAccount($container)

Retrieve Details of all Accounts in a Container

Retrieves details of the accounts that are in the specified identity domain. You can use this HTTP request to get details of the account that you must specify while creating a machine image.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt;&lt;p&gt;For example: &lt;code&gt;Compute-acme&lt;/code&gt;. | 

### Return type

[**AccountListResponse**](Account-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

