package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

type Shape uint8

const (
	// General purpose
	OC3 Shape = iota // 1CPU + 7.5 GB Memory
	OC4              // 2CPU + 15GB Memory
	OC5              // 4CPU + 30GB Memory
	OC6              // 8CPU + 60GB Memory
	OC7              // 16CPU + 120GB Memory
	// High memory
	OC1M // 1CPU + 15GB memory
	OC2M // 2CPU + 30GB Memory
	OC3M // 4CPU + 60GB Memory
	OC4M // 8CPU + 120GB Memory
	OC5M // 16CPU + 240GB Memory
)

// shapes is a mapping between shapes
var shapes = map[Shape]string{
	OC3:  "oc3",
	OC4:  "oc4",
	OC5:  "oc5",
	OC6:  "oc6",
	OC7:  "oc7",
	OC1M: "oc1m",
	OC2M: "oc2m",
	OC3M: "oc3m",
	OC4M: "oc4m",
	OC5M: "oc5m",
}

// InstanceParams used to feed the CreateInstance function
type InstanceParams struct {
	// Shape represents every instance in oracle cloud has a predefined shape
	// in order to create a virtual instance
	// we need to specify the computing power shape
	Shape Shape
	// Imagelist is the virtual image that will be used
	// in order to init the instance
	Imagelist string
	// Name is the name of the instance
	Name string
	// Label
	Label string
	// SSHKeys that will be installed on the Instance
	SSHKeys []string
}

func (c Client) CreateInstance(params []InstanceParams) (resp response.LaunchPlan, err error) {
	if params != nil || len(params) == 0 {
		return resp, errors.New("go-oracle-cloud: Empty slice of instance parameters")
	}

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	// no need to make this type public inside this package
	type args struct {
		Shape     string   `json:"shape"`
		Imagelist string   `json:"imagelist"`
		Name      string   `json:"name"`
		Label     string   `json:"label"`
		SSHKeys   []string `json:"sshkeys"`
	}

	instanceArgs := make([]args, len(params), len(params))
	n := len(params)
	for i := 0; i < n; i++ {
		if params[i].Imagelist == "" {
			return resp, errors.New("go-oracle-cloud: Empty image list in instance parameters")
		}
		instanceArgs[i].Imagelist = params[i].Imagelist
		if params[i].Label == "" {
			return resp, errors.New("go-oracle-cloud: Empty label in instance parameters")
		}
		instanceArgs[i].Label = params[i].Label
		instanceArgs[i].Name = fmt.Sprintf("Compute-%s/%s/%s",
			c.identify, c.username, params[i].Name)
		shape, ok := shapes[params[i].Shape]
		if !ok {
			return resp, ErrUndefinedShape
		}
		instanceArgs[i].Shape = shape
		keys := len(instanceArgs[i].SSHKeys)
		for j := 0; j < keys; j++ {
			instanceArgs[i].SSHKeys[i] = fmt.Sprintf("Compute-%s/%s/%s",
				c.identify, c.username, params[i].SSHKeys[i])
		}
	}

	url := fmt.Sprintf("%s/%s/", c.endpoint, "launchplan")
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &instanceArgs,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) DeleteInstance(uuid string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if uuid == "" {
		return errors.New("go-oracle-cloud: uuid provided is empty")
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "instance", c.identify, c.username, uuid)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusNoContent {
				return fmt.Errorf("go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
			return nil
		},
	}); err != nil {
		return err
	}

	return nil
}

func (c Client) AllInstance() (resp response.AllInstance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "instance", c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) InstanceDetails(uuid string) (resp response.Instance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "instance", c.identify, c.username, uuid)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) AllInstanceNames() (resp response.AllInstanceNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "instance", c.identify, c.username)

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		url:       url,
		verb:      "GET",
		treat:     defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) InstanceNames() (resp response.AllInstanceNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}
	url := fmt.Sprintf("%s/%s/", c.endpoint, "instance")

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		url:       url,
		verb:      "GET",
		treat:     defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
