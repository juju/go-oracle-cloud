/* 
 * REST API for Oracle Compute Cloud Service (IaaS)
 *
 * Use the Oracle Compute Cloud Service REST API to provision and manage instances and the associated resources
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// The request body contains details of the vpn endpoint that you want to update. 
type VpnEndpointPutRequest struct {

	// Specify the IP address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway. Your gateway device must support route-based VPN and IKE (Internet Key Exchange) configuration using pre-shared keys.
	CustomerVpnGateway string `json:"customer_vpn_gateway,omitempty"`

	// Enables the VPN endpoint. To start a VPN connection, set to <code>true</code>. A connection is established immediately, if possible. If you do not specify this option, the VPN endpoint is disabled and the connection is not established.
	Enabled bool `json:"enabled,omitempty"`

	// Two-part name of the object (<code><em>/Compute-acme/object</em></code>).
	Name string `json:"name,omitempty"`

	// Pre-shared VPN key. Enter the pre-shared key. This must be the same key that you provided when you requested the service. This secret key is shared between your network gateway and the Oracle Cloud network for authentication. Specify the full path and name of the text file that contains the pre-shared key. Ensure that the permission level of the text file is set to 400. The pre-shared VPN key must not exceed 256 characters.
	Psk string `json:"psk,omitempty"`

	// Specify a list of routes (CIDR prefixes) that are reachable through this VPN tunnel. You can specify a maximum of 20 IP subnet addresses. Specify IPv4 addresses in dot-decimal notation with or without mask.
	ReachableRoutes []string `json:"reachable_routes,omitempty"`
}
