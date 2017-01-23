# Restore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeUri** | **string** | The target volume created from this restore job | [optional] [default to null]
**RunAsUser** | **string** | Any actions on this model will be performed as this user. | [optional] [default to null]
**Bootable** | **bool** | True if the restored volume is bootable | [optional] [default to null]
**Description** | **string** | Description of the Restore | [optional] [default to null]
**ErrorMessage** | **string** | Human readable error message | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**Name** | **string** | The multi-part name of the object. Object names can contain only alphanumeric characters, hyphens, underscores, @ and periods. Object names are case-sensitive. | [optional] [default to null]
**BackupConfigurationName** | **string** | Multi-part name of the backup configuration. | [optional] [default to null]
**State** | **string** | State of this resource. | [optional] [default to null]
**DetailedErrorMessage** | **string** | Human readable detailed error message | [optional] [default to null]
**TagId** | [**Uuid**](UUID.md) | ID used to tag other cloud resources | [optional] [default to null]
**BackupName** | **string** | Multi-part name of the backup that you want to restore. The backup must be in the &lt;code&gt;completed&lt;/code&gt; state. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


