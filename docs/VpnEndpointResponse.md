# VpnEndpointResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerVpnGateway** | **string** | IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway. | [optional] [default to null]
**Enabled** | **bool** | &lt;code&gt;true&lt;/code&gt; indicates that the VPN endpoint is enabled and a connection is established immediately, if possible. | [optional] [default to null]
**Name** | **string** | Two-part name of the object (&lt;code&gt;&lt;em&gt;/Compute-acme/object&lt;/em&gt;&lt;/code&gt;) | [optional] [default to null]
**Psk** | **string** | Pre-shared VPN key. | [optional] [default to null]
**ReachableRoutes** | **[]string** | List of routes (CIDR prefixes) that are reachable through this VPN tunnel. | [optional] [default to null]
**Status** | **string** | Current status of the VPN tunnel. | [optional] [default to null]
**StatusDesc** | **string** | Describes the current status of the VPN tunnel. | [optional] [default to null]
**Uri** | **string** | Uniform Resource Identifier | [optional] [default to null]
**VpnStatistics** | [**Object**](object.md) | Statistics of VPN tunnels | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


