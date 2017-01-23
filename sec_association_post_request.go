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

// The request body contains details of the security association that you want to create. 
type SecAssociationPostRequest struct {

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).<p>If you don't specify a name for this object, then the name is generated automatically.<p>Object names can contain only alphanumeric characters, hyphens, underscores, and periods. Object names are case-sensitive.
	Name string `json:"name,omitempty"`

	// <p>Security list that you want to associate with the instance.
	Seclist string `json:"seclist,omitempty"`

	// <p>vcable of the instance that you want to associate with the security list.<p>For more information about the vcable of an instance, see <a class=\"xref\" href=\"op-instance-{name}-get.html\">Retrieve Details of an Instance.</a>
	Vcable string `json:"vcable,omitempty"`
}
