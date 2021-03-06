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

type Params29 struct {

	// 
	Description string `json:"description,omitempty"`

	// Duration of contract in minutes
	Duration int32 `json:"duration,omitempty"`

	// Bandwidth in kB/s
	Bandwidth int32 `json:"bandwidth,omitempty"`

	// Latency: 0=Low, 1=Standard, 2=Best Effort
	Latency int32 `json:"latency,omitempty"`

	// Identifier of a link
	Linkid string `json:"linkid,omitempty"`

	// Identifier of a contract
	Contractid string `json:"contractid,omitempty"`

	// 
	Price int32 `json:"price,omitempty"`

	// 
	ContractStatus int32 `json:"contractStatus,omitempty"`

	// 
	Version int32 `json:"version,omitempty"`

	// 
	Deletedtimestamp int32 `json:"deletedtimestamp,omitempty"`

	// 
	Tag string `json:"tag,omitempty"`

	// 
	Connection []string `json:"connection,omitempty"`

	// 
	Type_ string `json:"type,omitempty"`

	// 
	BillingId string `json:"billing-id,omitempty"`

	// Renewal Option: 0=Auto Disconnect, 1=Auto Renew, 2=Pay per hour
	RenewalOption int32 `json:"renewal-option,omitempty"`

	// 
	ContractStartTime int32 `json:"contract-start-time,omitempty"`

	// 
	ContractEndTime int32 `json:"contract-end-time,omitempty"`
}
