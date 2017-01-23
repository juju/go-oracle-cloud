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

// The request body contains details of the ssh key that you want to create. 
type SshKeyPostRequest struct {

	// <p>Indicates whether the key must be enabled (default) or disabled. Note that disabled keys cannot be associated with instances.<p>To explicitly enable the key, specify <code>true</code>. To disable the key, specify <code>false</code>.
	Enabled bool `json:"enabled,omitempty"`

	// <p>The SSH public key value.
	Key string `json:"key,omitempty"`

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.
	Name string `json:"name,omitempty"`
}
