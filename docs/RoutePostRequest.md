# RoutePostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminDistance** | **int32** | Specify 0,1, or 2 as the route&#39;s administrative distance. If you do not specify a value, the default value is 0. &lt;p&gt;The same prefix can be used in multiple routes. In this case, packets are routed over all the matching routes with the lowest administrative distance. In the case multiple routes with the same lowest administrative distance match, routing occurs over all these routes using ECMP. | [optional] [default to null]
**IpAddressPrefix** | **string** | The IPv4 address prefix, in CIDR format, of the external network (external to the vNIC set) from which you want to route traffic. | [default to null]
**Name** | **string** | The three-part name (&lt;code&gt;/Compute-&lt;em&gt;identity_domain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;object&lt;/em&gt;&lt;/code&gt;) of the route. &lt;p&gt;Object names can contain only alphanumeric, underscore (_), dash (-), and period (.) characters. Object names are case-sensitive. | [default to null]
**NextHopVnicSet** | **string** | Name of the virtual NIC set to route matching packets to. Routed flows are load-balanced among all the virtual NICs in the virtual NIC set. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


