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
	ipNetworkExchangeParams = struct {
		Description string
		Name        string
		Tags        []string
	}{
		Name:        "/Compute-acme/jack.jones@example.com/ipNetworkExchange1",
		Tags:        nil,
		Description: "",
	}

	ipNetworkExchangeDetails = response.IpNetworkExchange{
		Name: "/Compute-acme/jack.jones@example.com/ipNetworkExchange1",
		Uri:  "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/ipnetworkexchange/Compute-acme/jack.jones@example.com/ipNetworkExchange1",
	}

	allipnetworkexchanges = response.AllIpNetworkExchanges{
		Result: []response.IpNetworkExchange{
			ipNetworkExchangeDetails,
		},
	}
)

func (cl clientTest) TestIpNetworkExchangeWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpNetworkExchangeDetails(ipNetworkExchangeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpNetworkExchanges(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpNetworkExchange(
		ipNetworkExchangeParams.Description,
		ipNetworkExchangeParams.Name,
		ipNetworkExchangeParams.Tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpNetworkExchange(ipNetworkExchangeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestIpNetworkExchangeWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpNetworkExchangeDetails(ipNetworkExchangeParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpNetworkExchanges(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpNetworkExchange(
			ipNetworkExchangeParams.Description,
			ipNetworkExchangeParams.Name,
			ipNetworkExchangeParams.Tags,
		)
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
			err = client.DeleteIpNetworkExchange(ipNetworkExchangeParams.Name)
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

func (cl clientTest) TestIpNetworkExchangeWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpNetworkExchangeDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network exchange name"), gc.Equals, true)

	_, err = client.CreateIpNetworkExchange("", "", nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network exchange name"), gc.Equals, true)

	err = client.DeleteIpNetworkExchange("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network exchange name"), gc.Equals, true)
}

func (cl clientTest) TestIpNetworkExchangeDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipNetworkExchangeDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.IpNetworkExchangeDetails(ipNetworkExchangeParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipNetworkExchangeDetails)

}

func (cl clientTest) TestAllIpNetworkExchanges(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipnetworkexchanges),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllIpNetworkExchanges(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipnetworkexchanges)

}

func (cl clientTest) TestDeleteIpNetworkExchange(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpNetworkExchange(ipNetworkExchangeParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpNetworkExchange(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipNetworkExchangeDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Description string
				Name        string
				Tags        []string
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Description, gc.DeepEquals, ipNetworkExchangeParams.Description)
			c.Assert(req.Name, gc.DeepEquals, ipNetworkExchangeParams.Name)
			c.Assert(req.Tags, gc.DeepEquals, ipNetworkExchangeParams.Tags)
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpNetworkExchange(
		ipNetworkExchangeParams.Description,
		ipNetworkExchangeParams.Name,
		ipNetworkExchangeParams.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipNetworkExchangeDetails)
}
