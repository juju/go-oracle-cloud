// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/common"
	"github.com/hoenirvili/go-oracle-cloud/response"
)

type SecRuleParams struct {
	Permit      common.SecRule `json:"permit"`
	Application string         `json:"application"`
	Description string         `json:"description,omitempty"`
	Disabled    bool           `json:"disabled"`
	Dst_list    string         `json:"dst_list"`
	Name        string         `json:"name"`
	Src_list    string         `json:"src_list"`
}

func (s SecRuleParams) Validate() (err error) {
	if err = s.Permit.Validate(); err != nil {
		return err
	}

	if s.Application == "" {
		return errors.New("go-oracle-cloud: Empty application field")
	}

	if s.Name == "" {
		return errors.New("go-oracle-cloud: Empty secure rule name")
	}

	if s.Src_list == "" {
		return errors.New("go-oracle-cloud: Empty source list in secure rule")
	}

	if s.Dst_list == "" {
		return errors.New("go-oracle-cloud: Empty destination list in secure rule")
	}

	return nil
}

// CreateSecRule creates a new security rule. A security rule defines network access over a specified
// protocol between instances in two security lists, or from a
// set of external hosts (an IP list) to instances in a security list.
func (c *Client) CreateSecRule(p SecRuleParams) (resp response.SecRule, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if err = p.Validate(); err != nil {
		return resp, err
	}

	url := c.endpoints["secrule"] + "/"

	if err = c.request(paramsRequest{
		url:  url,
		body: &p,
		verb: "POST",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSecRule deletes a security role inside the oracle
// cloud account. If the security rule is not found this will return nil
func (c *Client) DeleteSecRule(name string) (err error) {
	if !c.isAuth() {
		return errNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure rule name")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["secrule"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "DELETE",
	}); err != nil {
		return err
	}

	return nil
}

// SecRuleDetails retrives details on a specific security rule
func (c *Client) SecRuleDetails(name string) (resp response.SecRule, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty secure rule name")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["secrule"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "GET",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// AllSecRules retrives all security rulues from the oracle cloud account
func (c *Client) AllSecRules() (resp response.AllSecRules, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["secrule"], c.identify, c.username)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "GET",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil

}

// UpdateSecRule modifies the security rule with the currentName
func (c *Client) UpdateSecRule(p SecRuleParams, currentName string) (resp response.SecRule, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if err = p.Validate(); err != nil {
		return resp, err
	}

	if currentName == "" {
		errors.New("go-oracle-cloud: Empty secure rule current name")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["secrule"], currentName)

	if err = c.request(paramsRequest{
		url:  url,
		body: &p,
		verb: "PUT",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// SecRuleNames retrives all secure rule names in the oracle cloud account
func (c *Client) SecRuleNames() (resp response.DirectoryNames, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/Compute-%s/%s/",
		c.endpoints["secrule"], c.identify, c.username)

	if err = c.request(paramsRequest{
		directory: true,
		url:       url,
		verb:      "GET",
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
