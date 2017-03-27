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
	secIpListParams = struct {
		Description  string
		Name         string
		Secipentries []string
	}{
		Secipentries: []string{
			"46.16.56.0/21",
			"46.6.0.0/16",
		},
		Name: "/Compute-acme/jack.jones@example.com/es_iplist",
	}

	secIpListDetails = response.SecIpList{
		Name: "/Compute-acme/jack.jones@example.com/es_iplist",
		Uri:  "https://api-z999.compute.us0.oraclecloud.com/seciplist/Compute-acme/jack.jones@example.com/es_iplist",
		Id:   "e9710ec4-51ac-4248-8d82-427851b634e0",
		Secipentries: []string{
			"46.16.56.0/21",
			"46.6.0.0/16",
		},
		Group_id: "29295",
	}

	allseciplists = response.AllSecIpLists{
		Result: []response.SecIpList{
			secIpListDetails,
		},
	}
)

func (cl clientTest) TestSecIpListResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllDefaultSecIpLists(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllSecIpLists(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecIpList(secIpListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecIpList(
		secIpListParams.Description,
		secIpListParams.Name,
		secIpListParams.Secipentries,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSecIpList(
		secIpListParams.Description,
		secIpListParams.Name,
		secIpListParams.Name,
		secIpListParams.Secipentries,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.SecIpListDetails(secIpListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecIpListResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecIpList(
			secIpListParams.Description,
			secIpListParams.Name,
			secIpListParams.Secipentries,
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
			err = client.DeleteSecIpList(secIpListParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecIpLists(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecIpListDetails(secIpListParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllDefaultSecIpLists(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecIpListResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecIpListDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure ip list name"), gc.Equals, true)

	_, err = client.CreateSecIpList("", "", nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure ip list name"), gc.Equals, true)

	err = client.DeleteSecIpList("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure ip list name"), gc.Equals, true)
}

func (cl clientTest) TestSecIpListDetials(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secIpListDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecIpListDetails(secIpListParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secIpListDetails)
}

func (cl clientTest) TestAllSecIpLists(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allseciplists),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecIpLists(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allseciplists)
}

func (cl clientTest) TestAllDefaultSecIpListss(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allseciplists),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllDefaultSecIpLists(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allseciplists)
}

func (cl clientTest) TestDeleteSecIpList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecIpList(secIpListParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecIpList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secIpListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Description  string
				Name         string
				Secipentries []string
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, secIpListParams.Name)
			c.Assert(req.Description, gc.DeepEquals, secIpListParams.Description)
			c.Assert(req.Secipentries, gc.DeepEquals, secIpListParams.Secipentries)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecIpList(
		secIpListParams.Description,
		secIpListParams.Name,
		secIpListParams.Secipentries,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secIpListDetails)
}

func (cl clientTest) TestUpdateSecIpList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secIpListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Description  string
				Name         string
				Secipentries []string
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, secIpListParams.Name)
			c.Assert(req.Description, gc.DeepEquals, secIpListParams.Description)
			c.Assert(req.Secipentries, gc.DeepEquals, secIpListParams.Secipentries)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateSecIpList(
		secIpListParams.Description,
		secIpListParams.Name,
		secIpListParams.Name,
		secIpListParams.Secipentries,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secIpListDetails)
}
