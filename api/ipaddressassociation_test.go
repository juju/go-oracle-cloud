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
	ipAddressAssociationParams = struct {
		name                 string
		ipAddressReservation string
		vnic                 string
		description          string
		tags                 []string
	}{
		name:                 "/Compute-acme/jack.jones@example.com/ipassociation1",
		ipAddressReservation: "/Compute-acme/jack.jones@example.com/ipreservation1",
		vnic:                 "/Compute-acme/jack.jones@example.com/2e6627de-6842-49bc-9c28-21da524c297d/eth0",
	}

	ipAddressAssociationDetails = response.IpAddressAssociation{
		Name:                 ipAddressAssociationParams.name,
		IpAddressReservation: ipAddressAssociationParams.ipAddressReservation,
		Description:          ipAddressAssociationParams.description,
		Tags:                 ipAddressAssociationParams.tags,
		Uri:                  "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipassociation/Compute-acme/jack.jones@example.com/ipassociation1",
		Vnic:                 ipAddressAssociationParams.vnic,
	}

	allipaddressassociations = response.AllIpAddressAssociations{
		Result: []response.IpAddressAssociation{
			ipAddressAssociationDetails,
		},
	}
)

func (cl clientTest) TestIpAddressAssociationWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressAssociationDetails(ipAddressAssociationParams.name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllIpAddressAssociations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateIpAddressAssociation(
		ipAddressAssociationParams.description,
		ipAddressAssociationParams.ipAddressReservation,
		ipAddressAssociationParams.vnic,
		ipAddressAssociationParams.name,
		ipAddressAssociationParams.tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteIpAddressAssociation(ipAddressAssociationParams.name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateIpAddressAssociation(
		ipAddressAssociationParams.name,
		ipAddressAssociationParams.ipAddressReservation,
		ipAddressAssociationParams.vnic,
		ipAddressAssociationParams.name,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestIpAddressAssociationWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.IpAddressAssociationDetails(ipAddressAssociationParams.name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllIpAddressAssociations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.CreateIpAddressAssociation(
			ipAddressAssociationParams.description,
			ipAddressAssociationParams.ipAddressReservation,
			ipAddressAssociationParams.vnic,
			ipAddressAssociationParams.name,
			ipAddressAssociationParams.tags,
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
			err = client.DeleteIpAddressAssociation(ipAddressAssociationParams.name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateIpAddressAssociation(
			ipAddressAssociationParams.name,
			ipAddressAssociationParams.ipAddressReservation,
			ipAddressAssociationParams.vnic,
			ipAddressAssociationParams.name,
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

func (cl clientTest) TestIpAddressAssociationWithEmptyName(c *gc.C) {

	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.IpAddressAssociationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association name"), gc.Equals, true)

	_, err = client.CreateIpAddressAssociation(
		ipAddressAssociationParams.description,
		ipAddressAssociationParams.ipAddressReservation,
		ipAddressAssociationParams.vnic,
		"",
		ipAddressAssociationParams.tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association name"), gc.Equals, true)

	_, err = client.CreateIpAddressAssociation(
		ipAddressAssociationParams.description,
		"",
		ipAddressAssociationParams.vnic,
		ipAddressAssociationParams.name,
		ipAddressAssociationParams.tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association reservation"), gc.Equals, true)

	_, err = client.CreateIpAddressAssociation(
		ipAddressAssociationParams.description,
		ipAddressAssociationParams.ipAddressReservation,
		"",
		ipAddressAssociationParams.name,
		ipAddressAssociationParams.tags,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association vnic"), gc.Equals, true)

	err = client.DeleteIpAddressAssociation("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association name"), gc.Equals, true)

	_, err = client.UpdateIpAddressAssociation(
		"",
		ipAddressAssociationParams.ipAddressReservation,
		ipAddressAssociationParams.vnic,
		"",
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association current name"), gc.Equals, true)

	_, err = client.UpdateIpAddressAssociation(
		ipAddressAssociationParams.name,
		"",
		ipAddressAssociationParams.vnic,
		ipAddressAssociationParams.name,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association reservation"), gc.Equals, true)

	_, err = client.UpdateIpAddressAssociation(
		ipAddressAssociationParams.name,
		ipAddressAssociationParams.ipAddressReservation,
		"",
		ipAddressAssociationParams.name,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ip address association vnic"), gc.Equals, true)
}

func (cl clientTest) TestIpAddressAssociationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressAssociationDetails),
		check: c,
	})

	defer ts.Close()

	resp, err := client.IpAddressAssociationDetails(ipAddressAssociationDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressAssociationDetails)

}

func (cl clientTest) TestAllIpAddressAssocitations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allipaddressassociations),
		check: c,
	})

	defer ts.Close()

	resp, err := client.AllIpAddressAssociations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allipaddressassociations)

}

func (cl clientTest) TestDeleteIpAddressAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})

	defer ts.Close()

	err := client.DeleteIpAddressAssociation(ipAddressAssociationDetails.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateIpAddressAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressAssociationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.IpAddressAssociation
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, ipAddressAssociationDetails.Name)
			c.Assert(req.Description, gc.DeepEquals,
				ipAddressAssociationDetails.Description)
			c.Assert(req.Tags, gc.DeepEquals, ipAddressAssociationDetails.Tags)
			c.Assert(req.Vnic, gc.DeepEquals, ipAddressAssociationDetails.Vnic)
			c.Assert(req.IpAddressReservation, gc.DeepEquals,
				ipAddressAssociationDetails.IpAddressReservation)
			c.Assert(req.Uri, gc.DeepEquals, "")
		},
	})
	defer ts.Close()

	resp, err := client.CreateIpAddressAssociation(
		ipAddressAssociationDetails.Description,
		ipAddressAssociationDetails.IpAddressReservation,
		ipAddressAssociationDetails.Vnic,
		ipAddressAssociationDetails.Name,
		ipAddressAssociationDetails.Tags,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressAssociationDetails)
}

func (cl clientTest) TestUpdateIpAssociation(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &ipAddressAssociationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req response.IpAddressAssociation
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)

			c.Assert(req.Name, gc.DeepEquals, ipAddressAssociationDetails.Name)
			c.Assert(req.Description, gc.DeepEquals,
				ipAddressAssociationDetails.Description)
			c.Assert(req.Tags, gc.DeepEquals, ipAddressAssociationDetails.Tags)
			c.Assert(req.Vnic, gc.DeepEquals, ipAddressAssociationDetails.Vnic)
			c.Assert(req.IpAddressReservation, gc.DeepEquals,
				ipAddressAssociationDetails.IpAddressReservation)
			c.Assert(req.Uri, gc.DeepEquals, "")
		},
	})
	defer ts.Close()

	resp, err := client.UpdateIpAddressAssociation(
		ipAddressAssociationDetails.Name,
		ipAddressAssociationDetails.IpAddressReservation,
		ipAddressAssociationDetails.Vnic,
		ipAddressAssociationDetails.Name,
	)

	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, ipAddressAssociationDetails)
}
