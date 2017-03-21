// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"

	gc "gopkg.in/check.v1"
)

var (
	shapeName    = "oc3"
	shapeDetails = response.Shape{
		Nds_iops_limit: 0,
		Ram:            7680,
		Cpus:           2.0,
		Root_disk_size: 0,
		Io:             200,
		Uri:            "https://api-z999.compute.us0.oraclecloud.com/shape/oc3",
		Name:           "oc3",
	}
	allshapes = response.AllShapes{Result: []response.Shape{shapeDetails}}
)

func (cl clientTest) TestShapeResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.ShapeDetails(shapeName)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllShapes(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestShapeWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.ShapeDetails(shapeName)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllShapes(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestShapeResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.ShapeDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty shape name"), gc.Equals, true)
}

func (cl clientTest) TestShapeDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &shapeDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.ShapeDetails(shapeName)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, shapeDetails)
}

func (cl clientTest) TestAllShapes(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allshapes),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllShapes(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allshapes)

}
