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
	ipNetworkParams = struct {
		Name                  string
		Description           string
		IpNetworkExchange     string
		IpAddressPrefix       string
		PublicNaptEnabledFlag bool
		Tags                  []string
	}{
		Name:            "/Compute-acme/jack.jones@example.com/ipnet1",
		IpAddressPrefix: "192.168.0.0/24",
	}

	ipNetworkDetails = response.IpNetwork{
		Name:                  "/Compute-acme/jack.jones@example.com/ipnet1",
		Uri:                   "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/ipnetwork/Compute-acme/jack.jones@example.com/ipnet1",
		Description:           nil,
		Tags:                  nil,
		IpAddressPrefix:       "192.168.0.0/24",
		IpNetworkExchange:     nil,
		PublicNaptEnabledFlag: false,
	}

	allipnetworks = response.AllIpNetworks{
		Result: []response.IpNetwork{
			ipNetworkDetails,
		},
	}
)

func (cl clientTest) TestIpNetworkWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpNetworkDetails(ipNetworkParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpNetworks(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpNetwork(
		ipNetworkParams.Description,
		ipNetworkParams.IpAddressPrefix,
		ipNetworkParams.IpNetworkExchange,
		ipNetworkParams.Name,
		ipNetworkParams.PublicNaptEnabledFlag,
		ipNetworkParams.Tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpNetwork(ipNetworkParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateIpNetwork(
		ipNetworkParams.Name,
		ipNetworkParams.Name,
		ipNetworkParams.Description,
		ipNetworkParams.IpNetworkExchange,
		ipNetworkParams.IpAddressPrefix,
		ipNetworkParams.PublicNaptEnabledFlag,
		ipNetworkParams.Tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestIpNetworkWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpNetworkDetails(ipNetworkParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpNetworks(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpNetwork(
			ipNetworkParams.Description,
			ipNetworkParams.IpAddressPrefix,
			ipNetworkParams.IpNetworkExchange,
			ipNetworkParams.Name,
			ipNetworkParams.PublicNaptEnabledFlag,
			ipNetworkParams.Tags,
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
			err = client.DeleteIpNetwork(ipNetworkParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateIpNetwork(
			ipNetworkParams.Name,
			ipNetworkParams.Name,
			ipNetworkParams.Description,
			ipNetworkParams.IpNetworkExchange,
			ipNetworkParams.IpAddressPrefix,
			ipNetworkParams.PublicNaptEnabledFlag,
			ipNetworkParams.Tags,
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

func (cl clientTest) TestIpNetworkWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpNetworkDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network name"), gc.Equals, true)

	_, err = client.CreateIpNetwork("", "", "", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network name"), gc.Equals, true)

	_, err = client.UpdateIpNetwork("", "", "", "", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network current name"), gc.Equals, true)

	_, err = client.UpdateIpNetwork(ipNetworkParams.Name, "", "", "", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network address prefix"), gc.Equals, true)

	err = client.DeleteIpNetwork("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip network name"), gc.Equals, true)
}

func (cl clientTest) TestIpNetworkDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipNetworkDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.IpNetworkDetails(ipNetworkParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipNetworkDetails)

}

func (cl clientTest) TestAllIpNetworks(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipnetworks),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllIpNetworks(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipnetworks)

}

func (cl clientTest) TestDeleteIpNetwork(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpNetwork(ipNetworkParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpNetwork(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipNetworkDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name                  string
				Description           string
				IpNetworkExchange     string
				IpAddressPrefix       string
				PublicNaptEnabledFlag bool
				Tags                  []string
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, ipNetworkDetails.Name)
			c.Assert(req.IpAddressPrefix, gc.DeepEquals, ipNetworkDetails.IpAddressPrefix)
			c.Assert(req.Tags, gc.DeepEquals, ipNetworkDetails.Tags)
			c.Assert(req.PublicNaptEnabledFlag, gc.DeepEquals, ipNetworkDetails.PublicNaptEnabledFlag)

		},
	})
	defer ts.Close()

	resp, err := client.CreateIpNetwork(
		ipNetworkParams.Description,
		ipNetworkParams.IpAddressPrefix,
		ipNetworkParams.IpNetworkExchange,
		ipNetworkParams.Name,
		ipNetworkParams.PublicNaptEnabledFlag,
		ipNetworkParams.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipNetworkDetails)
}

func (cl clientTest) TestUpdateIpNetwork(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipNetworkDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name                  string
				Description           string
				IpNetworkExchange     string
				IpAddressPrefix       string
				PublicNaptEnabledFlag bool
				Tags                  []string
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, ipNetworkDetails.Name)
			c.Assert(req.IpAddressPrefix, gc.DeepEquals, ipNetworkDetails.IpAddressPrefix)
			c.Assert(req.Tags, gc.DeepEquals, ipNetworkDetails.Tags)
			c.Assert(req.PublicNaptEnabledFlag, gc.DeepEquals, ipNetworkDetails.PublicNaptEnabledFlag)

		},
	})
	defer ts.Close()

	resp, err := client.UpdateIpNetwork(
		ipNetworkParams.Name,
		ipNetworkParams.Name,
		ipNetworkParams.Description,
		ipNetworkParams.IpNetworkExchange,
		ipNetworkParams.IpAddressPrefix,
		ipNetworkParams.PublicNaptEnabledFlag,
		ipNetworkParams.Tags,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipNetworkDetails)
}
