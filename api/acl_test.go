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

func (cl clientTest) TestAclResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	name := client.ComposeName("someName")
	description := "some random description"

	_, err := client.CreateAcl(name, description, false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteAcl(name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllAcls(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AclDetails(name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateAcl(name, name, description, false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestAclResourceWithErrors(c *gc.C) {
	description := "some random description"

	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		name := client.ComposeName("someName")
		_, err := client.CreateAcl(name, description, false, nil)
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
			err = client.DeleteAcl(name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllAcls(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AclDetails(name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateAcl(name, name, description, false, nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestAclResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AclDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty acl name"), gc.Equals, true)

	_, err = client.CreateAcl("", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty acl name"), gc.Equals, true)

	err = client.DeleteAcl("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty acl name"), gc.Equals, true)

	_, err = client.UpdateAcl("", "", "", false, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty acl name"), gc.Equals, true)
}

var (
	aclDetails = response.Acl{
		Name:        "/Compute-acme/jack.jones@example.com/acl1",
		Description: "Sample ACL 1",
		EnableFlag:  false,
		Tags:        nil,
	}

	allacls = response.AllAcls{
		Result: []response.Acl{
			aclDetails,
		},
	}
)

func (cl clientTest) TestCreateAcl(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &aclDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.Acl
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(aclDetails, gc.DeepEquals, req)
		},
	})
	defer ts.Close()

	resp, err := client.CreateAcl(
		aclDetails.Name,
		aclDetails.Description,
		aclDetails.EnableFlag,
		aclDetails.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, aclDetails)
}

func (cl clientTest) TestDeleteAcl(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteAcl(aclDetails.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestAclDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &aclDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AclDetails(aclDetails.Name)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, aclDetails)
}

func (cl clientTest) TestAllAcls(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allacls),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllAcls(nil)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allacls)

}

func (cl clientTest) TestUpdateAcl(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &aclDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.Acl
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(aclDetails, gc.DeepEquals, req)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateAcl(
		aclDetails.Name,
		aclDetails.Name,
		aclDetails.Description,
		aclDetails.EnableFlag,
		aclDetails.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, aclDetails)
}
