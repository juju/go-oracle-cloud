# OrchestrationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Shows the default account for your identity domain. | [optional] [default to null]
**Description** | **string** | Description of this orchestration plan. | [optional] [default to null]
**Info** | [**Object**](object.md) | The nested parameter &lt;code&gt;errors&lt;/code&gt; shows which object in the orchestration has encountered an error. Empty if there are no errors. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] [default to null]
**Oplans** | [**[]Object**](object.md) | List of oplans. An object plan, or oplan, is a top-level orchestration attribute. | [optional] [default to null]
**Relationships** | **[]string** | The relationship between the objects that are created by this orchestration. | [optional] [default to null]
**Schedule** | [**Object**](object.md) | The schedule for an orchestration consists of the start and stop dates and times. | [optional] [default to null]
**Status** | **string** | Shows the current status of the orchestration. | [optional] [default to null]
**StatusTimestamp** | **string** | This information is generally displayed at the end of the orchestration JSON. It indicates the time that the current view of the orchestration was generated. This information shows only when the orchestration is running. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**User** | **string** | Two-part name of the user (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;&lt;/code&gt;) who has most recently taken an action on the orchestration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


