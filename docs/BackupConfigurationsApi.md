# \BackupConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupConfiguration**](BackupConfigurationsApi.md#CreateBackupConfiguration) | **Post** /backupservice/v1/configuration | Create a Backup Configuration
[**DeleteBackupConfiguration**](BackupConfigurationsApi.md#DeleteBackupConfiguration) | **Delete** /backupservice/v1/configuration/{name} | Delete a Backup Configuration
[**GetBackupConfiguration**](BackupConfigurationsApi.md#GetBackupConfiguration) | **Get** /backupservice/v1/configuration/{name} | Retrieve Details of a Backup Configuration
[**ListBackupConfigurations**](BackupConfigurationsApi.md#ListBackupConfigurations) | **Get** /backupservice/v1/configuration | Retrieve Details of All Backup Configurations
[**UpdateBackupConfiguration**](BackupConfigurationsApi.md#UpdateBackupConfiguration) | **Put** /backupservice/v1/configuration/{name} | Update a Backup Configuration


# **CreateBackupConfiguration**
> BackupConfiguration CreateBackupConfiguration($body)

Create a Backup Configuration

Create a new backup configuration. Requires authorization to create backup configurations as well as appropriate authorization to create snapshots from the target volume.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBackupConfigurationInput**](CreateBackupConfigurationInput.md)| BackupConfiguration to create | 

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupConfiguration**
> DeleteBackupConfiguration($name)

Delete a Backup Configuration

Delete a backup configuration. In order to delete the configuration all backups and restores related to the configuration must already be deleted. If disabling a backup configuration is desired, consider setting <code>enabled</code> to false.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of BackupConfiguration to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupConfiguration**
> BackupConfiguration GetBackupConfiguration($name)

Retrieve Details of a Backup Configuration

Retrieves details of the specified backup configuration. You can use this request to verify whether the POST and PUT HTTP requests were completed successfully.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Multi-part name of BackupConfiguration to get | 

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupConfigurations**
> []BackupConfiguration ListBackupConfigurations()

Retrieve Details of All Backup Configurations

Retrieve details for all backup configuration objects the current user has permission to access.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code>role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]BackupConfiguration**](BackupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBackupConfiguration**
> UpdateBackupConfiguration($body, $name)

Update a Backup Configuration

Modify an existing backup configuration. All fields, including unmodifiable fields, must be provided for this operation. The following fields are unmodifiable: volumeUri, runAsUser, name.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateBackupConfigurationInput**](UpdateBackupConfigurationInput.md)| BackupConfiguration to update | 
 **name** | **string**| Multi-part name of BackupConfiguration to update | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json, application/json
 - **Accept**: application/oracle-compute-v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

