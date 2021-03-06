/*
 * Telstra Programmable Network API
 *
 * Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.
 *
 * API version: 2.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TelstraTPN

type EndpointPort struct {

	// 
	Cabletype string `json:"cabletype,omitempty"`

	// 
	Connectortype string `json:"connectortype,omitempty"`

	// 
	Cfa string `json:"cfa,omitempty"`

	// 
	Endpointuuid string `json:"endpointuuid,omitempty"`

	// 
	Interfacespeed string `json:"interfacespeed,omitempty"`

	// 
	Interfacetype string `json:"interfacetype,omitempty"`

	// 
	Vport []Vport `json:"vport,omitempty"`
}
