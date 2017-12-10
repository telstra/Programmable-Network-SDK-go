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

type Object50 struct {

	// 
	BriefDescription string `json:"brief_description,omitempty"`

	// 
	Buyers string `json:"buyers,omitempty"`

	// 
	Classifications []Classification `json:"classifications,omitempty"`

	// 
	Comments string `json:"comments,omitempty"`

	// 
	CreateAt int32 `json:"create_at,omitempty"`

	// 
	Creator string `json:"creator,omitempty"`

	// 
	Description string `json:"description,omitempty"`

	// 
	Eos string `json:"eos,omitempty"`

	// 
	Flavors []Flavor `json:"flavors,omitempty"`

	// 
	FlavorsPrice string `json:"flavors_price,omitempty"`

	// 
	GlanceComments string `json:"glance_comments,omitempty"`

	GlanceImage GlanceImage `json:"glance_image,omitempty"`

	// 
	GlanceName string `json:"glance_name,omitempty"`

	// 
	GlanceProperties string `json:"glance_properties,omitempty"`

	// 
	Id int32 `json:"id,omitempty"`

	// 
	ImageFormat string `json:"image_format,omitempty"`

	// 
	LicenseRequired bool `json:"license_required,omitempty"`

	// 
	Logo string `json:"logo,omitempty"`

	// 
	MaxPorts int32 `json:"max_ports,omitempty"`

	// 
	Md5 string `json:"md5,omitempty"`

	// 
	MinPorts int32 `json:"min_ports,omitempty"`

	// 
	Name string `json:"name,omitempty"`

	// 
	OsVersion string `json:"os_version,omitempty"`

	// 
	Owner string `json:"owner,omitempty"`

	Product Product `json:"product,omitempty"`

	// 
	PublishDate string `json:"publish_date,omitempty"`

	// 
	RestrictVncConsole bool `json:"restrict_vnc_console,omitempty"`

	// 
	Status string `json:"status,omitempty"`

	// 
	SupportHotPlug bool `json:"support_hot_plug,omitempty"`

	// 
	Tags []VnfTag `json:"tags,omitempty"`

	// 
	UploadAt int32 `json:"upload_at,omitempty"`

	// 
	VendorName string `json:"vendor_name,omitempty"`

	// 
	ZeroDayConfigSpec string `json:"zero_day_config_spec,omitempty"`
}