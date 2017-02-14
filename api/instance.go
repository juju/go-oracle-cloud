package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// DeleteInstance shuts down an instance and removes it permanently
// from the system.
func (c Client) DeleteInstance(uuid string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if uuid == "" {
		return errors.New("go-oracle-cloud: uuid provided is empty")
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, uuid)

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

// AllInstances retrieves details of the instances that are in the specified
// container and match the specified query criteria.
// If you don't specify any query criteria, then details
// of all the instances in the container are displayed.
func (c Client) AllInstances() (resp response.AllInstance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/",
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

// InstanceDetails retrieves details of the specified instance.
func (c Client) InstanceDetails(uuid string) (resp response.Instance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, uuid)

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

// AllInstanceNames retrieves the names of objects and subcontainers
// that you can access in the specified container.
func (c Client) AllInstanceNames() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		url:       url,
		verb:      "GET",
		treat:     defaultTreat,
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
