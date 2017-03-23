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
	vpnEndpointParams = api.VpnEndpointParams{
		Psk:  "********",
		Name: "/Compute-acme/vpn-to-LA",
		Reachable_routes: []string{
			"192.168.155.2/24",
			"192.168.135.0/24",
		},
		Enabled:              true,
		Customer_vpn_gateway: "192.168.111.2",
	}

	vpnEndpointDetails = response.VpnEndpoint{
		Psk:                  vpnEndpointParams.Psk,
		Name:                 vpnEndpointParams.Name,
		Reachable_routes:     vpnEndpointParams.Reachable_routes,
		Enabled:              vpnEndpointParams.Enabled,
		Customer_vpn_gateway: vpnEndpointParams.Customer_vpn_gateway,
		Uri:                  "http://api.oc.example.com/vpnendpoint/Compute-acme/vpn-to-LA",
		Status:               "UP",
	}

	allvpnendpoints = response.AllVpnEndpoints{
		Result: []response.VpnEndpoint{vpnEndpointDetails},
	}
)

func (cl clientTest) TestVpnEndpointWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VpnEndpointDetails(vpnEndpointParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllVpnEndpoints(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateVpnEndpoint(vpnEndpointParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteVpnEndpoint(vpnEndpointParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestVpnEndpointWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.VpnEndpointDetails(vpnEndpointParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllVpnEndpoints(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateVpnEndpoint(vpnEndpointParams)
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
			err = client.DeleteVpnEndpoint(vpnEndpointParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}
		ts.Close()
	}
}

func (cl clientTest) TestVpnEndpointWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VpnEndpointDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty vpn endpoint name"), gc.Equals, true)

	_, err = client.CreateVpnEndpoint(api.VpnEndpointParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty vpn endpoint name"), gc.Equals, true)

	err = client.DeleteVpnEndpoint("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty vpn endpoint name"), gc.Equals, true)
}

func (cl clientTest) TestVpnEndpointDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &vpnEndpointDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.VpnEndpointDetails(vpnEndpointParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, vpnEndpointDetails)

}

func (cl clientTest) TestAllVpnEndpoints(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allvpnendpoints),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllVpnEndpoints(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allvpnendpoints)

}

func (cl clientTest) TestDeleteVpnEndpoint(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteVpnEndpoint(vpnEndpointParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateVpnEndpoint(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &vpnEndpointDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.VpnEndpointParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, vpnEndpointParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateVpnEndpoint(vpnEndpointParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, vpnEndpointDetails)
}
