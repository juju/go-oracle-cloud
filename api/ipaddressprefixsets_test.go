// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	enc "encoding/json"
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"

	gc "gopkg.in/check.v1"
)

var (
	ipAddressPrefixSetParams = api.IpAddressPrefixSetParams{
		Name:              "/Compute-acme/jack.jones@example.com/ipaddressprefixset1",
		IpAddressPrefixes: []string{"192.168.0.0/16"},
	}

	ipAddressPrefixSetDetails = response.IpAddressPrefixSet{
		Name:              "/Compute-acme/jack.jones@example.com/ipaddressprefixset1",
		Uri:               "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipaddressprefixset/Compute-acme/jack.jones@example.com/ipaddressprefixset1",
		IpAddressPrefixes: ipAddressPrefixSetParams.IpAddressPrefixes,
	}

	allipaddressprefixsets = response.AllIpAddressPrefixSets{
		Result: []response.IpAddressPrefixSet{
			ipAddressPrefixSetDetails,
		},
	}
)

func (cl clientTest) TestIpAddressPrefixsetWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressPrefixSetDetails(ipAddressPrefixSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpAddressPrefixSets(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpAddressPrefixSet(ipAddressPrefixSetParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpAddressPrefixSet(ipAddressPrefixSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateIpAddressPrefixSet(ipAddressPrefixSetParams, ipAddressPrefixSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestIpAddressPrefixSetWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpAddressPrefixSetDetails(ipAddressPrefixSetParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpAddressPrefixSets(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpAddressPrefixSet(ipAddressPrefixSetParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		// when we encounter this case that,[
		// the delete method is recivng http.StatusNotFound
		// this means for the delete resource point of view to not be
		// an acutal error and it will return nil so we don't need to check this
		if key != http.StatusNotFound {
			err = client.DeleteIpAddressPrefixSet(ipAddressPrefixSetParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateIpAddressPrefixSet(
			ipAddressPrefixSetParams,
			ipAddressPrefixSetParams.Name,
		)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)
		ts.Close()
	}
}

func (cl clientTest) TestIpAddressPrefixSetWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressPrefixSetDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address prefix set name"), gc.Equals, true)

	_, err = client.CreateIpAddressPrefixSet(api.IpAddressPrefixSetParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address prefix set name"), gc.Equals, true)

	err = client.DeleteIpAddressPrefixSet("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address prefix set name"), gc.Equals, true)

	_, err = client.UpdateIpAddressPrefixSet(api.IpAddressPrefixSetParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address prefix set name"), gc.Equals, true)

	_, err = client.UpdateIpAddressPrefixSet(ipAddressPrefixSetParams, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address prefix set current name"), gc.Equals, true)

}

func (cl clientTest) TestIpAddressPrefixSetDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressPrefixSetDetails),
		check: c,
		u: &unmarshaler{
			raw:  ipAddressPrefixSetsRaw,
			into: &response.IpAddressPrefixSet{},
		},
	})

	defer ts.Close()

	resp, err := client.IpAddressPrefixSetDetails(ipAddressPrefixSetDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressPrefixSetDetails)
}

func (cl clientTest) TestAllIpAddressPrefixSets(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipaddressprefixsets),
		check: c,
		u: &unmarshaler{
			raw:  allIpAddressPrefixSetsRaw,
			into: &response.AllIpAddressPrefixSets{},
		},
	})

	defer ts.Close()

	resp, err := client.AllIpAddressPrefixSets(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipaddressprefixsets)

}

func (cl clientTest) TestDeleteIpAddressPrefixSet(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpAddressPrefixSet(ipAddressPrefixSetDetails.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpAddressPrefixset(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressPrefixSetDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.IpAddressPrefixSetParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, ipAddressPrefixSetParams)
		},
		u: &unmarshaler{
			raw:  ipAddressPrefixSetsRaw,
			into: &response.IpAddressPrefixSet{},
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpAddressPrefixSet(ipAddressPrefixSetParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressPrefixSetDetails)
}

func (cl clientTest) TestUpdateIpAddressPrefixSet(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressPrefixSetDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.IpAddressPrefixSetParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, ipAddressPrefixSetParams)
		},
		u: &unmarshaler{
			raw:  ipAddressPrefixSetsRaw,
			into: &response.IpAddressPrefixSet{},
		},
	})
	defer ts.Close()

	resp, err := client.UpdateIpAddressPrefixSet(ipAddressPrefixSetParams,
		ipAddressPrefixSetParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressPrefixSetDetails)
}
