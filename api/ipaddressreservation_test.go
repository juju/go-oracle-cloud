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
	ipAddressReservationParams = api.IpAddressReservationParams{
		Description:   "",
		IpAddressPool: common.NewIPPool(common.PublicIPPool, common.IPPoolType),
		Name:          "/Compute-acme/jack.jones@example.com/ipreservation1",
		Tags:          nil,
	}

	ipAddressReservationDetails = response.IpAddressReservation{
		Name:          "/Compute-acme/jack.jones@example.com/ipreservation1",
		Uri:           "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipreservation/Compute-acme/jack.jones@example.com/ipreservation1",
		IpAddressPool: common.NewIPPool(common.PublicIPPool, common.IPPoolType),
		IpAddress:     "10.252.152.90",
	}

	allipaddressreservations = response.AllIpAddressReservations{
		Result: []response.IpAddressReservation{
			ipAddressReservationDetails,
		},
	}
)

func (cl clientTest) TestIpAddressReservationWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressReservationDetails(ipAddressReservationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpAddressReservations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpAddressReservation(ipAddressReservationParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpAddressReservation(ipAddressReservationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateIpAddressReservation(
		ipAddressReservationParams,
		ipAddressReservationParams.Name,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestIpAddressReservationWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpAddressReservationDetails(ipAddressReservationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpAddressReservations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpAddressReservation(ipAddressReservationParams)
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
			err = client.DeleteIpAddressReservation(ipAddressReservationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateIpAddressReservation(ipAddressReservationParams,
			ipAddressReservationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)
		ts.Close()
	}
}

func (cl clientTest) TestIpAddressReservationWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressReservationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address reservation name"), gc.Equals, true)

	_, err = client.CreateIpAddressReservation(api.IpAddressReservationParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address reservation name"), gc.Equals, true)

	err = client.DeleteIpAddressReservation("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address reservation name"), gc.Equals, true)

	_, err = client.UpdateIpAddressReservation(api.IpAddressReservationParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address reservation name"), gc.Equals, true)

	_, err = client.UpdateIpAddressReservation(ipAddressReservationParams, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address reservation current name"), gc.Equals, true)

}

func (cl clientTest) TestIpAddressReservationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressReservationDetails),
		check: c,
		u: &unmarshaler{
			raw:  ipAddressReservationRaw,
			into: &response.IpAddressReservation{},
		},
	})

	defer ts.Close()

	resp, err := client.IpAddressReservationDetails(ipAddressReservationDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressReservationDetails)

}

func (cl clientTest) TestAllIpAddressReservations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipaddressreservations),
		check: c,
		u: &unmarshaler{
			raw:  allIpAddressReservationRaw,
			into: &response.AllIpAddressReservations{},
		},
	})

	defer ts.Close()

	resp, err := client.AllIpAddressReservations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipaddressreservations)

}

func (cl clientTest) TestDeleteIpAddressReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpAddressReservation(ipAddressReservationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpAddressReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressReservationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.IpAddressReservationParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, ipAddressReservationParams)
		},
		u: &unmarshaler{
			raw:  ipAddressReservationRaw,
			into: &response.IpAddressReservation{},
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpAddressReservation(ipAddressReservationParams)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressReservationDetails)
}

func (cl clientTest) TestUpdateIpAddressReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressReservationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.IpAddressReservationParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, ipAddressReservationParams)
		},
		u: &unmarshaler{
			raw:  ipAddressReservationRaw,
			into: &response.IpAddressReservation{},
		},
	})
	defer ts.Close()

	resp, err := client.UpdateIpAddressReservation(ipAddressReservationParams,
		ipAddressReservationParams.Name)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressReservationDetails)
}
