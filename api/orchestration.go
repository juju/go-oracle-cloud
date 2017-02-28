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

	// Description is the description of this orchestration plan
	Description string `json:"description,omitempty"`

	// List of oplans. An object plan, or oplan, is a top-level orchestration attribute.
	Oplans []Oplans `json:"oplans"`

	// Name is the name of the orchestration
	Name string `json:"name"`
}

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

type Objects struct {
	Instances        []InstancesOrchestration `json:"instances"`
	Status_timestamp string                   `json:"status_timestamp,omitmepty"`
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

type NetworkingOrchestration struct {
	Interfaces map[string]interface{}
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
func (c Client) CreateOrchestration(p OrchestrationParams) (resp response.Orchestration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := c.endpoints["orchestration"] + "/"

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "POST",
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

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
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

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
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

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["orchestration"], c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DirectoryOrchestration retrieves the names of containers that contain objects
// that you can access
func (c Client) DirectoryOrchestration() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := c.endpoints["orchestration"] + "/"

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		url:       url,
		verb:      "GET",
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) AllOrchestrationNames() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["orchestration"], c.identify, c.username)

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		url:       url,
		verb:      "GET",
		resp:      &resp,
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

	url := fmt.Sprintf("%s%s", c.endpoints["orchestration"], currentName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "PUT",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
