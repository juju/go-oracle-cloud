package response

// TODO(test this)

type LaunchPlan struct {
	Relationships []string   `json:"relationships,omitempty"`
	Instances     []Instance `json:"instances"`
}

type AllInstance struct {
	Result []Instance `json:"result"`
}

type AllInstanceNames struct {
	Result []string `json:"result"`
}

type List struct {
}

//todo(sgiulitti): still needs extensive testing but for now this should to the trick

type Instance struct {
	Domain                 string   `json:"domain"`
	Placement_requirements []string `json:"placement_requirements"`
	Ip                     string   `json:"ip"`
	Fingerprint            string   `json:"fingerprint,omitempty"`
	Site                   string   `json:"site,omitempty"`
	Shape                  string   `json:"shape"`
	//   	  "cluster": null,
	//       "last_state_change_time": null,
	//		 "error_exception": null,
	Start_requested bool `json:"start_requested"`
	// vethernets: null,
	//       "cluster_uri": null,
	// "target_node": null,
	Imagelist           ImageList  `json:"imagelist,omitempty"` // (todo)
	Image_format        string     `json:"image_format"`
	Relationships       []string   `json:"relationships,omitempty"`
	Availability_domain string     `json:"availability_domain"` // null
	Networking          Networking `json:"networking"`
	//  "seclist_associations": null,
	Hostname          string      `json:"hostname"`
	Quota_reservation interface{} `json:"quota_reservation,omitempty"` //null problem
	Disk_attach       string      `json:"disk_attach,omitempty"`
	//"suspend_file": null,
	// node: null,
	Resource_requirements ResourceRequirments `json:"resource_requirements"`
	Virtio                interface{}         `json:"virtio,omitempty"` // null problem
	Vnc                   string              `json:"vnc"`
	Desired_state         string              `json:"desired_state"`
	Storage_attachments   []Storage           `json:"storage_attachments,omitempty"`
	Start_time            string              `json:"start_time"`
	Label                 string              `json:"label,omitempty"`
	Id                    string              `json:"id,omitempty"`
	Priority              string              `json:"priority"`
	Platform              string              `json:"platform"`
	State                 string              `json:"state"`
	Tags                  []string            `json:"tags,omitempty"`
	//       "vnc_key": null,

	Quota        string      `json:"quota"`
	Entry        int         `json:"entry,omitempty"` // (todo)test
	Error_reason string      `json:"error_reason,omitempty"`
	SSHKeys      []string    `json:"sshkeys"`
	Resolvers    interface{} `json:"resolvers,omitempty"` // null problem
	Account      string      `json:"account"`
	Name         string      `json:"name"`
	Vcanble_id   string      `json:"vcable_id"`
	Hypervisor   Hypervisor  `json:"hypervisor"`
	Uri          string      `json:"uri"`
	Reverse_dns  bool        `json:"reverse_dns"`
	Attributes   Attributes  `json:"attributes"`
	Boot_order   []int       `json:"boot_order"`
	//"console": null,
	Launch_context string `json:"launch_context"`
	//       "delete_requested": null,
	//      "tracking_id": null,
	//"hypervisor_type": null,
	//"last_seen": null

}

type Networking struct {
	Eth0 Nic `json:"eth0"`
}

type Nic struct {
	Model     string   `json:"model,omitempty"`
	Seclists  []string `json:"seclists"`
	Dns       []string `json:"dns"`
	Nat       string   `json:"nat"` //null
	Vethernet string   `json:"vethernet"`
}

type Storage struct {
	Index               uint64 `json:"index"`
	Storage_volume_name string `json:"storage_volume_name"`
	Name                string `json:"name"`
}

type Hypervisor struct {
	Mode string `json:"mode"`
}

type Attributes struct {
	Userdata        map[string]interface{} `json:"userdata,omitempty"`
	SupportedShapes string                 `json:"supportedShapes"`
	DefaultShape    string                 `json:"defaultShape"`
	MinimumDiskSize string                 `json:"minimumDiskSize"`
	SSHKeys         []string               `json:"sshkeys"`
	Network         Network                `json:"network"`
}

type Dns struct {
	Domain      string `json:"domain"`
	Hostname    string `json:"hostname"`
	Vcable_eth0 string `json:"nimbula_vcable-eth0"`
}

type Network struct {
	Vcable_eth0    Vcable   `json:"nimbula_vcable-eth0"`
	Model          string   `json:"model,omitempty"`
	Vethernet_type string   `json:"vethernet_type"`
	Id             string   `json:"id"`
	Dhcp_options   []string `json"dhcp_options,omitempty"`
}

type Vcable struct {
	Vethernet_id   string   `json:"vethernet_id"`
	Vethernet      string   `json:"vethernet"`
	Address        []string `json:"address"`
	Model          string   `json:"model,omitempty"`
	Vethernet_type string   `json:"vethernet_type"`
	Id             string   `json:"id"`
	Dhcp_options   []string `json:"dhcp_options,omitempty"`
}

type ResourceRequirments struct {
	Compressed_size   uint64 `json:"compressed_size"`
	Is_root_ssd       bool   `json:"is_root_ssd"`
	Ram               uint64 `json:"ram"`
	Cpus              uint64 `json:"cpus"`
	Root_disk_size    uint64 `json:root_disk_size"`
	Io                uint64 `json:"io"`
	Decompressed_size uint64 `json:"decompressed_size"`
	Gpus              uint64 `json:"gpus"`
	Ssd_data_size     uint64 `json:"ssd_data_size"`
}
