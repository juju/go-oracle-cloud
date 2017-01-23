# SecListResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Description** | **string** | &lt;p&gt;A description of the security list. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**OutboundCidrPolicy** | **string** | &lt;p&gt;The policy for outbound traffic from the security list. You can specify one of the following values:&lt;ul&gt;&lt;li&gt;&lt;code&gt;deny&lt;/code&gt;: Packets are dropped. No response is sent.&lt;/li&gt;&lt;li&gt;&lt;code&gt;reject&lt;/code&gt;: Packets are dropped, but a response is sent.&lt;/li&gt;&lt;li&gt;&lt;code&gt;permit&lt;/code&gt;(default): Packets are allowed.&lt;/li&gt; | [optional] [default to null]
**Policy** | **string** | &lt;p&gt;The policy for inbound traffic to the security list. You can specify one of the following values:&lt;ul&gt;&lt;li&gt;&lt;code&gt;deny&lt;/code&gt;(default): Packets are dropped. No response is sent.&lt;/li&gt;&lt;li&gt;&lt;code&gt;reject&lt;/code&gt;: Packets are dropped, but a response is sent.&lt;/li&gt;&lt;li&gt;&lt;code&gt;permit&lt;/code&gt;: Packets are allowed. This policy effectively turns off the firewall for all instances in this security list.&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


