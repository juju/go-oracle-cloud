// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

type SecRule string

const (
	DefaultSecRule SecRule = "PERMIT"
)

type SecRuleParams struct {
	permit      string `json:"permit"`
	application string `json:"application"`
	description string `json:"description,omitempty"`
	disabled    bool   `json:"disabled"`
	dst_list    string `json:"dst_list"`
	name        string `json:"name"`
	Src_list    string `json:"src_list"`
}

// func (s SecRuleParams) validate() (err error) {
// 	if permit.validate() == "" {
//
// 	}
// }
//
// // CreateSecRule
// func (c Client) CreateSecRule(p SecRuleParams) (resp response.SecRule, err error) {
// 	if !c.isAuth() {
// 		return resp, errNotAuth
// 	}
//
// }
