package response

// Orchestration is an orchestration defines the attributes and interdependencies of a collection of compute,
// networking, and storage resources in Oracle Compute Cloud Service. You can use orchestrations to automate
// the provisioning and lifecycle operations of an entire virtual compute topology.
// After creating an orchestration (in a JSON-formatted file) and adding it to Oracle Compute Cloud Service,
// you can trigger the creation and removal all the resources defined in the orchestration with a single step.
// An orchestration contains one or more object plans (oplans). The attributes that you can specify in
// an oplan vary depending on the object type (obj_type). For detailed information about the object
// types that you can create by using orchestrations and the attributes for each object type,
// see Attributes in Orchestrations in Using Oracle Compute Cloud Service (IaaS).
// You can add, start, stop, get, update, and delete orchestrations.
type Orchestration struct {
	// Account shows the default account for your identity domain.
	Account string `json:"account"`

	// Description is the description of this orchestration plan
	Description string `json:"description,omitempty"`

	// Info the nested parameter errors shows which object
	//  in the orchestration has encountered an error.
	// Empty if there are no errors.
	Info interface{} `json:"info,omitempty"`

	// Name is the name of the orchestration
	Name string `json:"name"`

	// List of oplans. An object plan, or oplan, is a top-level orchestration attribute.
	Oplans []Oplans `json:"oplans"`

	// Relationships is the relationship between the objects
	// that are created by this orchestration.
	Relationships []string `json:"relationships,omitempty"`

	// Schedule for an orchestration consists
	// of the start and stop dates and times.
	Schedule Schedule `json:"schedule"`

	// Status shows the current status of the orchestration.
	Status string `json:"status"`

	// Status_timestamp this information is generally displayed
	// at the end of the orchestration JSON.
	// It indicates the time that the current view of the
	// orchestration was generated. This information shows only when
	// the orchestration is running.
	Status_timestamp string `json:"status_timestamp"`

	// Uri is the Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// User is the user of the orchestration
	User string `json:"user"`
}

// AllOrchestrations a holds a slice of all
// orchestrations of a oracle cloud account
type AllOrchestrations struct {
	Result []Orchestration `json:"result,omitmepty"`
}

type Schedule struct {
	Start_time *string `json:"start_time,omitempty"`
	Stop_time  *string `json:"stop_time,omitempty"`
}

type Oplans struct {
	// Ha_policy indicates that description is not available
	Ha_policy string `json:"ha_policy,omitempty"`

	// Info dictionary for the oplan.
	Info Info `json:"info,omitempty"`

	// Label is the description of this object plan.
	Label string `json:"label"`

	// Obj_type type of the object.
	Obj_type string `json:"obj_type"`

	// Objects list of object dictionaries
	// or object names.
	Objects []Objects `json:"objects"`

	// Status is the most recent status.
	Status string `json:"status"`

	// Status_timestamp Timestamp of the most-recent status change.
	Status_timestamp string `json:"status_timestamp"`
}

type Info struct {
	Errors map[string]string `json:"errors,omitempty"`
}

type Objects struct {
	Instances        []InstancesOrchestration `json:"instances"`
	Status_timestamp string                   `json:"status_timestamp,omitmepty"`
}

// todo(sgiulitti) test it make sure.
type NetworkingOrchestration struct {
	Interfaces map[string]interface{}
}

type InstancesOrchestration struct {
	Shape               string                  `json:"shape"`
	Label               string                  `json:"label"`
	Imagelist           string                  `json:"imagelist"`
	Name                string                  `json:"name"`
	Boot_order          []string                `json:"boot_order,omitempty"`
	Attributes          AttributesOrchestration `json:"attributes,omitmepty"`
	Storage_attachments []StorageOrhcestration  `json:"storage_attachments,omitmepty"`
	Uri                 *string                 `json:"uri,omitempty"`
	SSHkeys             []string                `json:"sshkeys,omitmepty"`
	Tags                []string                `json:"tags,omitmepty"`
	Networking          NetworkingOrchestration `json:"networking,omitempty"`
}

type StorageOrhcestration struct {
	Info map[string]string
}

type AttributesOrchestration struct {
	Userdata              map[string]string `json:"userdata,omitempty"`
	Nimbula_orchestration string            `json:"nimbula_orchestration,omitempty"`
}
