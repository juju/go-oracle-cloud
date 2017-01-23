# CreateBackupConfigurationInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of this Backup Configuration | [optional] [default to null]
**Interval** | [**Interval**](Interval.md) | Backup scheduler interval | [default to null]
**Enabled** | **bool** | When true, backups will automatically be generated based on the interval. | [optional] [default to null]
**BackupRetentionCount** | **int32** | How many backups to retain | [optional] [default to null]
**Name** | **string** | The multi-part name of the object. Object names can contain only alphanumeric characters, hyphens, underscores, @ and periods. Object names are case-sensitive. | [default to null]
**VolumeUri** | **string** | The complete URI of the storage volume that you want to backup. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


