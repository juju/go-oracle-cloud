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

type SecApplicationResponse struct {

	// <p>A description of the security application.
	Description string `json:"description,omitempty"`

	// <p>The TCP or UDP destination port number.<p>You can also specify a port range, such as 5900-5999 for TCP.<p>If you specify <code>tcp</code> or <code>udp</code> as the protocol, then the <code>dport</code> parameter is required; otherwise, it is optional.<p>This parameter isn't relevant to the <code>icmp</code> protocol.<p><b>Note:</b> This request fails if the range-end is lower than the range-start. For example, if you specify the port range as 5000-4000.
	Dport string `json:"dport,omitempty"`

	// The ICMP code.<p>This parameter is relevant only if you specify <code>icmp</code> as the protocol. You can specify one of the following values:<ul><li><code>network</code></li><li><code>host</code></li><li><code>protocol</code></li><li><code>port</code></li><li><code>df</code></li><li><code>admin</code></li></ul><p>If you specify <code>icmp</code> as the protocol and don't specify <code>icmptype</code> or <code>icmpcode</code>, then all ICMP packets are matched.
	Icmpcode string `json:"icmpcode,omitempty"`

	// The ICMP type.<p>This parameter is relevant only if you specify <code>icmp</code> as the protocol. You can specify one of the following values:<ul><li><code>echo</code></li><li><code>reply</code></li><li><code>ttl</code></li><li><code>traceroute</code></li><li><code>unreachable</code></li></ul><p>If you specify <code>icmp</code> as the protocol and don't specify <code>icmptype</code> or <code>icmpcode</code>, then all ICMP packets are matched.
	Icmptype string `json:"icmptype,omitempty"`

	// The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// The protocol to use.<p>The value that you specify can be either a text representation of a protocol or any unsigned 8-bit assigned protocol number in the range 0-254. See <a target=\"_blank\" href=\"http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml\">Assigned Internet Protocol Numbers</a>.<p>For example, you can specify either tcp or the number 6.<p>The following text representations are allowed: <code>tcp</code>, <code>udp</code>, <code>icmp</code>, <code>igmp</code>, <ocde>ipip</code>, <code>rdp</code>, <code>esp</code>, <code>ah</code>, <code>gre</code>, <code>icmpv6</code>, <code>ospf</code>, <code>pim</code>, <code>sctp</code>, <code>mplsip</code>, <code>all</code>.<p>To specify all protocols, set this to <code>all</code>.
	Protocol string `json:"protocol,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`
}
