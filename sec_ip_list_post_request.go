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

// The request body contains details of the security IP list that you want to create. 
type SecIpListPostRequest struct {

	// <p>A description of the security IP list.
	Description string `json:"description,omitempty"`

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.
	Name string `json:"name,omitempty"`

	// <p>A comma-separated list of the subnets (in CIDR format) or IPv4 addresses for which you want to create this security IP list.<p>For example, to create a security IP list containing the IP addresses 203.0.113.1 and 203.0.113.2, enter one of the following:<p><code>203.0.113.0/30</code><p><code>203.0.113.1, 203.0.113.2</code>
	Secipentries []string `json:"secipentries,omitempty"`
}
