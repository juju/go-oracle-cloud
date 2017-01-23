# StorageAttachmentPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | &lt;p&gt;Index number for the volume. The allowed range is 1-10. The index determines the device name by which the volume is exposed to the instance. Index 0 is allocated to the temporary boot disk, /dev/xvda An attachment with index 1 is exposed to the instance as /dev/xvdb, an attachment with index 2 is exposed as /dev/xvdc, and so on. | [default to null]
**InstanceName** | **string** | Multipart name of the instance to which you want to attach the volume. | [default to null]
**StorageVolumeName** | **string** | &lt;p&gt;Three part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of the volume that you want to attach.&lt;p&gt;For information about how to create a storage volume, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-storage-volume--post.html\&quot;&gt;Create a Storage Volume&lt;/a&gt;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


