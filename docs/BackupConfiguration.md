# BackupConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextScheduledRun** | [**time.Time**](time.Time.md) | Scheduled time for next backup execution | [optional] [default to null]
**Name** | **string** | The multi-part name of the object. Object names can contain only alphanumeric characters, hyphens, underscores, @ and periods. Object names are case-sensitive. | [optional] [default to null]
**RunAsUser** | **string** | Any actions on this model will be performed as this user. | [optional] [default to null]
**Description** | **string** | Description of this Backup Configuration | [optional] [default to null]
**Interval** | [**Interval**](Interval.md) | Backup scheduler interval | [optional] [default to null]
**Enabled** | **bool** | When true, backups will automatically be generated based on the interval. | [optional] [default to null]
**BackupRetentionCount** | **int32** | How many backups to retain | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**TagId** | [**Uuid**](UUID.md) | ID used to tag other cloud resources | [optional] [default to null]
**VolumeUri** | **string** | The complete URI of the storage volume that you want to backup. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


