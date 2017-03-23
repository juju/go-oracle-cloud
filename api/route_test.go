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
	routeParams = api.RouteParams{
		Name:            "/Compute-acme/jack.jones@example.com/route1",
		AdminDistance:   common.AdminDistanceZero,
		IpAddressPrefix: "10.33.88.0/24",
		NextHopVnicSet:  "/Compute-acme/jack.jones@example.com/vnicset1",
	}

	routeDetails = response.Route{
		Name:            "/Compute-acme/jack.jones@example.com/route1",
		Uri:             "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/route/Compute-acme/jack.jones@example.com/route1",
		NextHopVnicSet:  "/Compute-acme/jack.jones@example.com/vnicset1",
		IpAddressPrefix: "10.33.88.0/24",
		AdminDistance:   common.AdminDistanceZero,
	}

	allroutes = response.AllRoutes{
		Result: []response.Route{
			routeDetails,
		},
	}
)

func (cl clientTest) TestRouteResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.RouteDetails(routeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllRoutes(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteRoute(routeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateRoute(routeParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateRoute(routeParams, routeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestRouteWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.RouteDetails(routeParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllRoutes(nil)
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
			err = client.DeleteRoute(routeParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.CreateRoute(routeParams)
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateRoute(routeParams, routeParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestRouteResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.RouteDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty route name"), gc.Equals, true)

	_, err = client.CreateRoute(api.RouteParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty route name"), gc.Equals, true)

	err = client.DeleteRoute("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty route name"), gc.Equals, true)

	_, err = client.UpdateRoute(routeParams, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty route current name"), gc.Equals, true)

	_, err = client.UpdateRoute(api.RouteParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty route name"), gc.Equals, true)
}

func (cl clientTest) TestRouteDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &routeDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.RouteDetails(routeParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, routeDetails)
}

func (cl clientTest) TestAllRoutes(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allroutes),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllRoutes(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allroutes)
}

func (cl clientTest) TestDeleteRoute(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteRoute(routeParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateRoute(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &routeDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.RouteParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, routeParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateRoute(routeParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, routeDetails)

}

func (cl clientTest) TestUpdateRoute(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &routeDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.RouteParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, routeParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateRoute(routeParams, routeParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, routeDetails)

}
