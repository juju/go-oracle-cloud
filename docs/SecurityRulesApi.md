# \SecurityRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecRule**](SecurityRulesApi.md#AddSecRule) | **Post** /secrule/ | Create a Security Rule
[**DeleteSecRule**](SecurityRulesApi.md#DeleteSecRule) | **Delete** /secrule/{name} | Delete a Security Rule
[**DiscoverRootSecRule**](SecurityRulesApi.md#DiscoverRootSecRule) | **Get** /secrule/ | Retrieve Names of Containers
[**DiscoverSecRule**](SecurityRulesApi.md#DiscoverSecRule) | **Get** /secrule/{container} | Retrieve Names of all Security Rules in a Container
[**GetSecRule**](SecurityRulesApi.md#GetSecRule) | **Get** /secrule/{name} | Retrieve Details of a Security Rule
[**ListSecRule**](SecurityRulesApi.md#ListSecRule) | **Get** /secrule/{container}/ | Retrieve Details of all Security Rules in a Container
[**UpdateSecRule**](SecurityRulesApi.md#UpdateSecRule) | **Put** /secrule/{name} | Update a Security Rule


# **AddSecRule**
> SecRuleResponse AddSecRule($body)

Create a Security Rule

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecRulePostRequest**](SecRulePostRequest.md)|  | 

### Return type

[**SecRuleResponse**](SecRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecRule**
> DeleteSecRule($name)

Delete a Security Rule

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootSecRule**
> SecRuleDiscoverResponse DiscoverRootSecRule()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SecRuleDiscoverResponse**](SecRule-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverSecRule**
> SecRuleDiscoverResponse DiscoverSecRule($container)

Retrieve Names of all Security Rules in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 

### Return type

[**SecRuleDiscoverResponse**](SecRule-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecRule**
> SecRuleResponse GetSecRule($name)

Retrieve Details of a Security Rule

Retrieves details of the specified security rule. You can use this request to verify whether POST and PUT HTTP requests were completed successfully.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

[**SecRuleResponse**](SecRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecRule**
> SecRuleListResponse ListSecRule($container, $dstList, $description, $srcList, $disabled, $application, $name)

Retrieve Details of all Security Rules in a Container

Retrieves details of the security rules that are in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the security rules in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;/code&gt; for user-defined security rules. | 
 **dstList** | **string**| &lt;p&gt;The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of the destination security list or security IP list.&lt;p&gt;You must use the prefix &lt;code&gt;seclist&lt;/code&gt;: or &lt;code&gt;seciplist&lt;/code&gt;: to identify the list type.&lt;p&gt;&lt;b&gt;Note:&lt;/b&gt; You can specify a security IP list as the destination in a secrule, provided &lt;code&gt;src_list&lt;/code&gt; is a security list that has DENY as its outbound policy.&lt;p&gt;You cannot specify any of the security IP lists in the &lt;code&gt;/oracle/public&lt;/code&gt; container as a destination in a secrule. | [optional] 
 **description** | **string**| &lt;p&gt;A description of the security rule. | [optional] 
 **srcList** | **string**| &lt;p&gt;The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of the source security list or security IP list.&lt;p&gt;You must use the prefix &lt;code&gt;seclist&lt;/code&gt;: or &lt;code&gt;seciplist&lt;/code&gt;: to identify the list type. | [optional] 
 **disabled** | **bool**| &lt;p&gt;Indicates whether the security rule is enabled (set to &lt;code&gt;false&lt;/code&gt;) or disabled (&lt;code&gt;true&lt;/code&gt;). The default setting is &lt;code&gt;false&lt;/code&gt;. | [optional] 
 **application** | **string**| &lt;p&gt;The three-part name of the security application: (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) for user-defined security applications and &lt;code&gt;/oracle/public/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt; for predefined security applications. | [optional] 
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] 

### Return type

[**SecRuleListResponse**](SecRule-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecRule**
> SecRuleResponse UpdateSecRule($name, $body)

Update a Security Rule

Disables or enables the specified security rule. You can also update the description of the security rule.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 
 **body** | [**SecRulePutRequest**](SecRulePutRequest.md)|  | 

### Return type

[**SecRuleResponse**](SecRule-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

