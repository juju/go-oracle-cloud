// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

type Instances struct {
	// Shape represents every instance in oracle cloud has a predefined shape
	// in order to create a virtual instance
	// here it can be specify the computing power shape
	Shape string `json:"shape"`

	// Imagelist is the virtual image that will be used
	// in order to init the instance
	Imagelist string `json:"imagelist,omitempty"`

	// Name is the name of the instance
	Name string `json:"name,omitempty"`

	// Label is used when defining relationships in an orchestration.
	Label string `json:"label,omitempty"`

	// SSHKeys that will be installed on the Instance
	SSHKeys []string `json:"sshkeys,omitempty"`

	// hostname is assigned to the instance
	// on an Oracle Linux instance, this host name
	// is displayed in response to the hostname command.
	Hostname string `json:"hostname,omitempty"`

	// Tags by assigning a human-friendly tag to an instance
	// you can identify the instance easily when you perform
	// an instance listing.
	// These tags aren’t available from within the instance.
	Tags []string `json:"tags,omitempty"`

	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// If set to true (default), then reverse DNS records are created.
	// If set to false, no reverse DNS records are created.
	Reverse_dns bool `json:"reverse_dns,omiempty"`
}

// InstanceParams used to feed the CreateInstance function
type InstanceParams struct {
	Relationships []string    `json:"relationships,omitempty"`
	Instances     []Instances `json:"instances"`
}

func (i InstanceParams) validate() (err error) {
	for _, val := range i.Instances {
		if val.Imagelist == "" {
			return errors.New(
				"go-oracle-cloud: Empty image list in instance parameters",
			)
		}
		if val.Label == "" {
			return errors.New(
				"go-oracle-cloud: Empty label in instance parameters",
			)
		}
	}

	return nil
}

func (c Client) CreateInstance(params InstanceParams) (resp response.LaunchPlan, err error) {
	if params.Instances == nil || len(params.Instances) == 0 {
		return resp, errors.New("go-oracle-cloud: Empty slice of instance parameters")
	}

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if err := params.validate(); err != nil {
		return resp, err
	}

	url := c.endpoints["launchplan"]

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &params,
		resp:   &resp,
		treat:  debugTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
