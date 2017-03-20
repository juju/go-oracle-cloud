// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	enc "encoding/json"
	"net/http"
	"strings"

	gc "gopkg.in/check.v1"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
)

var (
	orchestrationParams = api.OrchestrationParams{
		Relationships: nil,
		Description:   "simple orchestration",
		Name:          "/acme/jack.jones@example.com/webportal",
		Oplans: []api.Oplan{
			{
				Ha_policy: "active",
				Label:     "webserver_ip",
				Obj_type:  "launchplan",
				Objects: []api.Object{
					{
						Instances: []api.InstancesOrchestration{
							{
								Shape:     "medium",
								Label:     "primary_webserver",
								Imagelist: "/oracle/public/compute_oel_6.4_2GB",
								Name:      "/acme/jack.jones@example.com/primary_webserver",
							},
						},
					},
				},
			},
		},
	}

	orchestrationDetails = response.Orchestration{
		Relationships: nil,
		Status:        "stopped",
		Account:       "/acme/default",
		Description:   "simple orchestration",
		Uri:           "https://api-z999.compute.us0.oraclecloud.com/orchestration/Compute-acme/jack.jones@example.com/webportal",
		Oplans: []response.Oplan{
			{
				Status:    "stopped",
				Obj_type:  "launchplan",
				Ha_policy: "active",
				Label:     "webserver_Lp",
				Objects: []response.Object{
					{
						Instances: []response.InstancesOrchestration{
							{
								Shape:     "medium",
								Label:     "primary_webserver",
								Imagelist: "/oracle/public/compute_oel_6.4_2GB",
								Name:      "/acme/jack.jones@example.com/primary_webserver",
							},
						},
					},
				},
			},
		},
	}

	allorchestrations = response.AllOrchestrations{
		Result: []response.Orchestration{
			orchestrationDetails,
		},
	}

	orchestrationnames = response.DirectoryNames{
		Result: []string{
			"/acme/jack.jones@example.com/",
			"/acme/jack.jones@example.com/",
		},
	}
)

func (cl clientTest) TestOrchestrationResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllOrchestrations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.OrchestrationDetails(orchestrationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteOrchestration(orchestrationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateOrchestration(orchestrationParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateOrchestration(orchestrationParams, orchestrationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllOrchestrationNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.DirectoryOrchestration()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestOrchestrationWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.OrchestrationDetails(orchestrationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllOrchestrations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		// when we encounter this case that,
		// the delete method is recivng http.StatusNotFound
		// this means for the delete resource point of view to not be
		// an acutal error and it will return nil so we don't need to check this
		if key != http.StatusNotFound {
			err = client.DeleteOrchestration(orchestrationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.CreateOrchestration(orchestrationParams)
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateOrchestration(orchestrationParams, orchestrationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllOrchestrationNames()
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.DirectoryOrchestration()
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestOrchestrationResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.OrchestrationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty orchestration name"), gc.Equals, true)

	err = client.DeleteOrchestration("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty orchestration name"), gc.Equals, true)

	_, err = client.CreateOrchestration(api.OrchestrationParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty orchestration name"), gc.Equals, true)

	_, err = client.UpdateOrchestration(api.OrchestrationParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty orchestration name"), gc.Equals, true)
}

func (cl clientTest) TestOrchestrationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &orchestrationDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.OrchestrationDetails(orchestrationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, orchestrationDetails)
}

func (cl clientTest) TestAllOrchestrations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allorchestrations),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllOrchestrations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allorchestrations)
}

func (cl clientTest) TestDeleteOrchestraiton(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteOrchestration(orchestrationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestAllOrchestrationNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &orchestrationnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllOrchestrationNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, orchestrationnames)
}

func (cl clientTest) TestDirectoryOrchestration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &orchestrationnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.DirectoryOrchestration()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, orchestrationnames)
}

func (cl clientTest) TestCreateOrchestration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &orchestrationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.Orchestration
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)

			c.Assert(req.Name, gc.Equals, orchestrationParams.Name)
			c.Assert(req.Description, gc.Equals, orchestrationParams.Description)

		},
	})
	defer ts.Close()

	resp, err := client.CreateOrchestration(orchestrationParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, orchestrationDetails)

}

func (cl clientTest) TestUpdateOrchestration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &orchestrationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.Orchestration
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)

			c.Assert(req.Name, gc.Equals, orchestrationParams.Name)
			c.Assert(req.Description, gc.Equals, orchestrationParams.Description)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateOrchestration(orchestrationParams, orchestrationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, orchestrationDetails)

}
