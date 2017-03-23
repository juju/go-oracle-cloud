// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	enc "encoding/json"
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/common"
	"github.com/juju/go-oracle-cloud/response"

	gc "gopkg.in/check.v1"
)

var (
	securityProtocolParams = api.SecurityProtocolParams{
		Description: "Sample security protocol",
		IpProtocol:  common.TCP,
		SrcPortSet:  []string{"1018-1040"},
		DstPortSet:  []string{"2018-2040"},
		Name:        "/Compute-acme/jack.jones@example.com/secprotocol1",
	}

	securityProtocolDetails = response.SecurityProtocol{
		Uri:         "https://api-z999.compute.us0.oraclecloud.com/network/v1/secprotocol/Compute-acme/jack.jones@example.com/secprotocol1",
		Description: securityProtocolParams.Description,
		IpProtocol:  securityProtocolParams.IpProtocol,
		SrcPortSet:  securityProtocolParams.SrcPortSet,
		DstPortSet:  securityProtocolParams.DstPortSet,
		Name:        securityProtocolParams.Name,
	}

	allsecurityprotocols = response.AllSecurityProtocols{
		Result: []response.SecurityProtocol{
			securityProtocolDetails,
		},
	}
)

func (cl clientTest) TestSecurityProtocolResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllSecurityProtocols(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecurityProtocol(securityProtocolParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecurityProtocol(securityProtocolParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSecurityProtocol(securityProtocolParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.SecurityProtocolDetails(securityProtocolParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecurityProtocolResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecurityProtocol(securityProtocolParams)
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
			err = client.DeleteSecurityProtocol(securityProtocolParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecurityProtocols(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecurityProtocolDetails(securityProtocolParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateSecurityProtocol(securityProtocolParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecurityProtocolResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecurityProtocolDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security protocol name"), gc.Equals, true)

	_, err = client.CreateSecurityProtocol(api.SecurityProtocolParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security protocol name"), gc.Equals, true)

	err = client.DeleteSecurityProtocol("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security protocol name"), gc.Equals, true)

	_, err = client.UpdateSecurityProtocol(api.SecurityProtocolParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security protocol name"), gc.Equals, true)

}

func (cl clientTest) TestSecurityProtocolDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityProtocolDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecurityProtocolDetails(securityProtocolParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityProtocolDetails)
}

func (cl clientTest) TestAllSecurityProtocols(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecurityprotocols),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecurityProtocols(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecurityprotocols)
}

func (cl clientTest) TestDeleteSecurityProtocol(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecurityProtocol(securityProtocolParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecurityProtocol(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityProtocolDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecurityProtocolParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, securityProtocolParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecurityProtocol(securityProtocolParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityProtocolDetails)
}

func (cl clientTest) TestUpdateSecurityProtocol(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityProtocolDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecurityProtocolParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, securityProtocolParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateSecurityProtocol(securityProtocolParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityProtocolDetails)
}
