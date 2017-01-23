# \RestoresApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestore**](RestoresApi.md#CreateRestore) | **Post** /backupservice/v1/restore | Create a Restore
[**DeleteRestore**](RestoresApi.md#DeleteRestore) | **Delete** /backupservice/v1/restore/{name} | Delete Restore
[**GetRestore**](RestoresApi.md#GetRestore) | **Get** /backupservice/v1/restore/{name} | Retrieve Details of a Restore
[**ListRestores**](RestoresApi.md#ListRestores) | **Get** /backupservice/v1/restore | Retrieve Details of all Restores


# **CreateRestore**
> Restore CreateRestore($body)

Create a Restore

Restores a storage volume from a previously completed backup. Before sending a request to restore, verify that the backup to be restored is in the <code>completed</code> state by retrieving details of the backup and checking the <code>state</code> field. You can only restore backups that are in the Completed state.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RestoreInput**](RestoreInput.md)| Restore to execute | 

### Return type

[**Restore**](Restore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRestore**
> DeleteRestore($name)

Delete Restore

Deletes Restore model only. Restored volume will remain.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of the Restore to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRestore**
> Restore GetRestore($name)

Retrieve Details of a Restore

Returns a json object of restore object, requires /restore/<PathName>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of the Restore you want to get | 

### Return type

[**Restore**](Restore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRestores**
> []Restore ListRestores($backupConfigurationName)

Retrieve Details of all Restores

Returns a json object of restore objects.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupConfigurationName** | **string**| Multi-part name of BackupConfiguration. Filters Restores by BackupConfiguration | [optional] 

### Return type

[**[]Restore**](Restore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

