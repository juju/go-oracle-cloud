package api_test

import (
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
	gc "gopkg.in/check.v1"
)

var (
	instanceParams = api.Instances{
		Shape:     "oc3",
		Imagelist: "/oracle/public/oel_6.4_2GB_v1",
		Name:      "/Compute-acme/jack.jones@example.com/dev-vm",
		Label:     "dev-vm",
		SSHKeys: []string{
			"/Compute-acme/jack.jones@example.com/dev-key1",
		},
	}

	createInstanceParams = api.InstanceParams{
		Instances: []api.Instances{instanceParams},
	}

	instanceDetails = response.Instance{
		Name:      instanceParams.Name,
		Imagelist: instanceParams.Imagelist,
		Label:     instanceParams.Label,
		SSHKeys:   instanceParams.SSHKeys,
		Shape:     instanceParams.Shape,
	}

	allinstances = response.AllInstances{
		Result: []response.Instance{
			instanceDetails,
		},
	}
)

func (cl clientTest) TestInstanceResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.InstanceDetails(instanceParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllInstances(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteInstance(instanceParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllInstanceNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateInstance(createInstanceParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestInstanceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.InstanceDetails(instanceParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllInstances(nil)
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
			err = client.DeleteInstance(instanceParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllInstanceNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateInstance(createInstanceParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestInstanceResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.InstanceDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty instance name"), gc.Equals, true)

	err = client.DeleteInstance("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty instance name"), gc.Equals, true)

	_, err = client.CreateInstance(api.InstanceParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty instances in instance params"), gc.Equals, true)

	_, err = client.CreateInstance(api.InstanceParams{
		Instances: []api.Instances{
			api.Instances{},
		},
	})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty instance name"), gc.Equals, true)
	//TODO(sgiulitti) more test cases
}

func (cl clientTest) TestInstnaceDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &instanceDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.InstanceDetails(instanceParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, instanceDetails)
}

func (cl clientTest) TestDeleteInstance(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteInstance(instanceParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestAllInstances(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allinstances),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllInstances(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allinstances)
}

// TODO(sgiulitti) test here the launchplan also
