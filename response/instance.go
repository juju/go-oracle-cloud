// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package response

import (
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/common"
)

type LaunchPlan struct {
	Relationships []string   `json:"relationships,omitempty"`
	Instances     []Instance `json:"instances"`
}

// AllInstances a slice of all instances in the
// oracle cloud account
type AllInstances struct {
	Result []Instance `json:"result"`
}

// AllInstanceNames a slice of all instance
// names in the oracle cloud account
type AllInstanceNames struct {
	Result []string `json:"result"`
}

// Instance represents an Oracle Compute Cloud Service
// instance is a virtual machine running a specific
// operating system and with CPU and memory resources that you specify.
type Instance struct {
	Domain                          string               `json:"domain"`
	Placement_requirements          []string             `json:"placement_requirements"`
	Ip                              string               `json:"ip"`
	Fingerprint                     string               `json:"fingerprint,omitempty"`
	Site                            string               `json:"site,omitempty"`
	Last_state_change_time          interface{}          `json:"last_state_change_time,omitempty"`
	Error_exception                 interface{}          `json:"error_exception,omitempty"`
	Cluster                         interface{}          `json:"cluster,omitempty"`
	Shape                           string               `json:"shape"`
	Start_requested                 bool                 `json:"start_requested"`
	Vethernets                      interface{}          `json:"vethernets,omitempty"`
	Imagelist                       string               `json:"imagelist,omitempty"`
	Image_format                    string               `json:"image_format"`
	Cluster_uri                     interface{}          `json:"cluster_uri,omitempty"`
	Relationships                   []string             `json:"relationships,omitempty"`
	Target_node                     interface{}          `json:"target_node,omitempty"`
	Availability_domain             interface{}          `json:"availability_domain,omitempty"`
	Networking                      Networking           `json:"networking"`
	Seclist_associations            interface{}          `json:"seclist_associations,omitempty"`
	Hostname                        string               `json:"hostname"`
	State                           common.InstanceState `json:"state"`
	Disk_attach                     string               `json:"disk_attach,omitempty"`
	Label                           string               `json:"label,omitempty"`
	Priority                        string               `json:"priority"`
	Platform                        string               `json:"platform"`
	Quota_reservation               interface{}          `json:"quota_reservation,omitempty"`
	Suspend_file                    interface{}          `json:"suspend_file,omitempty"`
	Node                            interface{}          `json:"node,omitempty"`
	Resource_requirements           ResourceRequirments  `json:"resource_requirements"`
	Virtio                          interface{}          `json:"virtio,omitempty"`
	Vnc                             string               `json:"vnc,omitempty"`
	Desired_state                   common.InstanceState `json:"desired_state"`
	Storage_attachments             []Storage            `json:"storage_attachments,omitempty"`
	Start_time                      string               `json:"start_time"`
	Storage_attachment_associations []interface{}        `json:"storage_attachment_associations,omitempty"`
	Quota                           string               `json:"quota"`
	Vnc_key                         interface{}          `json:"vnc_key,omitempty"`
	Numerical_priority              uint64               `json:"numerical_priority"`
	Suspend_requested               bool                 `json:"suspend_requested"`
	Entry                           int                  `json:"entry"`
	Error_reason                    string               `json:"error_reason,omitempty"`
	Nat_associations                interface{}          `json:"nat_associations,omitempty"`
	SSHKeys                         []string             `json:"sshkeys,omitemtpy"`
	Tags                            []string             `json:"tags,omitempty"`
	Resolvers                       interface{}          `json:"resolvers,omitempty"`
	Metrics                         interface{}          `json:"metrics,omitempty"`
	Account                         string               `json:"account"`
	Node_uuid                       interface{}          `json:"node_uuid,omitempty"`
	Name                            string               `json:"name"`
	Vcable_id                       common.VcableID      `json:"vcable_id,omitempty"`
	Higgs                           interface{}          `json:"higgs,omitempty"`
	Hypervisor                      Hypervisor           `json:"hypervisor"`
	Uri                             string               `json:"uri"`
	Console                         interface{}          `json:"console,omitempty"`
	Reverse_dns                     bool                 `json:"reverse_dns"`
	Launch_context                  string               `json:"launch_context"`
	Delete_requested                interface{}          `json:"delete_requested,omitempty"`
	Tracking_id                     interface{}          `json:"tracking_id,omitempty"`
	Hypervisor_type                 interface{}          `json:"hypervisor_type,omitempty"`
	Attributes                      Attributes           `json:"attributes"`
	Boot_order                      []int                `json:"boot_order,omitempty"`
	Last_seen                       interface{}          `json:"last_seen,omitempty"`
}

// Attributes holds a map of attributes that is returned from
// the instance response
// This attributes can have user data scripts or any other
// key, value passed to be executed when the instance will strat, inits
type Attributes map[string]interface{}

// Networking is a json object of string keys
// Every key is the name of the interface example eth0,eth1, etc.
// And every valu is a predefined json objects that holds infromation
// about the interface
type Networking map[string]map[string]interface{}

// Nic type used to hold information from a
// given interface card
// This wil be used to dump all information from the
// Netowrking type above
type Nic struct {
	Vethernet string
	Nat       string
	Model     string
	Seclists  []string
	Dns       []string
}

// Interfaces returns a map of Nics from the response instance
// If we have some unexpected values in the response other than what we
// expect(declared in the Nic struct above), we return a descriptfull error
// If the response is empty the func will return nil, nil
func (n Networking) Interfaces() (interfaces map[string]Nic, err error) {
	if n == nil || len(n) == 0 {
		return nil, nil
	}

	interfaces = make(map[string]Nic)

	// for every interface nic eth0, eth1..
	for nic, v := range n {
		// make a new default nic
		ic := Nic{}

		// for every key in the nic
		for key, k := range v {

			// decide on what kind of type the key has
			switch k.(type) {

			// if we are dealing with the string type
			// that means we should assume we have
			// this keys "vethernet", "model" and "nat"
			case string:
				switch key {
				case "vethernet":
					ic.Vethernet = k.(string)
				case "model":
					ic.Model = k.(string)
				case "nat":
					ic.Nat = k.(string)

				// if there is other key bail out
				default:
					return nil, fmt.Errorf(
						"Unknown key %s in nic", key,
					)
				}

			// if we are dealing with the []interface type
			// that means we should assume we have
			// this keys seclists and dns
			case []interface{}:
				switch key {
				case "seclists":
					for _, m := range k.([]interface{}) {
						// append into the slice the value
						elem, ok := m.(string)
						// if it's not a slice of strings we should bail out
						if !ok {
							return nil, fmt.Errorf(
								"Different type in seclists, type %T", m,
							)
						}

						ic.Seclists = append(ic.Seclists, elem)
					}
				case "dns":
					for _, m := range k.([]interface{}) {
						elem, ok := m.(string)
						if !ok {
							return nil, fmt.Errorf(
								"Different type in dns, type %T", m,
							)
						}

						ic.Dns = append(ic.Dns, elem)
					}

				// if there is other key bail out
				default:
					return nil, fmt.Errorf(
						"Unknown key %s in nic ", key,
					)
				}

			// if have other type, bail out
			default:
				return nil, fmt.Errorf(
					"Unknown key types, recived type %T", k,
				)
			}
		}
		// assign the new nic
		interfaces[nic] = ic
	}

	return interfaces, nil
}

type Storage struct {
	Index               uint64 `json:"index"`
	Storage_volume_name string `json:"storage_volume_name"`
	Name                string `json:"name"`
}

type Hypervisor struct {
	Mode string `json:"mode"`
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
	Dhcp_options   []string `json:"dhcp_options,omitempty"`
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
	Root_disk_size    uint64  `json:"root_disk_size"`
	Io                uint64  `json:"io"`
	Decompressed_size uint64  `json:"decompressed_size"`
	Gpus              uint64  `json:"gpus"`
	Ssd_data_size     uint64  `json:"ssd_data_size"`
}
