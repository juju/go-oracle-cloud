# IpReservationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Ip** | **string** | Public IP address.&lt;p&gt;An IP reservation is a public IP address that you can attach to an Oracle Compute Cloud Service instance that requires access to or from the Internet. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**Parentpool** | **string** | &lt;code&gt;/oracle/public/ippool&lt;/code&gt;&lt;p&gt;Pool of public IP addresses | [optional] [default to null]
**Permanent** | **bool** | &lt;code&gt;true&lt;/code&gt; indicates that the IP reservation has a persistent public IP address. You can associate either a temporary or a persistent public IP address with an instance when you create the instance.&lt;p&gt;Temporary public IP addresses are assigned dynamically from a pool of public IP addresses. When you associate a temporary public IP address with an instance, if the instance is restarted or is deleted and created again later, its public IP address might change. | [optional] [default to null]
**Quota** | **string** | Not used | [optional] [default to null]
**Tags** | **[]string** | A comma-separated list of strings which helps you to identify IP reservation. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**Used** | **bool** | &lt;code&gt;true&lt;/code&gt; indicates that the IP reservation is associated with an instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


