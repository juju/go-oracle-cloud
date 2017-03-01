// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// InstanceConsoleDetails retrieves the messages that
// appear when an instance boots. Use these messages
// to diagnose unresponsive instances and failures in
// the boot up process.
func (c *Client) InstanceConsoleDetails(
	name string,
) (resp response.InstanceConsole, err error) {

	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s%s", c.endpoints["instanceconsole"], name)

	if err = c.request(paramsRequest{
		url:  url,
		verb: "GET",
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
