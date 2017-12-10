/* 
 * Telstra Programmable Network API
 *
 * Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.
 *
 * OpenAPI spec version: 2.1.2
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package TelstraTPN

type Vport struct {

	// 
	Status string `json:"status,omitempty"`

	// 
	Bandwidth int32 `json:"bandwidth,omitempty"`

	// 
	Interfacesequence int32 `json:"interfacesequence,omitempty"`

	// 
	Vporttype string `json:"vporttype,omitempty"`

	// 
	Vportuuid string `json:"vportuuid,omitempty"`

	Vportvalue Vportvalue `json:"vportvalue,omitempty"`
}
