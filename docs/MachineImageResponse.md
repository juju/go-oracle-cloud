# MachineImageResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Two-part name of the account (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/cloud_storage&lt;/code&gt; that contains the credentials and access details of the associated Oracle Storage Cloud Service instance. | [optional] [default to null]
**Attributes** | [**Object**](object.md) | &lt;p&gt;An optional JSON object or dictionary of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-268FE284-E5A0-4A18-BA58-345660925FB7\&quot;&gt;Retrieving User-Defined Instance Attributes&lt;/a&gt; in &lt;em&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/em&gt;. | [optional] [default to null]
**Audited** | **string** | Last time when this image was audited | [optional] [default to null]
**Checksums** | [**Object**](object.md) | Not used | [optional] [default to null]
**ErrorReason** | **string** | Description of the state of the machine image if there is an error. | [optional] [default to null]
**File** | **string** | Name of the machine image file. | [optional] [default to null]
**Hypervisor** | [**Object**](object.md) | A dictionary of hypervisor-specific attributes. | [optional] [default to null]
**ImageFormat** | **string** | The format of the image. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object. | [optional] [default to null]
**NoUpload** | **bool** | &lt;code&gt;true&lt;/code&gt; indicates that the image file is available in Oracle Cloud Storage Service. | [optional] [default to null]
**Platform** | **string** | The OS platform of the image. | [optional] [default to null]
**Quota** | **string** | Not used | [optional] [default to null]
**Signature** | **string** | Not used | [optional] [default to null]
**SignedBy** | **string** | Not used | [optional] [default to null]
**Sizes** | [**Object**](object.md) | Size values of the image file. | [optional] [default to null]
**State** | **string** | The state of the uploaded machine image. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


