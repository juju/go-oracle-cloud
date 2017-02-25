// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// DeleteInstance shuts down an instance and removes it permanently
// from the system.
// Example of name f653a677-b566-4f92-8e93-71d47b364119
func (c Client) DeleteInstance(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty instance name")
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/%s",
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

	for key := range resp.Result {
		strip(&resp.Result[key].Imagelist)
		for alt := range resp.Result[key].SSHKeys {
			strip(&resp.Result[key].SSHKeys[alt])
		}
		list := strings.Split(resp.Result[key].Name, "/")
		resp.Result[key].Name = list[len(list)-2] + "/" + list[len(list)-1]
	}
	return resp, nil
}

// InstanceDetails retrieves details of the specified instance.
// Name is the form of dev-name/uuid
func (c Client) InstanceDetails(name string) (resp response.Instance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/instance/Compute-%s/%s/%s",
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

	strip(&resp.Imagelist)
	for alt := range resp.SSHKeys {
		strip(&resp.SSHKeys[alt])
	}
	list := strings.Split(resp.Name, "/")
	resp.Name = list[len(list)-2] + "/" + list[len(list)-1]

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
