# SnapshotPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | **string** | Use this option when you want to take preserve the custom changes you have made to an instance before deleting the instance. The only permitted value is shutdown. Snapshot of the instance is not taken immediately. It creates a machine image which preserves the changes you have made to the instance, and then the instance is deleted.&lt;p&gt;&lt;b&gt;Note:&lt;/b&gt; This option has no effect if you shutdown the instance from inside it. Any pending snapshot request on that instance goes into error state. You must delete the instance (&lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-instance-{name}-delete.html\&quot;&gt;DELETE /instance/{name}&lt;/a&gt;). | [optional] [default to null]
**Instance** | **string** | Multipart name of the instance that you want to clone. | [default to null]
**Machineimage** | **string** | Specify the three-part name (&lt;code&gt;/Compute-identity_domain/user/object&lt;/code&gt;) of the machine image created by the snapshot request.&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.&lt;p&gt;If you don&#39;t specify a name for this object, then the name is generated automatically. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


