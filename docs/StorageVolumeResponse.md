# StorageVolumeResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Bootable** | **bool** | &lt;code&gt;true&lt;/code&gt; indicates that the storage volume can also be used as a boot disk for an instance. | [optional] [default to null]
**Description** | **string** | &lt;p&gt;The description of the storage volume. | [optional] [default to null]
**Hypervisor** | **string** | The hypervisor that this volume is compatible with. | [optional] [default to null]
**Imagelist** | **string** | Name of machine image to extract onto this volume when created. This information is provided only for bootable storage volumes. | [optional] [default to null]
**ImagelistEntry** | **int32** | Specific imagelist entry version to extract. | [optional] [default to null]
**MachineimageName** | **string** | Three-part name of the machine image. This information is available if the volume is a bootable storage volume. | [optional] [default to null]
**Managed** | **bool** | All volumes are managed volumes. Default value is &lt;code&gt;true&lt;/code&gt;. | [optional] [default to null]
**Name** | **string** | The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**Platform** | **string** | The OS platform this volume is compatible with. | [optional] [default to null]
**Properties** | **[]string** | The storage-pool property: &lt;code&gt;/oracle/public/storage/latency&lt;/code&gt; or &lt;code&gt;/oracle/public/storage/default&lt;/code&gt;. | [optional] [default to null]
**Quota** | **string** | Not used | [optional] [default to null]
**Shared** | **bool** | Not used | [optional] [default to null]
**Size** | **string** | The size of this storage volume. | [optional] [default to null]
**Snapshot** | **string** | Name of the parent snapshot from which the storage volume is restored or cloned. | [optional] [default to null]
**SnapshotAccount** | **string** | Account of the parent snapshot from which the storage volume is restored. | [optional] [default to null]
**SnapshotId** | **string** | Id of the parent snapshot from which the storage volume is restored or cloned. | [optional] [default to null]
**Status** | **string** | The current state of the storage volume. | [optional] [default to null]
**StatusDetail** | **string** | Details about the latest state of the storage volume. | [optional] [default to null]
**StatusTimestamp** | **string** | It indicates the time that the current view of the storage volume was generated. | [optional] [default to null]
**StoragePool** | **string** | The storage pool from which this volume is allocated. | [optional] [default to null]
**Tags** | **[]string** | Comma-separated strings that tag the storage volume. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**Writecache** | **bool** | Not used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


