# LaunchPlanResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]Object**](object.md) | &lt;p&gt;A list of instances.&lt;p&gt;Each instance is defined using the instance attributes. | [optional] [default to null]
**Relationships** | **[]string** | &lt;p&gt;The relationships between various instances.&lt;p&gt;Valid values:&lt;ul&gt;&lt;li&gt;same_node: The specified instances are created on the same physical server. This is useful if you want to ensure low latency across instances.&lt;/li&gt;&lt;li&gt;different_node: The specified instances aren&#39;t created on the same physical server. This is useful if you want to isolate instances for security or redundancy.&lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


