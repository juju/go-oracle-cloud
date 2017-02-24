// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package response

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

type Instance struct {
	Domain                          string              `json:"domain"`
	Placement_requirements          []string            `json:"placement_requirements"`
	Ip                              string              `json:"ip"`
	Fingerprint                     string              `json:"fingerprint,omitempty"`
	Site                            string              `json:"site,omitempty"`
	Last_state_change_time          interface{}         `json:"last_state_change_time,omitempty"`
	Error_exception                 interface{}         `json:"error_exception,omitempty"`
	Cluster                         interface{}         `json:"cluster,omitempty"`
	Shape                           string              `json:"shape"`
	Start_requested                 bool                `json:"start_requested"`
	Vethernets                      interface{}         `json:"vethernets,omitempty"`
	Imagelist                       string              `json:"imagelist,omitempty"`
	Image_format                    string              `json:"image_format"`
	Id                              string              `json:"id,omitempty"`
	Cluster_uri                     interface{}         `json:"cluster_uri,omitempty"`
	Relationships                   []string            `json:"relationships,omitempty"`
	Target_node                     interface{}         `json:"target_node,omitempty"`
	Availability_domain             interface{}         `json:"availability_domain,omitempty"`
	Networking                      Networking          `json:"networking"`
	Seclist_associations            interface{}         `json:"seclist_associations,omitempty"`
	Hostname                        string              `json:"hostname"`
	State                           string              `json:"state"`
	Disk_attach                     string              `json:"disk_attach,omitempty"`
	Label                           string              `json:"label,omitempty"`
	Priority                        string              `json:"priority"`
	Platform                        string              `json:"platform"`
	Quota_reservation               interface{}         `json:"quota_reservation,omitempty"`
	Suspend_file                    interface{}         `json:"suspend_file,omitempty"`
	Node                            interface{}         `json:"node,omitempty"`
	Resource_requirements           ResourceRequirments `json:"resource_requirements"`
	Virtio                          interface{}         `json:"virtio,omitempty"`
	Vnc                             string              `json:"vnc,omitempty"`
	Desired_state                   string              `json:"desired_state"`
	Storage_attachments             []Storage           `json:"storage_attachments,omitempty"`
	Start_time                      string              `json:"start_time"`
	Storage_attachment_associations []interface{}       `json:"storage_attachment_associations,omitempty"`
	Quota                           string              `json:"quota"`
	Vnc_key                         interface{}         `json:"vnc_key,omitempty"`
	Numerical_priority              uint64              `json:"numerical_priority"`
	Suspend_requested               bool                `json:"suspend_requested"`
	Entry                           int                 `json:"entry"`
	Error_reason                    string              `json:"error_reason,omitempty"`
	Nat_associations                interface{}         `json:"nat_associations,omitempty"`
	SSHKeys                         []string            `json:"sshkeys,omitemtpy"`
	Tags                            []string            `json:"tags,omitempty"`
	Resolvers                       interface{}         `json:"resolvers,omitempty"`
	Metrics                         interface{}         `json:"metrics,omitempty"`
	Account                         string              `json:"account"`
	Node_uuid                       interface{}         `json:"node_uuid,omitempty"`
	Name                            string              `json:"name"`
	Vcanble_id                      interface{}         `json:"vcable_id,omitempty"`
	Higgs                           interface{}         `json:"higgs,omitempty"`
	Hypervisor                      Hypervisor          `json:"hypervisor"`
	Uri                             string              `json:"uri"`
	Console                         interface{}         `json:"console,omitempty"`
	Reverse_dns                     bool                `json:"reverse_dns"`
	Launch_context                  string              `json:"launch_context"`
	Delete_requested                interface{}         `json:"delete_requested,omitempty"`
	Tracking_id                     interface{}         `json:"tracking_id,omitempty"`
	Hypervisor_type                 interface{}         `json:"hypervisor_type,omitempty"`
	Attributes                      Attributes          `json:"attributes"`
	Boot_order                      []int               `json:"boot_order,omitempty"`
	Last_seen                       interface{}         `json:"last_seen,omitempty"`
}

type Networking struct {
	Eth0 Nic `json:"eth0"`
}

type Nic struct {
	Model     string      `json:"model,omitempty"`
	Seclists  []string    `json:"seclists"`
	Dns       []string    `json:"dns"`
	Nat       interface{} `json:"nat,omitempty"`
	Vethernet string      `json:"vethernet"`
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
	Compressed_size   uint64  `json:"compressed_size"`
	Is_root_ssd       bool    `json:"is_root_ssd"`
	Ram               uint64  `json:"ram"`
	Cpus              float64 `json:"cpus"`
	Root_disk_size    uint64  `json:root_disk_size"`
	Io                uint64  `json:"io"`
	Decompressed_size uint64  `json:"decompressed_size"`
	Gpus              uint64  `json:"gpus"`
	Ssd_data_size     uint64  `json:"ssd_data_size"`
}
