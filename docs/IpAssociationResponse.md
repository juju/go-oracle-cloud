# IpAssociationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Ip** | **string** | The public IP address which is attached to an Oracle Compute Cloud Service instance that requires access to or from the Internet. | [optional] [default to null]
**Name** | **string** | The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**Parentpool** | **string** | &lt;ul&gt;&lt;li&gt;To associate a temporary IP address from the pool, specify ippool:/oracle/public/ippool.&lt;/li&gt;&lt;li&gt;To associate a persistent IP address, specify ipreservation:ipreservation_name, where ipreservation_name is three-part name of an existing IP reservation in the &lt;code&gt;/Compute-identity_domain/user/object_name&lt;/code&gt; format. For more information about how to create an IP reservation, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-ip-reservation--post.html\&quot;&gt;Create an IP Reservation&lt;/a&gt;.&lt;/li&gt;&lt;ul&gt; | [optional] [default to null]
**Reservation** | **string** | The three-part name of the IP reservation object in the format (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). An IP reservation is a public IP address which is attached to an Oracle Compute Cloud Service instance that requires access to or from the Internet. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**Vcable** | **string** | The three-part name of a vcable ID of an instance that is associated with the IP reservation. The three-part name is in the format: &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


