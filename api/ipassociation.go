// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.
package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// IpAssociation retrieves the names of objects and subcontainers
// that you can access in the specified container.
func (c Client) AllIpAssociation() (resp response.AllIpAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/ip/association/Compute-%s/%s/",
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

// IpAssociationDetails retrieves details of the IP associations that are
// available in the specified container
func (c Client) IpAssociationDetails(name string) (resp response.IpAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty ip association name")
	}

	url := fmt.Sprintf("%s/ip/association/Compute-%s/%s/%s",
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

// Creates an association between an IP address
// and the vcable ID of an instance.
func (c Client) CreateIpAssociation(
	parentpool string,
	vcable string,
) (resp response.IpAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	params := struct {
		Parentpool string `json:"parentpool"`
		Vcable     string `json:"vcable"`
	}{
		Parentpool: parentpool,
		Vcable: fmt.Sprintf("Compute-%s/%s/%s",
			c.identify, c.username, vcable),
	}

	url := fmt.Sprintf("%s/ip/ipassociation/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		body:   &params,
		url:    url,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// Deletes the specified IP association
func (c Client) DeleteIpAssociation(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty ip association name provided")
	}

	url := fmt.Sprintf("%s/ip/ipassociation/%s/%s/%s",
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
