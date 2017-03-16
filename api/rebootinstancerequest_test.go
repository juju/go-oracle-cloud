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
	rebootInstanceRequestParams = struct {
		Hard         bool
		InstanceName string
	}{
		Hard:         true,
		InstanceName: "/Compute-acme/jack.jones@example.com/7b036196-d077-4b6a-b2c8-64bc9d8d97ce",
	}

	rebootInstanceRequestDetails = response.RebootInstanceRequest{
		Name:          rebootInstanceRequestParams.InstanceName,
		Hard:          true,
		Creation_time: "2016-05-10T16:08:10Z",
		Error_reason:  "",
		Uri:           "https://api-z999.compute.us0.oraclecloud.com/rebootinstancerequest/Compute-acme/jack.jones@example.com/7b036196-d077-4b6a-b2c8-64bc9d8d97ce/8666a368-6264-4cee-9ec4-63996234deb1",
		Instance_id:   "7b036196-d077-4b6a-b2c8-64bc9d8d97ce",
		Instance:      "/Compute-acme/jack.jones@example.com/7b036196-d077-4b6a-b2c8-64bc9d8d97ce",
		State:         "active",
		Request_id:    "8666a368-6264-4cee-9ec4-63996234deb1",
	}

	allrebootinstancerequests = response.AllRebootInstanceRequests{
		Result: []response.RebootInstanceRequest{
			rebootInstanceRequestDetails,
		},
	}
)

func (cl clientTest) TestRebootInstanceRequestWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.RebootInstanceRequestDetails(rebootInstanceRequestParams.InstanceName)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllRebootInstanceRequests(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateRebootInstanceRequest(
		rebootInstanceRequestParams.Hard,
		rebootInstanceRequestParams.InstanceName,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteRebootInstanceRequest(rebootInstanceRequestParams.InstanceName)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestRebootInstanceRequestWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.RebootInstanceRequestDetails(
			rebootInstanceRequestParams.InstanceName,
		)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllRebootInstanceRequests(nil)
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
			err = client.DeleteRebootInstanceRequest(
				rebootInstanceRequestParams.InstanceName,
			)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.CreateRebootInstanceRequest(
			rebootInstanceRequestParams.Hard,
			rebootInstanceRequestParams.InstanceName,
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

func (cl clientTest) TestInstanceRebootRequestResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.RebootInstanceRequestDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty reboot instance request name"), gc.Equals, true)

	_, err = client.CreateRebootInstanceRequest(rebootInstanceRequestParams.Hard, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty reboot instance request name"), gc.Equals, true)

	err = client.DeleteRebootInstanceRequest("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty reboot instance request name"), gc.Equals, true)
}

func (cl clientTest) TestRebootInstanceRequest(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &rebootInstanceRequestDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.RebootInstanceRequestDetails(
		rebootInstanceRequestParams.InstanceName,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, rebootInstanceRequestDetails)
}

func (cl clientTest) TestAllRebootInstanceRequests(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allrebootinstancerequests),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllRebootInstanceRequests(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allrebootinstancerequests)
}

func (cl clientTest) TestDeleteRebootInstanceRequest(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteRebootInstanceRequest(
		rebootInstanceRequestParams.InstanceName,
	)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateRebootInstancerequests(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &rebootInstanceRequestDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name string `json:"name"`
				Hard bool   `json:"hard"`
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Hard, gc.DeepEquals,
				rebootInstanceRequestParams.Hard)
			c.Assert(req.Name, gc.DeepEquals,
				rebootInstanceRequestParams.InstanceName)
		},
	})
	defer ts.Close()

	resp, err := client.CreateRebootInstanceRequest(
		rebootInstanceRequestParams.Hard,
		rebootInstanceRequestParams.InstanceName,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, rebootInstanceRequestDetails)
}
