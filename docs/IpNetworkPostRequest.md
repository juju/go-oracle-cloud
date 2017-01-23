# IpNetworkPostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressPrefix** | **string** | Specify the size of the IP subnet. It is a range of IPv4 addresses assigned in the virtual network, in CIDR address prefix format.&lt;p&gt;While specifying the IP address prefix take care of the following points:&lt;p&gt;* These IP addresses aren&#39;t part of the common pool of Oracle-provided IP addresses used by the shared network.&lt;p&gt;* There&#39;s no conflict with the range of IP addresses used in another IP network, the IP addresses used your on-premises network, or with the range of private IP addresses used in the shared network. If IP networks with overlapping IP subnets are linked to an IP exchange, packets going to and from those IP networks are dropped.&lt;p&gt;* The upper limit of the CIDR block size for an IP network is /16.&lt;p&gt;Note: The first IP address of any IP network is reserved for the default gateway, the DHCP server, and the DNS server of that IP network.  | [default to null]
**IpNetworkExchange** | **string** | Specify the IP network exchange to which the IP network belongs. You can add an IP network to only one IP network exchange, but an IP network exchange can include multiple IP networks. An IP network exchange enables access between IP networks that have non-overlapping addresses, so that instances on these networks can exchange packets with each other without NAT.  | [optional] [default to null]
**Name** | **string** | The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the IP network. &lt;p&gt;Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


