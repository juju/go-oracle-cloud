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
	secAssociationParams = struct {
		Name    string
		Seclist string
		Vcable  common.VcableID
	}{
		Name:    "/Compute-acme/jack.jones@example.com/2128a81c-a9e2-49f8-a003-3d43a95197b9",
		Vcable:  common.VcableID("/Compute-acme/jack.jones@example.com/allowed_video_servers"),
		Seclist: "/Compute-acme/jack.jones@example.com/allowed_video_servers",
	}

	secAssociationDetails = response.SecAssociation{
		Name:    "/Compute-acme/jack.jones@example.com/2128a81c-a9e2-49f8-a003-3d43a95197b9",
		Uri:     "https://api-z999.compute.us0.oraclecloud.com/secassociation/Compute-acme/jack.jones@example.com/2128a81c-a9e2-49f8-a003-3d43a95197b9",
		Seclist: "/Compute-acme/jack.jones@example.com/allowed_video_servers",
		Vcable:  "/Compute-acme/jack.jones@example.com/e4d0564b-1e95-464f-92d8-d74c1c583883",
	}

	allsecassociations = response.AllSecAssociations{
		Result: []response.SecAssociation{
			secAssociationDetails,
		},
	}

	secassociationnames = response.DirectoryNames{}
)

func (cl clientTest) TestSecAssociationResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecAssociationDetails(secAssociationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllSecAssociations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecAssociation(secAssociationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecAssociation(secAssociationParams.Name,
		secAssociationParams.Seclist, secAssociationParams.Vcable)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllSecAssociationNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecAssociationResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecAssociation(secAssociationParams.Name,
			secAssociationParams.Seclist, secAssociationParams.Vcable)
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
			err = client.DeleteSecAssociation(secAssociationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecAssociations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecAssociationDetails(secAssociationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllSecAssociationNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecAssociationResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecAssociationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure association name"), gc.Equals, true)

	_, err = client.CreateSecAssociation("", "", "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure association name"), gc.Equals, true)

	err = client.DeleteSecAssociation("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure association name"), gc.Equals, true)

}

func (cl clientTest) TestSecAssociationDetials(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secAssociationDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecAssociationDetails(secAssociationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secAssociationDetails)
}

func (cl clientTest) TestAllSecAssociations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecassociations),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecAssociations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecassociations)
}

func (cl clientTest) TestSecAssociationNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secassociationnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecAssociationNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secassociationnames)
}

func (cl clientTest) TestDeleteSecAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecAssociation(secAssociationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secAssociationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name    string
				Seclist string
				Vcable  common.VcableID
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, secAssociationParams.Name)
			c.Assert(req.Seclist, gc.DeepEquals, secAssociationParams.Seclist)
			c.Assert(req.Vcable, gc.DeepEquals, secAssociationParams.Vcable)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecAssociation(secAssociationParams.Name,
		secAssociationParams.Seclist, secAssociationParams.Vcable)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secAssociationDetails)
}
