# MachineImagePostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Two-part name of the account (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/cloud_storage&lt;/code&gt; that contains the credentials and access details of the associated Oracle Storage Cloud Service instance. | [default to null]
**Attributes** | [**Object**](object.md) | &lt;p&gt;An optional JSON object or dictionary of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;STCSG-GUID-268FE284-E5A0-4A18-BA58-345660925FB7\&quot;&gt;Retrieving User-Defined Instance Attributes&lt;/a&gt; in &lt;em&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/em&gt;. | [optional] [default to null]
**File** | **string** | &lt;code&gt;&lt;em&gt;image_file&lt;/em&gt;.tar.gz&lt;/code&gt;, where &lt;code&gt;&lt;em&gt;image_file&lt;/em&gt;&lt;/code&gt; is the .tar.gz name of the machine image file that you have uploaded to Oracle Storage Cloud Service. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.oracle.com/pls/topic/lookup?ctx&#x3D;stcomputecs&amp;id&#x3D;MMOCS-GUID-799D6F6D-BDED-4DDE-9B3D-BE23BE5F687F\&quot;&gt;Uploading Machine Image Files to Oracle Storage Cloud Service&lt;/a&gt; in &lt;em&gt;Using Oracle Compute Cloud Service (IaaS)&lt;/em&gt;. | [optional] [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. | [default to null]
**NoUpload** | **bool** | &lt;p&gt;Set this to &lt;code&gt;true&lt;/code&gt;. Indicates that the image file is available in Oracle Cloud Storage Service. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


