# VpnEndpointPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerVpnGateway** | **string** | Specify the IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway. Your gateway device must support route-based VPN and IKE (Internet Key Exchange) configuration using pre-shared keys. | [default to null]
**Enabled** | **bool** | Enables the VPN endpoint. To start a VPN connection, set to &lt;code&gt;true&lt;/code&gt;. A connection is established immediately, if possible. If you do not specify this option, the VPN endpoint is disabled and the connection is not established. | [optional] [default to null]
**Name** | **string** | Two-part name of the object (&lt;code&gt;&lt;em&gt;/Compute-acme/object&lt;/em&gt;&lt;/code&gt;). | [default to null]
**Psk** | **string** | Pre-shared VPN key. Enter the pre-shared key. This must be the same key that you provided when you requested the service. This secret key is shared between your network gateway and the Oracle Cloud network for authentication. Specify the full path and name of the text file that contains the pre-shared key. Ensure that the permission level of the text file is set to 400. The pre-shared VPN key must not exceed 256 characters. | [default to null]
**ReachableRoutes** | **[]string** | Specify a list of routes (CIDR prefixes) that are reachable through this VPN tunnel. You can specify a maximum of 20 IP subnet addresses. Specify IPv4 addresses in dot-decimal notation with or without mask. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


