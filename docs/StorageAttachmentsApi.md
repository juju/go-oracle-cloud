# \StorageAttachmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStorageAttachment**](StorageAttachmentsApi.md#AddStorageAttachment) | **Post** /storage/attachment/ | Create a Storage Attachment
[**DeleteStorageAttachment**](StorageAttachmentsApi.md#DeleteStorageAttachment) | **Delete** /storage/attachment/{name} | Delete a Storage Attachment
[**DiscoverRootStorageAttachment**](StorageAttachmentsApi.md#DiscoverRootStorageAttachment) | **Get** /storage/attachment/ | Retrieve Names of Containers
[**DiscoverStorageAttachment**](StorageAttachmentsApi.md#DiscoverStorageAttachment) | **Get** /storage/attachment/{container} | Retrieve Names of all Storage Attachments in a Container
[**GetStorageAttachment**](StorageAttachmentsApi.md#GetStorageAttachment) | **Get** /storage/attachment/{name} | Retrieve Details of a Storage Attachment
[**ListStorageAttachment**](StorageAttachmentsApi.md#ListStorageAttachment) | **Get** /storage/attachment/{container}/ | Retrieve Details of all Storage Attachments in a Container


# **AddStorageAttachment**
> StorageAttachmentResponse AddStorageAttachment($body)

Create a Storage Attachment

Attaches a storage volume to an instance.<p>Note that, after attaching the volume, you must create a file system and mount the file system on the instance. For more information, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&amp;id=STCSG-GUID-093D0406-28B9-4E8F-B650-57D5CDC56557\">Mounting and Unmounting a Storage Volume</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StorageAttachmentPostRequest**](StorageAttachmentPostRequest.md)|  | 

### Return type

[**StorageAttachmentResponse**](StorageAttachment-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageAttachment**
> DeleteStorageAttachment($name)

Delete a Storage Attachment

Deletes the specified storage attachment. No response is returned.<p>Before deleting the storage attachment, you must unmount the associated storage volume. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&amp;id=STCSG-GUID-093D0406-28B9-4E8F-B650-57D5CDC56557\">Mounting and Unmounting a Storage Volume</a> in <i>Using Oracle Compute Cloud Service (IaaS)</i>.<p>Note that volumes attached to an instance at launch time can't be detached.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the object. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootStorageAttachment**
> StorageAttachmentDiscoverResponse DiscoverRootStorageAttachment()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**StorageAttachmentDiscoverResponse**](StorageAttachment-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverStorageAttachment**
> StorageAttachmentDiscoverResponse DiscoverStorageAttachment($container)

Retrieve Names of all Storage Attachments in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. | 

### Return type

[**StorageAttachmentDiscoverResponse**](StorageAttachment-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageAttachment**
> StorageAttachmentResponse GetStorageAttachment($name)

Retrieve Details of a Storage Attachment

Retrieves details of the specified storage attachment.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multipart name of the object. | 

### Return type

[**StorageAttachmentResponse**](StorageAttachment-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStorageAttachment**
> StorageAttachmentListResponse ListStorageAttachment($container, $instanceName, $state, $storageVolumeName, $name)

Retrieve Details of all Storage Attachments in a Container

Retrieves details of the storage attachments that are available in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the storage attachments in the container are displayed. To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt; or &lt;p&gt;&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;&lt;/code&gt; | 
 **instanceName** | **string**| Multipart name of an instance. | [optional] 
 **state** | **string**| &lt;p&gt;Specify one of the following states of the storage attachment:&lt;p&gt;&lt;code&gt;attaching&lt;/code&gt;: The storage attachment is in the process of attaching to the instance.&lt;p&gt;&lt;code&gt;attached&lt;/code&gt;: The storage attachment is attached to the instance.&lt;p&gt;&lt;code&gt;detaching&lt;/code&gt;: The storage attachment is in the process of detaching from the instance.&lt;p&gt;&lt;code&gt;unavailable&lt;/code&gt;: The storage attachment is unavailable.&lt;p&gt;&lt;code&gt;unknown&lt;/code&gt;: The state of the storage attachment is not known. | [optional] 
 **storageVolumeName** | **string**| &lt;p&gt;Three part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of a storage volume. | [optional] 
 **name** | **string**| &lt;p&gt;Multipart name of the object. | [optional] 

### Return type

[**StorageAttachmentListResponse**](StorageAttachment-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

