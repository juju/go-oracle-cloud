package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

type OrchestrationParams struct {
	// Relationships is the relationship between the objects
	// that are created by this orchestration.
	Relationships []string `json:"relationships,omitempty"`

	// Account shows the default account for your identity domain.
	Account string `json:"account"`

	// Description is the description of this orchestration plan
	Description string `json:"description,omitempty"`

	// Schedule for an orchestration consists
	// of the start and stop dates and times.
	Schedule response.Schedule `json:"schedule"`

	// Uri is the Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// List of oplans. An object plan, or oplan, is a top-level orchestration attribute.
	Oplans []response.Oplans `json:"oplans"`

	// Name is the name of the orchestration
	Name string `json:"name"`
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

// AddOrchestration Adds an orchestration to Oracle Compute Cloud Service.
func (c Client) AddOrchestration(p OrchestrationParams) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/orchestration/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteOrchestration deletes an orchestration from the system.
// The orchestration must be stopped to be deleted. All the objects created by
// the orchestration are deleted when you stop the orchestration.
// No response is returned for the delete action.
func (c Client) DeleteOrchestration(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure list")
	}

	url := fmt.Sprintf("%s/orchestration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat:  defaultDeleteTreat,
	}); err != nil {
		return err
	}

	return nil
}

// OrchestrationDetails retrieves details of the orchestrations
// that are available in the specified container
func (c Client) OrchestrationDetails(name string) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty secure list")
	}

	url := fmt.Sprintf("%s/orchestration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// AllOrchestration retrives all orchestration
func (c Client) AllOrchestration() (resp response.AllOrchestration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/orchestration/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateOrchestration updates an orchestration.
func (c Client) UpdateOrchestration(p OrchestrationParams, currentName string) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
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

	url := fmt.Sprintf("%s/orchestration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, p.Name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "PUT",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
