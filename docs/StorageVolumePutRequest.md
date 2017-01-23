# StorageVolumePutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | &lt;p&gt;The description of the storage volume. | [optional] [default to null]
**Name** | **string** | The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [default to null]
**Platform** | **string** | Specify the operating system platform for a bootable storage volume, such as Linux or Windows. | [optional] [default to null]
**Properties** | **[]string** | &lt;p&gt;The storage-pool property.&lt;p&gt;For storage volumes that require low latency and high IOPS, such as for storing database files, specify &lt;code&gt;/oracle/public/storage/latency&lt;/code&gt;.&lt;p&gt;For all other storage volumes, specify &lt;code&gt;/oracle/public/storage/default&lt;/code&gt;. | [default to null]
**Size** | **string** | &lt;p&gt;The size of this storage volume. Use one of the following abbreviations for the unit of measurement:&lt;ul&gt;&lt;li&gt;&lt;code&gt;B&lt;/code&gt; or &lt;code&gt;b&lt;/code&gt; for bytes&lt;/li&gt;&lt;li&gt;&lt;code&gt;K&lt;/code&gt; or &lt;code&gt;k&lt;/code&gt; for kilobytes&lt;/li&gt;&lt;li&gt;&lt;code&gt;M&lt;/code&gt; or &lt;code&gt;m&lt;/code&gt; for megabytes&lt;/li&gt;&lt;li&gt;&lt;code&gt;G&lt;/code&gt; or &lt;code&gt;g&lt;/code&gt; for gigabytes&lt;/li&gt;&lt;li&gt;&lt;code&gt;T&lt;/code&gt; or &lt;code&gt;t&lt;/code&gt; for terabytes&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;For example, to create a volume of size 10 gigabytes, you can specify &lt;code&gt;10G&lt;/code&gt;, or &lt;code&gt;10240M&lt;/code&gt;, or &lt;code&gt;10485760K&lt;/code&gt;, and so on.&lt;p&gt;The allowed range is from 1 GB to 2 TB, in increments of 1 GB. | [default to null]
**Snapshot** | **string** | Multipart name of the storage volume snapshot if this storage volume is a clone. | [optional] [default to null]
**SnapshotAccount** | **string** | Account of the parent snapshot from which the storage volume is restored. | [optional] [default to null]
**SnapshotId** | **string** | Id of the parent snapshot from which the storage volume is restored or cloned. | [optional] [default to null]
**Tags** | **[]string** | &lt;p&gt;Strings that you can use to tag the storage volume. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


