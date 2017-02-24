// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// AddSSHKey adds into the oracle cloud account an ssh key
func (c Client) AddSHHKey(
	name string,
	key string,
	enabled bool,
) (resp response.SSH, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: empty key name")
	}

	if key == "" {
		return resp, errors.New("go-oracle-cloud: ssh key provided is empty")
	}

	ssh := struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
		Name    string `json:"name"`
	}{
		Enabled: enabled,
		Key:     key,
		Name:    fmt.Sprintf("/Compute-%s/%s/%s", c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/%s/", c.endpoint, "sshkey")
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &ssh,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSSHKey deteles a ssh key with a specific name
func (c Client) DeleteSSHKey(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: empty key name")
	}

	keyname := fmt.Sprintf("Compute-%s/%s/%s",
		c.identify, c.username, name)
	url := fmt.Sprintf("%s/%s/%s",
		c.endpoint, "sshkey", keyname)

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

// SSHKeyDetails returns all details of a specific key
func (c Client) SSHKeyDetails(name string) (resp response.SSH, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: empty key name")
	}

	keyname := fmt.Sprintf("Compute-%s/%s/%s", c.identify, c.username, name)
	url := fmt.Sprintf("%s/%s/%s", c.endpoint, "sshkey", keyname)
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

// AllSShKeysDetails returns list of all keys with all the details
func (c Client) AllSSHKeyDetails() (resp response.AllSSH, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/", c.endpoint, "sshkey", c.identify, c.username)
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

// AllSSHKeyNames returns a list of all ssh keys by names of the user
func (c Client) AllSSHKeyNames() (resp response.AllSSHNames, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "sshkey", c.identify, c.username)
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

// UpdateSSHKey change the content and details of a specific ssh key
// If the key is invalid it will retrun 400 status code. Make sure the key is a valid ssh public key
func (c Client) UpdateSSHKey(
	name string,
	key string,
	enabled bool,
) (resp response.SSH, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: empty key name")
	}

	if key == "" {
		return resp, errors.New("go-oracle-cloud: ssh key provided is empty")
	}

	ssh := struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
		Name    string `json:"name"`
	}{
		Enabled: enabled,
		Key:     key,
		Name:    fmt.Sprintf("/Compute-%s/%s/%s", c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/%s%s",
		c.endpoint, "sshkey", ssh.Name)
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		body:   &ssh,
		url:    url,
		verb:   "PUT",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
