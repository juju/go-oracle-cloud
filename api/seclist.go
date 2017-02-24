package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// CreatesSecList a security list. After creating security
// lists, you can add instances to them by using the HTTP request,
// CreateSecAssociation (Create a Security Association).
func (c Client) CreateSecList(
	description string,
	name string,
	outbound_cidr_policy string,
	policy string,
) (resp response.SecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty secure list name")
	}

	params := struct {
		Description          string `json:"description,omitempty"`
		Name                 string `json:"name"`
		Outbound_cidr_policy string `json:"outbound_cidr_policy"`
		Policy               string `json:"policy"`
	}{
		Description:          description,
		Name:                 name,
		Outbound_cidr_policy: outbound_cidr_policy,
		Policy:               policy,
	}

	url := fmt.Sprintf("%s/seclist/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSecList the specified security list. No response is returned.<Paste>
func (c Client) DeleteSecList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure list")
	}

	url := fmt.Sprintf("%s/seclist/Compute-%s/%s/%s",
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

// AllSecList retrieves details of the security lists that are in the specified
// container and match the specified query criteria.
func (c Client) AllSecList() (resp response.AllSecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/seclist/Compute-%s/%s/",
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

// SecListDetails retrieves information about the specified security list.
func (c Client) SecListDetails(name string) (resp response.SecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/seclist/Compute-%s/%s/%s",
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
