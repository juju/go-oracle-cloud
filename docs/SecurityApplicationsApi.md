# \SecurityApplicationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecApplication**](SecurityApplicationsApi.md#AddSecApplication) | **Post** /secapplication/ | Create a Security Application
[**DeleteSecApplication**](SecurityApplicationsApi.md#DeleteSecApplication) | **Delete** /secapplication/{name} | Delete a Security Application
[**DiscoverRootSecApplication**](SecurityApplicationsApi.md#DiscoverRootSecApplication) | **Get** /secapplication/ | Retrieve Names of Containers
[**DiscoverSecApplication**](SecurityApplicationsApi.md#DiscoverSecApplication) | **Get** /secapplication/{container} | Retrieve Names of all Security Applications in a Container
[**GetSecApplication**](SecurityApplicationsApi.md#GetSecApplication) | **Get** /secapplication/{name} | Retrieve Details of a Security Application
[**ListSecApplication**](SecurityApplicationsApi.md#ListSecApplication) | **Get** /secapplication/{container}/ | Retrieve Details of all Security Applications in a Container


# **AddSecApplication**
> SecApplicationResponse AddSecApplication($body)

Create a Security Application

Creates a security application. After creating security applications, you can use them in security rules by using the HTTP request, POST /secrule/ <a class=\"xref\" href=\"op-secrule--post.html\">(Create a Security Rule)</a>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecApplicationPostRequest**](SecApplicationPostRequest.md)|  | 

### Return type

[**SecApplicationResponse**](SecApplication-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecApplication**
> DeleteSecApplication($name)

Delete a Security Application

Deletes a security application. No response is returned.<p>You can't delete system-provided security application that are available in the <code>/oracle/public</code> container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverRootSecApplication**
> SecApplicationDiscoverResponse DiscoverRootSecApplication()

Retrieve Names of Containers

Retrieves the names of containers that contain objects that you can access. You can use this information to construct the multipart name of an object.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SecApplicationDiscoverResponse**](SecApplication-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverSecApplication**
> SecApplicationDiscoverResponse DiscoverSecApplication($container)

Retrieve Names of all Security Applications in a Container

Retrieves the names of objects and subcontainers that you can access in the specified container.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;i&gt;user&lt;/i&gt;/&lt;/code&gt; to retrieve the names of objects that you can access. Specify &lt;code&gt;/Compute-&lt;i&gt;identityDomain&lt;/i&gt;/&lt;/code&gt; to retrieve the names of containers that contain objects that you can access. Specify &lt;code&gt;/oracle/public&lt;/code&gt; to retrieve the names of system-provided objects. | 

### Return type

[**SecApplicationDiscoverResponse**](SecApplication-discover-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+directory+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecApplication**
> SecApplicationResponse GetSecApplication($name)

Retrieve Details of a Security Application

<b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | 

### Return type

[**SecApplicationResponse**](SecApplication-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecApplication**
> SecApplicationListResponse ListSecApplication($container, $protocol, $description, $icmptype, $dport, $icmpcode, $name)

Retrieve Details of all Security Applications in a Container

Retrieves details of the security applications that are in the specified container and match the specified query criteria. If you don't specify any query criteria, then details of all the security applications in the container are displayed. You can use this HTTP request to validate the results of POST and PUT operations.<p>To filter the search results, you can pass one or more of the following query parameters, by appending them to the URI in the following syntax:<p><code>?parameter1=value1&ampparameter2=value2&ampparameterN=valueN</code><p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Monitor</code> or <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **string**| &lt;code&gt;/Compute-identity_domain/user&lt;/code&gt; or &lt;code&gt;/Compute-identity_domain&lt;/code&gt; for user-created security applications and &lt;code&gt;/oracle/public&lt;/code&gt; for system-provided security applications. | 
 **protocol** | **string**| The protocol to use.&lt;p&gt;The value that you specify can be either a text representation of a protocol or any unsigned 8-bit assigned protocol number in the range 0-254. See &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml\&quot;&gt;Assigned Internet Protocol Numbers&lt;/a&gt;.&lt;p&gt;For example, you can specify either tcp or the number 6.&lt;p&gt;The following text representations are allowed: &lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, &lt;code&gt;icmp&lt;/code&gt;, &lt;code&gt;igmp&lt;/code&gt;, &lt;ocde&gt;ipip&lt;/code&gt;, &lt;code&gt;rdp&lt;/code&gt;, &lt;code&gt;esp&lt;/code&gt;, &lt;code&gt;ah&lt;/code&gt;, &lt;code&gt;gre&lt;/code&gt;, &lt;code&gt;icmpv6&lt;/code&gt;, &lt;code&gt;ospf&lt;/code&gt;, &lt;code&gt;pim&lt;/code&gt;, &lt;code&gt;sctp&lt;/code&gt;, &lt;code&gt;mplsip&lt;/code&gt;, &lt;code&gt;all&lt;/code&gt;.&lt;p&gt;To specify all protocols, set this to &lt;code&gt;all&lt;/code&gt;. | [optional] 
 **description** | **string**| &lt;p&gt;A description of the security application. | [optional] 
 **icmptype** | **string**| The ICMP type.&lt;p&gt;This parameter is relevant only if you specify &lt;code&gt;icmp&lt;/code&gt; as the protocol. You can specify one of the following values:&lt;ul&gt;&lt;li&gt;&lt;code&gt;echo&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;reply&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;ttl&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;traceroute&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;unreachable&lt;/code&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;If you specify &lt;code&gt;icmp&lt;/code&gt; as the protocol and don&#39;t specify &lt;code&gt;icmptype&lt;/code&gt; or &lt;code&gt;icmpcode&lt;/code&gt;, then all ICMP packets are matched. | [optional] 
 **dport** | **string**| &lt;p&gt;The TCP or UDP destination port number.&lt;p&gt;You can also specify a port range, such as 5900-5999 for TCP.&lt;p&gt;If you specify &lt;code&gt;tcp&lt;/code&gt; or &lt;code&gt;udp&lt;/code&gt; as the protocol, then the &lt;code&gt;dport&lt;/code&gt; parameter is required; otherwise, it is optional.&lt;p&gt;This parameter isn&#39;t relevant to the &lt;code&gt;icmp&lt;/code&gt; protocol.&lt;p&gt;&lt;b&gt;Note:&lt;/b&gt; This request fails if the range-end is lower than the range-start. For example, if you specify the port range as 5000-4000. | [optional] 
 **icmpcode** | **string**| The ICMP code.&lt;p&gt;This parameter is relevant only if you specify &lt;code&gt;icmp&lt;/code&gt; as the protocol. You can specify one of the following values:&lt;ul&gt;&lt;li&gt;&lt;code&gt;network&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;host&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;protocol&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;port&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;df&lt;/code&gt;&lt;/li&gt;&lt;li&gt;&lt;code&gt;admin&lt;/code&gt;&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;If you specify &lt;code&gt;icmp&lt;/code&gt; as the protocol and don&#39;t specify &lt;code&gt;icmptype&lt;/code&gt; or &lt;code&gt;icmpcode&lt;/code&gt;, then all ICMP packets are matched. | [optional] 
 **name** | **string**| The three-part name of the object (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;). | [optional] 

### Return type

[**SecApplicationListResponse**](SecApplication-list-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

