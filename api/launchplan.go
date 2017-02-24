// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

	Attributes []map[string]interface{} `json:"attributes,omitempty"`

	// If set to true (default), then reverse DNS records are created.
	// If set to false, no reverse DNS records are created.
	Reverse_dns bool `json:"reverse_dns,omiempty"`
}

// InstanceParams used to feed the CreateInstance function
type InstanceParams struct {
	Relationships []string    `json:"relationships,omitempty"`
	Instances     []Instances `json:"instances"`
}

func (c Client) CreateInstance(params InstanceParams) (resp response.LaunchPlan, err error) {
	if params.Instances == nil || len(params.Instances) == 0 {
		return resp, errors.New("go-oracle-cloud: Empty slice of instance parameters")
	}

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	n := len(params.Instances)

	// here we are constructing the post body json
	for i := 0; i < n; i++ {
		// add the imagelist
		if params.Instances[i].Imagelist == "" {
			return resp, errors.New(
				"go-oracle-cloud: Empty image list in instance parameters",
			)
		}
		params.Instances[i].Imagelist = fmt.Sprintf(
			"/Compute-%s/%s/%s", c.identify, c.username, params.Instances[i].Imagelist,
		)

		// add the label
		if params.Instances[i].Label == "" {
			return resp, errors.New(
				"go-oracle-cloud: Empty label in instance parameters",
			)
		}

		// make the name oracle cloud complaint
		params.Instances[i].Name = fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, params.Instances[i].Name)

		// add the ssh keys
		keys := len(params.Instances[i].SSHKeys)
		for j := 0; j < keys; j++ {
			params.Instances[i].SSHKeys[j] = fmt.Sprintf("/Compute-%s/%s/%s",
				c.identify, c.username, params.Instances[i].SSHKeys[j],
			)
		}
	}

	url := fmt.Sprintf("%s/launchplan/", c.endpoint)
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &params,
		resp:   &resp,
		treat: func(resp *http.Response) (err error) {
			type Instances struct {
				Instances string `json:"instances,omitempty"`
			}
			type m struct {
				Message Instances `json:"message,omitempty"`
			}
			type refereceError struct {
				Message   string `json:"message,omitempty"`
				Reference string `json:"reference,omitmepty"`
			}

			var (
				errOut m
				refOut refereceError
			)

			if resp.StatusCode != http.StatusCreated {
				if resp.StatusCode >= http.StatusInternalServerError {
					json.NewDecoder(resp.Body).Decode(&refOut)
					return fmt.Errorf(
						"go-oracle-cloud: Error Api response %d %s Reference : %s",
						resp.StatusCode, refOut.Message, refOut.Reference,
					)
				} else {
					json.NewDecoder(resp.Body).Decode(&errOut)
					return fmt.Errorf(
						"go-oracle-cloud: Error Api response %d %s",
						resp.StatusCode, errOut.Message.Instances,
					)
				}
			}

			return nil
		},
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
