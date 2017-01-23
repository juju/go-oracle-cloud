# StorageSnapshotPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of this storage snapshot. | [optional] [default to null]
**Name** | **string** | Three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. If you do not specify a name for this object, the object name is generated automatically. | [optional] [default to null]
**Platform** | **string** | Specify the operating system platform for a bootable storage volume, such as Linux or Windows. | [optional] [default to null]
**Property** | **string** | &lt;ul&gt;&lt;li&gt;If you don&#39;t specify a value, a remote snapshot is created. Remote snapshots aren&#39;t stored in the same location as the original storage volume. Instead, they are reduced and stored in the associated Oracle Storage Cloud Service instance. Remote snapshots are useful if your domain spans multiple sites. With remote snapshots, you can create a snapshot in one site, then switch to another site and create a copy of the storage volume on that site. However, creating a remote snapshot and restoring a storage volume from a remote snapshot can take quite a long time depending on the size of the storage volume, as data is written to and from the Oracle Storage Cloud Service instance.&lt;/li&gt;&lt;li&gt;Specify &lt;code&gt;/oracle/private/storage/snapshot/collocated&lt;/code&gt; to create a collocated snapshot. Colocated snapshots are stored in the same physical location as the original storage volume and each snapshot uses the same amount of storage as the original volume. Colocated snapshots and volumes from colocated snapshots can be created very quickly. Colocated snapshots are useful for quickly cloning storage volumes within a site. However, you can&#39;t restore volumes across sites using colocated snapshots.&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**Tags** | **[]string** | Strings that describe the storage snapshot and help you identify it. | [optional] [default to null]
**Volume** | **string** | Three-part name of the storage volume (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of which you want to create a snapshot. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


