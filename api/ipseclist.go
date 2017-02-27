// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// CreateSecIpList a security IP list. Note that, after creating a
// security IP list, you can add additional IP addresses to the list
// by using the CreateIpSecList again with just the additional IP addresses
// description is a description of the security IP list.
// name is the name of the SecIpList you wish to create
// secipentries a comma-separated list of the subnets
// (in CIDR format) or IPv4 addresses for which you want
// to create this security IP list.
// For example, to create a security IP list containing the
// IP addresses 203.0.113.1 and 203.0.113.2, enter one of the following:
// 203.0.113.0/30
// 203.0.113.1, 203.0.113.2
func (c Client) CreateSecIpList(
	description string,
	name string,
	secipentries []string,
) (resp response.SecIpList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty secure ip list name",
		)
	}

	if secipentries == nil || len(secipentries) == 0 {
		return resp, errors.New(
			"go-oracle-cloud: Slice secure ip entries nil or empty",
		)
	}

	params := struct {
		Description  string   `json:"description,omitempty"`
		Name         string   `json:"name"`
		Secipentries []string `json:"secipentries"`
	}{
		Description:  description,
		Name:         name,
		Secipentries: secipentries,
	}

	url := fmt.Sprintf("%s/seciplist/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "POST",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSecIpList deletes the specified security IP list. No response is returned.
// You can't delete system-provided security application that are
// available in the /oracle/public container.
func (c Client) DeleteSecIpList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure ip list name")
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

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

// SecIpListDetail retrieves information about the specified security IP list.
// You can use this request to verify whether CreateSecIpList
// or UpdateSecIpList operations were completed successfully.
func (c Client) IpSecListDetail(name string) (resp response.SecIpList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty sec ip list name",
		)
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// AllSecIpList retrieves details of the security IP lists that are in the account
func (c Client) AllSecIpList() (resp response.AllSecIpList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateSecIpList updates IP addresses and description of
// the specified security IP list. Note that this command replaces
// the values in the secipentries and description fields with
// the new values that you specify. To add one or more IP addresses
// to the existing list, run the add seciplist command and
// specify just the additional IP addresses.
func (c Client) UpdateSecIpList(
	description string,
	currentName string,
	newName string,
	secipentries []string,
) (resp response.SecIpList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty current secure ip list name",
		)
	}

	if newName == "" {
		newName = currentName
	}

	if secipentries == nil || len(secipentries) == 0 {
		return resp, errors.New(
			"go-oracle-cloud: Slice secure ip entries nil or empty",
		)
	}

	params := struct {
		Description  string   `json:"description,omitempty"`
		Name         string   `json:"name"`
		Secipentries []string `json:"secipentries"`
	}{
		Description:  description,
		Name:         newName,
		Secipentries: secipentries,
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "PUT",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
