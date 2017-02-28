// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// VirtualNic retrives a virtual nic with that has a given name
func (c Client) VirtualNic(name string) (resp response.VirtualNic, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty virtual nic name",
		)
	}

	url := fmt.Sprintf("%s%s", c.endpoints["vnc"], name)

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

	return resp, nil
}

// AllVirtualNic returns all virtual nic that are in the oracle account
func (c Client) AllVirtualNic() (resp response.AllVirtualNic, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s%s", c.endpoints["vnc"], name)

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["vnc"], c.identify, c.username)

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

	return resp, nil
}
