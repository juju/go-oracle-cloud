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
	virtualNicDetails = response.VirtualNic{
		Name:        "/Compute-acme/jack.jones@example.com/inst1-vnic1-ipnet1",
		Uri:         "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/vnic/Compute-acme/jack.jones@example.com/inst1-vnic1-ipnet1",
		Description: nil,
		Tags:        nil,
		MacAddress:  "c6:b0:c0:a8:03:0b",
		TransitFlag: false,
	}

	allvirtualnics = response.AllVirtualNics{
		Result: []response.VirtualNic{virtualNicDetails},
	}
)

func (cl clientTest) TestVirtualNicResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VirtualNicDetails(virtualNicDetails.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllVirtualNics(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestVirtualNicWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.VirtualNicDetails(virtualNicDetails.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllVirtualNics(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestVirtualNicResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VirtualNicDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty virtual nic name"), gc.Equals, true)
}

func (cl clientTest) TestVirtualNicDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &virtualNicDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.VirtualNicDetails(virtualNicDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, virtualNicDetails)
}

func (cl clientTest) TestAllVirtualNics(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allvirtualnics),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllVirtualNics(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allvirtualnics)
}
