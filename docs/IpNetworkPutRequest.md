# IpNetworkPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressPrefix** | **string** | Range of IPv4 addresses assigned in the virtual network, in CIDR address prefix format.&lt;p&gt;You can expand the network by specifying the same IP address prefix but with a shorter prefix length. For example, you can expand 192.168.1.0/24 to 192.168.1.0/20. Don&#39;t, however, change the IP address. This ensures that all IP addresses that have been currently allocated to instances remain valid in the updated IP network. | [default to null]
**IpNetworkExchange** | **string** | Specify the three-part name of the IP network exchange to which you want to add the IP network. You can add an IP network to only one IP network exchange, but an IP network exchange can include multiple IP networks. An IP network exchange enables access between IP networks that have non-overlapping addresses, so that instances on these networks can exchange packets with each other without NAT.  | [optional] [default to null]
**Name** | **string** | Specify the three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network that you want to update. &lt;p&gt;Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


