# SecRulePostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | &lt;p&gt;Set this parameter to &lt;code&gt;PERMIT&lt;/code&gt;. | [default to null]
**Application** | **string** | &lt;p&gt;The three-part name of the security application: (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) for user-defined security applications and &lt;code&gt;/oracle/public/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt; for predefined security applications. | [default to null]
**Description** | **string** | &lt;p&gt;A description of the security rule. | [optional] [default to null]
**Disabled** | **bool** | &lt;p&gt;Indicates whether the security rule is enabled (set to &lt;code&gt;false&lt;/code&gt;) or disabled (&lt;code&gt;true&lt;/code&gt;). The default setting is &lt;code&gt;false&lt;/code&gt;. | [optional] [default to null]
**DstList** | **string** | &lt;p&gt;The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of the destination security list or security IP list.&lt;p&gt;You must use the prefix &lt;code&gt;seclist&lt;/code&gt;: or &lt;code&gt;seciplist&lt;/code&gt;: to identify the list type.&lt;p&gt;&lt;b&gt;Note:&lt;/b&gt; You can specify a security IP list as the destination in a secrule, provided &lt;code&gt;src_list&lt;/code&gt; is a security list that has DENY as its outbound policy.&lt;p&gt;You cannot specify any of the security IP lists in the &lt;code&gt;/oracle/public&lt;/code&gt; container as a destination in a secrule. | [default to null]
**Name** | **string** | &lt;p&gt;The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;).&lt;p&gt;Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive. | [default to null]
**SrcList** | **string** | &lt;p&gt;The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object_name&lt;/em&gt;&lt;/code&gt;) of the source security list or security IP list.&lt;p&gt;You must use the prefix &lt;code&gt;seclist&lt;/code&gt;: or &lt;code&gt;seciplist&lt;/code&gt;: to identify the list type. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


