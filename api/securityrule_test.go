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
	securityRuleParams = api.SecurityRuleParams{
		Name:                   "/Compute-acme/jack.jones@example.com/secrule1",
		FlowDirection:          "egress",
		Description:            "Sample security rule",
		Acl:                    "/Compute-acme/jack.jones@example.com/acl1",
		SrcVnicSet:             "/Compute-acme/jack.jones@example.com/vnicset1",
		DstVnicSet:             "/Compute-acme/jack.jones@example.com/vnicset2",
		SecProtocols:           []string{"/Compute-acme/jack.jones@example.com/secprotocol1"},
		SrcIpAddressPrefixSets: []string{"/Compute-acme/jack.jones@example.com/ipaddressprefixset1"},
	}

	securityRuleDetails = response.SecurityRule{
		Name:                   securityRuleParams.Name,
		FlowDirection:          securityRuleParams.FlowDirection,
		Description:            securityRuleParams.Description,
		Acl:                    securityRuleParams.SrcVnicSet,
		SrcVnicSet:             securityRuleParams.SrcVnicSet,
		DstVnicSet:             securityRuleParams.DstVnicSet,
		SecProtocols:           securityRuleParams.SecProtocols,
		SrcIpAddressPrefixSets: securityRuleParams.SrcIpAddressPrefixSets,
		Uri:                    "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/secrule/Compute-acme/jack.jones@example.com/secrule1",
		EnabledFlag:            true,
		DstIpAddressPrefixSets: nil,
	}

	allsecurityRules = response.AllSecurityRules{
		Result: []response.SecurityRule{
			securityRuleDetails,
		},
	}
)

func (cl clientTest) TestSecurityRuleResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllSecurityRules(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecurityRule(securityRuleParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecurityRule(securityRuleParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSecurityRule(securityRuleParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.SecurityRuleDetails(securityRuleParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecurityRuleResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecurityRule(securityRuleParams)
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
			err = client.DeleteSecurityRule(securityRuleParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecurityRules(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecurityRuleDetails(securityRuleParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateSecurityRule(securityRuleParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecurityRuleResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecurityRuleDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security rule name"), gc.Equals, true)

	_, err = client.CreateSecurityRule(api.SecurityRuleParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security rule name"), gc.Equals, true)

	err = client.DeleteSecurityRule("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security rule name"), gc.Equals, true)

	_, err = client.UpdateSecurityRule(api.SecurityRuleParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty security rule name"), gc.Equals, true)

}

func (cl clientTest) TestSecurityRuleDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityRuleDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecurityRuleDetails(securityRuleParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityRuleDetails)
}

func (cl clientTest) TestAllSecurityRules(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecurityRules),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecurityRules(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecurityRules)
}

func (cl clientTest) TestDeleteSecurityRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecurityRule(securityRuleParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecurityRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityRuleDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecurityRuleParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, securityRuleParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecurityRule(securityRuleParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityRuleDetails)
}

func (cl clientTest) TestUpdateSecurityRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &securityRuleDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecurityRuleParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, securityRuleParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateSecurityRule(securityRuleParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, securityRuleDetails)
}
