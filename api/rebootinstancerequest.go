// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

//You can reboot a running instance by creating a rebootinstancerequest object.

// CreateRebootInstanceRequest is used when we want to launch a restart on a instnace
// If your instance hangs after it starts running, you can use this request to reboot
// your instance. After creating this request, use GET /rebootinstancerequest/{name}
// to retrieve the status of the request. When the status of the rebootinstancerequest
// changes to complete, you know that the instance has been rebooted.
func (c Client) CreateRebootInstanceRequest(
	hard bool,
	instanceName string,
) (resp response.RebootInstanceRequest, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if instanceName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty instance name",
		)
	}

	url := fmt.Sprintf("%s/rebootinstancerequest/", c.endpoint)

	params := struct {
		Name string `json:"name"`
		Hard bool   `json:"hard"`
	}{
		Name: fmt.Sprintf("/Compute-%s/%s/%s/",
			c.identify, c.username, instanceName),
		Hard: hard,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		body:   &params,
		verb:   "POST",
		url:    url,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)
	strip(&resp.Instance)

	return resp, nil
}

// DeleteRebootInstanceRequest deletes a reboot instance request.
// No response is returned for the delete action.
func (c Client) DeleteRebootInstanceRequest(instanceName string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if instanceName == "" {
		return errors.New("go-oracle-cloud: Empty instance name")
	}

	url := fmt.Sprintf("%s/rebootinstancerequest/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, instanceName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "DELETE",
		url:    url,
		treat:  defaultDeleteTreat,
	}); err != nil {
		return err
	}

	return nil
}

// RebootInstanceRequestDetails retrieves details of the specified reboot instance request.
// You can use this request when you want to find out the status of a reboot instance request.
func (c Client) RebootInstanceRequestDetails(
	instanceName string,
) (resp response.RebootInstanceRequest, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if instanceName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty instance name",
		)
	}

	url := fmt.Sprintf("%s/rebootinstancerequest/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, instanceName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)
	strip(&resp.Instance)

	return resp, nil
}

//AllRebootInstanceRequest retrieves details of the reboot instance requests that are available in the specified container
func (c Client) AllRebootInstanceRequest() (resp response.AllRebootInstanceRequest, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/rebootinstancerequest/Compute-%s/%s",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	for key := range resp.Result {
		strip(&resp.Result[key].Name)
		strip(&resp.Result[key].Instance)
	}

	return resp, nil
}
