// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/juju/go-oracle-cloud/response"
)

type EndpointParams struct{}

func (e EndpointParams) validate() (err error) {
	return nil
}

// CreateVpnEndpoint creates a VPN tunnel between
// your data center and your Oracle Compute
// Cloud Service site. You can create up to 20 VPN
// tunnels to your Oracle Compute Cloud Service site
func (c *Client) CreateVpnEndpoint(
	p EndpointParams,
) (resp response.VpnEndpoint, err error) {

	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := c.endpoints["vpnendpoint"] + "/"

	if err = c.request(paramsRequest{
		verb: "POST",
		resp: &resp,
		body: &p,
		url:  url,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) DeleteVpnEndpoint(name string) (err error) {
	if !c.isAuth() {
		return errNotAuth
	}

	if name == "" {
		return errors.New(
			"go-oracle-cloud: Empty vpn endpoint name",
		)
	}

	url := fmt.Sprintf("%s%s", c.endpoints["vpnendpoint"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "DELETE",
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) VpnEndpointsDetails(
	name string,
) (resp response.VpnEndpoint, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty vpn endpoints name",
		)
	}

	url := fmt.Sprintf("%s%s", c.endpoints["vpnendpoint"], name)

	if err = c.request(paramsRequest{
		url:  url,
		resp: &resp,
		verb: "GET",
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) AllVpnEndpoints(
	filter []Filter,
) (resp response.AllVpnEndpoints, err error) {

	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["vpnendpoint"], c.identify, c.username)

	if err = c.request(paramsRequest{
		url:    url,
		filter: filter,
		verb:   "GET",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
