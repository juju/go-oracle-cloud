# \SecurityAssociationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecAssociation**](SecurityAssociationsApi.md#AddSecAssociation) | **Post** /secassociation/ | Create a Security Association
[**DeleteSecAssociation**](SecurityAssociationsApi.md#DeleteSecAssociation) | **Delete** /secassociation/{name} | Delete a Security Association
[**DiscoverRootSecAssociation**](SecurityAssociationsApi.md#DiscoverRootSecAssociation) | **Get** /secassociation/ | Retrieve Names of Containers
[**DiscoverSecAssociation**](SecurityAssociationsApi.md#DiscoverSecAssociation) | **Get** /secassociation/{container} | Retrieve Names of all Security Associations in a Container
[**GetSecAssociation**](SecurityAssociationsApi.md#GetSecAssociation) | **Get** /secassociation/{name} | Retrieve Details of a Security Association
[**ListSecAssociation**](SecurityAssociationsApi.md#ListSecAssociation) | **Get** /secassociation/{container}/ | Retrieve Details of all Security Associations in a Container


# **AddSecAssociation**
> SecAssociationResponse AddSecAssociation($body)

Create a Security Association

Adds an instance to a security list.<p><b>Caution: </b>When you add an instance to a security list, all the security rules that use that security list-as either the source or destination-are applicable to the instance. Consider a security list that is the destination in two security rules, one rule that allows SSH access from the public Internet and another rule permitting HTTPS traffic from the public Internet. When you add an instance to this security list, the instance is accessible from the public Internet over both SSH and HTTPS. Keep this in mind when you decide the security lists that you want to add an instance to.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecAssociationPostRequest**](SecAssociationPostRequest.md)|  | 

### Return type

[**SecAssociationResponse**](SecAssociation-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecAssociation**
> DeleteSecAssociation($name)

Delete a Security Association

Deletes the specified security association. After you delete a security association, it takes a few minutes for the change to take effect. No response is returned.<p><b>Note:</b> You can delete all security associations, if you want to. If you delete all the security associations, instances will not be associated with any security list, and you cannot access the instance by NAT or flat IP. You can only access instances which are added to at least one security association.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


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

# **DiscoverRootSecAssociation**
> SecAssociationDiscoverResponse DiscoverRootSecAssociation()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SecAssociationDiscoverResponse**](SecAssociation-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverSecAssociation**
> SecAssociationDiscoverResponse DiscoverSecAssociation($container)

Retrieve Names of all Security Associations in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Specify &lt;code&gt;/oracle/public&lt;/code&gt; to retrieve the names of system-provided objects. | 

### Return type

[**SecAssociationDiscoverResponse**](SecAssociation-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecAssociation**
> SecAssociationResponse GetSecAssociation($name)

Retrieve Details of a Security Association

Retrieves details about the specified security association.<p>You can use this request to verify whether the POST HTTP request was completed successfully.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

[**SecAssociationResponse**](SecAssociation-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecAssociation**
> SecAssociationListResponse ListSecAssociation($container, $vcable, $seclist, $name)

Retrieve Details of all Security Associations in a Container

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; | 
 **vcable** | **string**| &lt;p&gt;vcable of the instance that you want to associate with the security list.&lt;p&gt;For more information about the vcable of an instance, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-instance-{name}-get.html\&quot;&gt;Retrieve Details of an Instance.&lt;/a&gt; | [optional] 
 **seclist** | **string**| &lt;p&gt;Security list that you want to associate with the instance. | [optional] 
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] 

### Return type

[**SecAssociationListResponse**](SecAssociation-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

