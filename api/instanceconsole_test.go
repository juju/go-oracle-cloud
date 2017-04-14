// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
	gc "gopkg.in/check.v1"
)

var (
	instanceConsoleName    = "/Compute-acme/jack.jones@example.com/68a3c40c-466e-41df-a7f2-00fbfbd590e5"
	instanceConsoleDetails = response.InstanceConsole{
		Name:      instanceConsoleName,
		Timestamp: "2016-06-17T09:21:19.662570",
		Output:    "k [LNKD] (IRQs *5 10 11)\r\nvgaarb: device added: ... login: ",
		Uri:       "https://api-z999.compute.us0.oraclecloud.com/instanceconsole/Compute-acme/jack.jones@example.com/68a3c40c-466e-41df-a7f2-00fbfbd590e5",
	}
)

func (cl clientTest) TestInstanceConsoleResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.InstanceConsoleDetails(instanceConsoleName)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestInstanceConsoleWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.InstanceConsoleDetails(instanceConsoleName)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestInstanceConsoleResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.InstanceConsoleDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty instance console name"), gc.Equals, true)
}

func (cl clientTest) TestInstanceConsoleDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &instanceConsoleDetails),
		check: c,
		u: &unmarshaler{
			raw:  instanceConsoleDetailsRaw,
			into: &response.InstanceConsole{},
		},
	})
	defer ts.Close()

	resp, err := client.InstanceConsoleDetails(instanceConsoleName)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, instanceConsoleDetails)
}
