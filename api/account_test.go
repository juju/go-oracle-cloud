// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"fmt"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

func (cl clientTest) TestAccountDetailsWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	name := client.ComposeName("someName")
	_, err := client.AccountDetails(name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestAccountDetailsWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AccountDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty account name"), gc.Equals, true)
}

func (cl clientTest) TestAccountDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	name := client.ComposeName("someName")
	_, err := client.AccountDetails(name)
	c.Assert(err, gc.NotNil)
	fmt.Println(err.Error())
}
