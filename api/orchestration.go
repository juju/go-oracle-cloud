// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/juju/go-oracle-cloud/common"
	"github.com/juju/go-oracle-cloud/response"
)

// OrchestrationParams orchestraiton params used as params in
// CreateOrchestration and UpdateOrchestration
type OrchestrationParams struct {

	// Relationships is the relationship between the objects
	// that are created by this orchestration.
	Relationships []Relationship `json:"relationships,omitempty"`

	// Description is the description of this orchestration plan
	Description string `json:"description,omitempty"`

	// List of oplans. An object plan, or oplan, is a top-level orchestration attribute.
	Oplans []Oplans `json:"oplans"`

	// Name is the name of the orchestration
	Name string `json:"name"`

	// Schedule for an orchestration consists of
	// the start and stop dates and times
	Schedule Schedule `json:"schedule"`
}

// Relationship type that will describe the relationship
// between objects
type Relationship struct {

	// ToOplan to witch orchestration plan should
	// be the orchestration in a relationship
	ToOplan string `json:"to_oplan,omitempty"`

	// Oplan orchestration plan
	Oplan string `json:"oplan,omitempty"`

	// The type of relationship that this orchestration
	// has with the other one in the ToOplan field
	Type string `json:"type,omitempty"`
}

// Schedule for an orchestration consists of
// the start and stop dates and times
type Schedule struct {

	// Start_time when the orchestration will start
	// Date and time, in ISO 8601 format, when you want to start the orchestration
	// If you do not specify a value, the orchestration starts immediately
	Start_time *string `json:"start_time,omitempty"`

	// Stop_time when the orchestration will stop
	// Date and time, in ISO 8601 format, when you want
	// to stop the orchestration
	Stop_time *string `json:"stop_time,omitempty"`
}

// Oplans orchestration plans holds important details
// about the orchestration
type Oplans struct {

	// Ha_policy indicates that description is not available
	Ha_policy string `json:"ha_policy,omitempty"`

	// Label is the description of this object plan.
	Label string `json:"label"`

	// Obj_type type of the object.
	Obj_type string `json:"obj_type"`

	// Objects list of object dictionaries
	// or object names.
	Objects []Objects `json:"objects"`
}

// Objects types used for storing object dictionaries
// or object names for the orchestration
type Objects struct {
	// Instances list of instances
	Instances []InstancesOrchestration `json:"instances"`

	Status_timestamp string `json:"status_timestamp,omitmepty"`
}

// InstancesOrchestration holds information for
// an instances inside the orchestration object
type InstancesOrchestration struct {

	// Shape is the shape of the instnace
	Shape string `json:"shape"`

	// Label is the label of the instance
	Label string `json:"label"`

	// Imagelist is the image from what was created
	Imagelist string `json:"imagelist"`

	// Name of the instance
	Name string `json:"name"`

	// Boot_order is the number in what order the instance is booting
	Boot_order []string `json:"boot_order,omitempty"`

	// Attributes list of orchestration attributes
	Attributes AttributesOrchestration `json:"attributes,omitmepty"`

	// Storage_attachments list of storages that the instnaces has
	Storage_attachments []StorageOrhcestration `json:"storage_attachments,omitmepty"`

	// Uri of the instnace
	Uri *string `json:"uri,omitempty"`

	// SSHKeys of the instance
	SSHkeys []string `json:"sshkeys,omitmepty"`

	// Tags are a list of tags, aliases for the instance
	Tags []string `json:"tags,omitmepty"`

	// Networking information of the instance
	Networking common.Networking `json:"networking,omitempty"`
}

//TODO
type StorageOrhcestration struct {
	Info map[string]string
}

type AttributesOrchestration struct {
	Userdata              map[string]string `json:"userdata,omitempty"`
	Nimbula_orchestration string            `json:"nimbula_orchestration,omitempty"`
}

func (o OrchestrationParams) validate() (err error) {
	if o.Oplans == nil || len(o.Oplans) == 0 {
		return errors.New(
			"go-oracle-cloud: Empty orchestration plans",
		)
	}

	if o.Name == "" {
		return errors.New(
			"go-oracle-clod: Empty name in orchestration",
		)
	}

	for _, val := range o.Oplans {
		if val.Ha_policy == "" {
			return errors.New(
				"go-oracle-cloud: Empty HA_policy in orchestration plan",
			)
		}

		if val.Label == "" {
			return errors.New(
				"go-oracle-cloud: Empty label in orchestration plan",
			)
		}

		if val.Objects == nil || len(val.Objects) == 0 {
			return errors.New(
				"go-oracle-cloud: Empty Objects in orchestration plan",
			)
		}
		if val.Obj_type == "" {
			return errors.New(
				"go-oracle-cloud: Empty object type in orchestration plan",
			)
		}
	}

	return nil
}

// CreateOrchestration Adds an orchestration to Oracle Compute Cloud Service.
func (c *Client) CreateOrchestration(p OrchestrationParams) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := c.endpoints["orchestration"] + "/"

	if err = c.request(paramsRequest{
		url:  url,
		body: &p,
		verb: "POST",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteOrchestration deletes an orchestration from the system.
// The orchestration must be stopped to be deleted. All the objects created by
// the orchestration are deleted when you stop the orchestration.
// No response is returned for the delete action.
func (c *Client) DeleteOrchestration(name string) (err error) {
	if !c.isAuth() {
		return errNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure list")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "DELETE",
	}); err != nil {
		return err
	}

	return nil
}

// OrchestrationDetails retrieves details of the orchestrations
// that are available in the specified container
func (c *Client) OrchestrationDetails(name string) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty secure list")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "GET",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// AllOrchestrations retrives all orchestration
// You can filter by status
func (c *Client) AllOrchestrations(filter []Filter) (resp response.AllOrchestrations, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["orchestration"], c.identify, c.username)

	if err = c.request(paramsRequest{
		url:    url,
		verb:   "GET",
		resp:   &resp,
		filter: filter,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DirectoryOrchestration retrieves the names of containers that contain objects
// that you can access
func (c *Client) DirectoryOrchestration() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := c.endpoints["orchestration"] + "/"

	if err = c.request(paramsRequest{
		directory: true,
		url:       url,
		verb:      "GET",
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) AllOrchestrationNames() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["orchestration"], c.identify, c.username)

	if err = c.request(paramsRequest{
		directory: true,
		url:       url,
		verb:      "GET",
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil

}

// UpdateOrchestration updates an orchestration.
func (c *Client) UpdateOrchestration(p OrchestrationParams, currentName string) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty orchestration name",
		)
	}

	if p.Name == "" {
		p.Name = currentName
	}

	if err := p.validate(); err != nil {
		return resp, err
	}

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], currentName)

	if err = c.request(paramsRequest{
		url:  url,
		body: &p,
		verb: "PUT",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
