# OrchestrationPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of this orchestration plan. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [default to null]
**Oplans** | [**[]Object**](object.md) | List of oplans. An object plan, or oplan, is a top-level orchestration attribute. | [default to null]
**Relationships** | **[]string** | A list of relationship specifications to be satisfied on this orchestration. | [optional] [default to null]
**Schedule** | [**Object**](object.md) | The schedule for an orchestration consists of the start and stop dates and times. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


