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
	vnicSetParams = api.VnicSetParams{
		Name: "/Compute-acme/jack.jones@example.com/vnicset_new_instance",
	}

	vnicSetDetails = response.VnicSet{
		Name:        vnicSetParams.Name,
		Uri:         "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/vnicset/Compute-acme/jack.jones@example.com/vnicset_new_instance",
		Description: nil,
		Tags:        nil,
		Vnics:       nil,
		AppliedAcls: nil,
	}

	allvnicsets = response.AllVnicSets{
		Result: []response.VnicSet{vnicSetDetails},
	}
)

func (cl clientTest) TestVnicSetWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VnicSetDetails(vnicSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllVnicSets(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateVnicSet(vnicSetParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteVnicSet(vnicSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateVnicSet(vnicSetParams, vnicSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestVnicSetWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.VnicSetDetails(vnicSetParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllVnicSets(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateVnicSet(vnicSetParams)
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
			err = client.DeleteVnicSet(vnicSetParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateVnicSet(vnicSetParams, vnicSetParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestVnicSetWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.VnicSetDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty virtual nic set name"), gc.Equals, true)

	_, err = client.CreateVnicSet(api.VnicSetParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty virtual nic set name"), gc.Equals, true)

	_, err = client.UpdateVnicSet(api.VnicSetParams{}, vnicSetParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty virtual nic set name"), gc.Equals, true)

	err = client.DeleteVnicSet("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty virtual nic set name"), gc.Equals, true)
}

func (cl clientTest) TestVnicSetDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &vnicSetDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.VnicSetDetails(vnicSetParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, vnicSetDetails)

}

func (cl clientTest) TestAllVnicSets(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allvnicsets),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllVnicSets(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allvnicsets)

}

func (cl clientTest) TestDeleteVnicSet(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteVnicSet(vnicSetParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateVnicSet(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &vnicSetDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.VnicSetParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, vnicSetParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateVnicSet(vnicSetParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, vnicSetDetails)
}

func (cl clientTest) TestUpdateVnicSet(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &vnicSetDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.VnicSetParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, vnicSetParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateVnicSet(vnicSetParams, vnicSetParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, vnicSetDetails)
}
