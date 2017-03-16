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
	ipReservationParams = struct {
		Name       string
		Parentpool common.IPPool
		Permanent  bool
		Tags       []string
	}{
		Name:       "/Compute-acme/jack.jones@example.com/c7588bef-dab7-4b0d-adb4-648ba462c2c1",
		Parentpool: common.NewIPPool(common.PublicIPPool, common.IPPoolType),
		Permanent:  true,
		Tags:       nil,
	}

	ipReservationDetails = response.IpReservation{
		Account:    "/Compute-acme/default",
		Ip:         "129.144.53.38",
		Name:       "/Compute-acme/jack.jones@example.com/c7588bef-dab7-4b0d-adb4-648ba462c2c1",
		Parentpool: ipReservationParams.Parentpool,
		Permanent:  true,
		Used:       false,
	}

	allipreservations = response.AllIpReservations{
		Result: []response.IpReservation{
			ipReservationDetails,
		},
	}
)

func (cl clientTest) TestIpReservationWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpNetworkDetails(ipReservationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpReservations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpReservation(
		ipReservationParams.Name,
		ipReservationParams.Parentpool,
		ipReservationParams.Permanent,
		ipReservationParams.Tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpReservation(ipReservationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateIpReservation(
		ipReservationParams.Name,
		ipReservationParams.Name,
		ipReservationParams.Parentpool,
		ipReservationParams.Permanent,
		ipReservationParams.Tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestIpReservationWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpReservationDetails(ipReservationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpReservations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpReservation(
			ipReservationParams.Name,
			ipReservationParams.Parentpool,
			ipReservationParams.Permanent,
			ipReservationParams.Tags,
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
			err = client.DeleteIpReservation(ipReservationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateIpReservation(
			ipReservationParams.Name,
			ipReservationParams.Name,
			ipReservationParams.Parentpool,
			ipReservationParams.Permanent,
			ipReservationParams.Tags,
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

func (cl clientTest) TestIpReservationWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpReservationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip reservation name"), gc.Equals, true)

	_, err = client.CreateIpReservation("", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip reservation name"), gc.Equals, true)

	_, err = client.UpdateIpReservation("", "", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip reservation current name"), gc.Equals, true)

	err = client.DeleteIpReservation("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip reservation name"), gc.Equals, true)
}

func (cl clientTest) TestIpReservationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipReservationDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.IpReservationDetails(ipReservationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipReservationDetails)

}

func (cl clientTest) TestAllIpReservations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipreservations),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllIpReservations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipreservations)
}

func (cl clientTest) TestDeleteIpReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpReservation(ipReservationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipReservationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Permanent  bool          `json:"permanent"`
				Tags       []string      `json:"tags,omitempty"`
				Name       string        `json:"name"`
				Parentpool common.IPPool `json:"parentpool"`
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, ipReservationParams.Name)
			c.Assert(req.Parentpool, gc.DeepEquals, ipReservationParams.Parentpool)
			c.Assert(req.Permanent, gc.DeepEquals, ipReservationParams.Permanent)
			c.Assert(req.Tags, gc.DeepEquals, ipReservationParams.Tags)
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpReservation(
		ipReservationParams.Name,
		ipReservationParams.Parentpool,
		ipReservationParams.Permanent,
		ipReservationParams.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipReservationDetails)
}

func (cl clientTest) TestUpdateIpReservation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipReservationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Permanent  bool          `json:"permanent"`
				Tags       []string      `json:"tags,omitempty"`
				Name       string        `json:"name"`
				Parentpool common.IPPool `json:"parentpool"`
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, ipReservationParams.Name)
			c.Assert(req.Parentpool, gc.DeepEquals, ipReservationParams.Parentpool)
			c.Assert(req.Permanent, gc.DeepEquals, ipReservationParams.Permanent)
			c.Assert(req.Tags, gc.DeepEquals, ipReservationParams.Tags)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateIpReservation(
		ipReservationParams.Name,
		ipReservationParams.Name,
		ipReservationParams.Parentpool,
		ipReservationParams.Permanent,
		ipReservationParams.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipReservationDetails)
}
