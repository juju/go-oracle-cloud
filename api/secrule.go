// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

type SecRule string

func (s SecRule) validate() (err error) {
	if s == "" {
		return errors.New("go-oracle-cloud: Empty secure rule permission")
	}

	return nil
}

const (
	DefaultSecRule SecRule = "PERMIT"
)

type SecRuleParams struct {
	Permit      SecRule `json:"permit"`
	Application string  `json:"application"`
	Description string  `json:"description,omitempty"`
	Disabled    bool    `json:"disabled"`
	Dst_list    string  `json:"dst_list"`
	Name        string  `json:"name"`
	Src_list    string  `json:"src_list"`
}

func (s SecRuleParams) validate() (err error) {
	if err = s.Permit.validate(); err != nil {
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
func (c Client) CreateSecRule(p SecRuleParams) (resp response.SecRule, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if err = p.validate(); err != nil {
		return resp, err
	}

	url := c.endpoints["secrule"] + "/"

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "POST",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSecRule deletes a security role inside the oracle
// cloud account. If the security rule is not found this will return nil
func (c Client) DeleteSecRule(name string) (err error) {
	if !c.isAuth() {
		return errNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty secure rule name")
	}

	url := fmt.Sprintf("%s%s", c.endpoints["secrule"], name)

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
