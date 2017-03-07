package response

// VpnEndpoint you can create secure IPSec-based tunnels between your data
// center and the instances in your Oracle Compute Cloud
// Service site to securely access your instances.
// A vpnendpoint object represents a VPN tunnel to
// your Oracle Compute Cloud Service site. You can
// create up to 20 VPN tunnels to your Oracle Compute
// Cloud Service site. You can use any internet service
// provider to access your Oracle Compute Cloud Service
// site, provided you have a VPN device to terminate an IPSec VPN tunnel

type VpnEndpoint struct {
}

type AllVpnEndpoints struct {
	Result []VpnEndpoint `json:"result,omitmepty"`
}
