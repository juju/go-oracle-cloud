# IpAssociationPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parentpool** | **string** | &lt;ul&gt;&lt;li&gt;To associate a temporary IP address from the pool, specify ippool:/oracle/public/ippool.&lt;/li&gt;&lt;li&gt;To associate a persistent IP address, specify ipreservation:ipreservation_name, where ipreservation_name is three-part name of an existing IP reservation in the &lt;code&gt;/Compute-identity_domain/user/object_name&lt;/code&gt; format. For more information about how to create an IP reservation, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-ip-reservation--post.html\&quot;&gt;Create an IP Reservation&lt;/a&gt;.&lt;/li&gt;&lt;ul&gt; | [default to null]
**Vcable** | **string** | &lt;p&gt;The three-part name of the vcable ID of the instance that you want to associate with an IP address. The three-part name is in the format: &lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;.&lt;p&gt;For more information about the vcable of an instance, see &lt;a class&#x3D;\&quot;xref\&quot; href&#x3D;\&quot;op-instance-{name}-get.html\&quot;&gt;Retrieve Details of an Instance&lt;/a&gt;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


