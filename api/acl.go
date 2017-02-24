// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// CreateAcl creates an access control list (ACL) to control
// the traffic between virtual NICs.
// An ACL consists of one or more security rules that is applied
// to a virtual NIC set. Each security rule may refer to a virtual
// NIC set in either the source or destination.See Workflow for
// After creating an ACL, you can associate it to one or more virtual NIC sets.
func (c Client) CreateAcl(
	name string,
	description string,
	enabledFlag bool,
	tags []string,
) (resp response.Acl, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Cannot create acl because name provided is empty",
		)
	}

	url := fmt.Sprintf("%s/network/v1/acl/", c.endpoint)

	acl := struct {
		Name        string   `json:"name"`
		Description string   `json:"description,omitempty"`
		EnabledFlag bool     `json:"enabledFlag"`
		Tags        []string `json:"tags,omitempty"`
	}{
		Name:        name,
		Description: description,
		EnabledFlag: enabledFlag,
		Tags:        tags,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &acl,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)

	return resp, nil
}

// DeleteAcl deletes specific acl that has the name
//
// If you no longer need to use an ACL, you can delete it.
// Remember, however, that security rules reference ACLs and
// ACLs are applied to vNICsets.
//
// If you delete an ACL that is referenced in one or more security rjkkkjules,
// those security rules can no longer be used.
//
// If you delete an ACL that is applied to a vNICset, the security rules in
// that ACL no longer apply to that vNICset. Before deleting an ACL,
// ensure that other ACLs are in place to provide access to relevant vNICsets.
//
// If you delete all the ACLs applied to a vNICset, some vNICs in that vNICset
// might become unreachable.
//
// If you want to disable an ACL and not delete it, use the UpdateAcl method
func (c Client) DeleteAcl(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	url := fmt.Sprintf("%s/network/v1/acl/Compute-%s/%s/%s",
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

// AllAcl retrieves details of all the ACLs
// that are available in the specified container.
func (c Client) AllAcl() (resp response.AllAcl, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/network/v1/acl/Compute-%s",
		c.endpoint, c.identify)

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

	for key, _ := range resp.Result {
		strip(&resp.Result[key].Name)
	}

	return resp, nil
}

// AclDetails retrieves information about the specified ACL.
func (c Client) AclDetails(name string) (resp response.Acl, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Cannot list acl details because name provided is empty")
	}

	url := fmt.Sprintf("%s/network/v1/acl/Compute-%s/%s/%s",
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

	strip(&resp.Name)

	return resp, nil
}

// UpdateAcl can update the description and tag fields for an ACL.
// You can also disable an ACL by setting the value of the enabledFlag to false.
// When you disable an ACL, it also disables the flow of traffic
// allowed by the security rules in scope of the ACL.
func (c Client) UpdateAcl(
	currentName string,
	newName string,
	description string,
	enableFlag bool,
	tags []string,
) (resp response.Acl, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: acl name provided is empty",
		)
	}

	if newName == "" {
		newName = currentName
	}

	url := fmt.Sprintf("%s/network/v1/acl/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	acl := response.Acl{
		Description: description,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
		EnableFlag: enableFlag,
		Tags:       tags,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		body:   &acl,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)

	return resp, nil
}
