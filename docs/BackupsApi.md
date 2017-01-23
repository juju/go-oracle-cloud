# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupsApi.md#CreateBackup) | **Post** /backupservice/v1/backup | Create a Backup
[**DeleteBackup**](BackupsApi.md#DeleteBackup) | **Delete** /backupservice/v1/backup/{name} | Delete a Backup
[**GetBackupByName**](BackupsApi.md#GetBackupByName) | **Get** /backupservice/v1/backup/{name} | Retrieves Details of the Specified Backup
[**ListBackups**](BackupsApi.md#ListBackups) | **Get** /backupservice/v1/backup | Retrieve Details of all Backups


# **CreateBackup**
> Backup CreateBackup($body)

Create a Backup

Schedules a backup immediately using the specified backup configuration. The storage volume that you have specified in the backup configuration is backed up immediately, irrespective of the status of <code>enabled</code> in the specified backup configuration.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BackupInput**](BackupInput.md)| Multi-part name of BackupConfiguration to execute | 

### Return type

[**Backup**](Backup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackup**
> DeleteBackup($name)

Delete a Backup

Delete a backup and it's associated snapshot. In progress backups may not be deleted.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of backup to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupByName**
> Backup GetBackupByName($name)

Retrieves Details of the Specified Backup

Get the backup specified by the provided multi-part object name.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of Backup to get | 

### Return type

[**Backup**](Backup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackups**
> []Backup ListBackups($backupConfigurationName)

Retrieve Details of all Backups

Retrieves details of the backups that are available and match the specified query criteria. If you don't specify any query criteria, then details of all the backups in the container are displayed. To filter the search results, you can pass one or more of the documented query parameters by appending them to the URI in the following syntax: <p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupConfigurationName** | **string**| Multi-part name of BackupConfiguration. Filters Backups by BackupConfiguration | [optional] 

### Return type

[**[]Backup**](Backup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

