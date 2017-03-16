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
	vcable common.VcableID = "/Compute-acme/jack.jones@example.com/0d363a97-2353-4294-903a-696d290bdac8"

	ipAssociationParams = struct {
		Parentpool common.IPPool
		Vcable     common.VcableID
		Name       string
	}{
		Parentpool: common.NewIPPool(common.PublicIPPool, common.IPPoolType),
		Vcable:     vcable,
		Name:       "/Compute-acme/jack.jones@example.com/1f36bc5a-c104-44ce-9c7e-1742a260f6b5",
	}

	ipAssociationDetails = response.IpAssociation{
		Parentpool: ipAssociationParams.Parentpool,
		Vcable:     ipAssociationParams.Vcable,
	}

	allipassociations = response.AllIpAssociations{
		Result: []response.IpAssociation{
			ipAssociationDetails,
		},
	}
)

func (cl clientTest) TestIpAssociationWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAssociationDetails(ipAssociationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpAssociations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpAssociation(
		ipAssociationParams.Parentpool,
		ipAssociationParams.Vcable,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpAssociation(ipAssociationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestIpAssociationWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpAssociationDetails(ipAssociationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpAssociations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpAssociation(
			ipAssociationParams.Parentpool,
			ipAssociationParams.Vcable,
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
			err = client.DeleteIpAssociation(ipAssociationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		ts.Close()
	}
}

func (cl clientTest) TestIpAssociationWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAssociationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip association name"), gc.Equals, true)

	_, err = client.CreateIpAssociation("", "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty vcable id"), gc.Equals, true)

	err = client.DeleteIpAssociation("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip association name"), gc.Equals, true)
}

func (cl clientTest) TestIpAssociationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAssociationDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.IpAssociationDetails(ipAssociationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAssociationDetails)

}

func (cl clientTest) TestAllIpAssocitations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipassociations),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllIpAssociations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipassociations)

}

func (cl clientTest) TestDeleteIpAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpAssociation(ipAssociationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAssociationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Parentpool common.IPPool   `json:"parentpool"`
				Vcable     common.VcableID `json:"vcable"`
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Parentpool, gc.DeepEquals, ipAssociationDetails.Parentpool)
			c.Assert(req.Vcable, gc.DeepEquals, ipAssociationDetails.Vcable)
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpAssociation(
		ipAssociationDetails.Parentpool,
		ipAssociationDetails.Vcable,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAssociationDetails)
}
