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
	secApplicationParams = api.SecApplicationParams{
		Dport:    "70",
		Name:     "/Compute-acme/jack.jones@example.com/vid_stream_udp",
		Protocol: common.TCP,
	}

	secApplicationDetails = response.SecApplication{
		Dport:    "70",
		Name:     "/Compute-acme/jack.jones@example.com/vid_stream_udp",
		Uri:      "https://api-z999.compute.us0.oraclecloud.com/secapplication/Compute-acme/jack.jones@example.com/vid_stream_udp",
		Protocol: common.TCP,
	}

	allsecapplications = response.AllSecApplications{
		Result: []response.SecApplication{
			secApplicationDetails,
		},
	}
)

func (cl clientTest) TestSecApplicationResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecApplicationDetails(secApplicationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllSecApplications(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecApplication(secApplicationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecApplication(secApplicationParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.DefaultSecApplications(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecApplicationResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecApplication(secApplicationParams)
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
			err = client.DeleteSecApplication(secApplicationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecApplications(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecApplicationDetails(secApplicationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.DefaultSecApplications(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecApplicationResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecApplicationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure application name"), gc.Equals, true)

	_, err = client.CreateSecApplication(api.SecApplicationParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure application name"), gc.Equals, true)

	err = client.DeleteSecApplication("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure application name"), gc.Equals, true)

}

func (cl clientTest) TestSecApplicationDetials(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secApplicationDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecApplicationDetails(secApplicationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secApplicationDetails)
}

func (cl clientTest) TestAllSecApplications(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecapplications),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecApplications(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecapplications)
}

func (cl clientTest) TestDefaultSecApplication(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecapplications),
		check: c,
	})
	defer ts.Close()

	resp, err := client.DefaultSecApplications(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecapplications)
}

func (cl clientTest) TestDeleteSecApplication(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecApplication(secApplicationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecApplication(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secApplicationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecApplicationParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, secApplicationParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecApplication(secApplicationParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secApplicationDetails)
}
