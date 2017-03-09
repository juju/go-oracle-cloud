// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
	gc "gopkg.in/check.v1"
)

func (cl clientTest) TestAccountResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	name := client.ComposeName("someName")
	_, err := client.AccountDetails(name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllAccounts(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllAccountNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.DirectoryAccount()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestAccountResourceWithErrors(c *gc.C) {

	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		name := client.ComposeName("someName")
		_, err := client.AccountDetails(name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllAccounts(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllAccountNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.DirectoryAccount()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
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

var (
	accountDetails = response.Account{
		Uri:  "https://api-z999.compute.us0.oraclecloud.com/account/Compute-acme/cloud_storage",
		Name: "/Compute-acme/cloud_storage",
	}

	allaccounts = response.AllAccounts{
		Result: []response.Account{
			accountDetails,
		},
	}
	accountnames = response.DirectoryNames{
		Result: []string{
			fmt.Sprintf("/Compute-%s/", identify),
			"/Compute-public/",
		},
	}
)

func (cl clientTest) TestAccountDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &accountDetails),
		check: c,
	})
	defer ts.Close()

	name := client.ComposeName("someName")
	resp, err := client.AccountDetails(name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, accountDetails)
}

func (cl clientTest) TestAllAccounts(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allaccounts),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllAccounts(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allaccounts)
}

func (cl clientTest) TestAccountNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &accountnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllAccountNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, accountnames)
}

func (cl clientTest) TestDirectoryAccount(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &accountnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.DirectoryAccount()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, accountnames)
}
