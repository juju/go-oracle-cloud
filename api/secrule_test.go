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
	secRuleParams = api.SecRuleParams{
		Action:      common.SecRulePermit,
		Application: "/Compute-acme/jack.jones@example.com/video_streaming_udp",
		Name:        "/Compute-acme/jack.jones@example.com/es_to_videoservers_stream",
		Dst_list:    "seclist:/Compute-acme/jack.jones@example.com/allowed_video_servers",
		Src_list:    "seciplist:/Compute-acme/jack.jones@example.com/es_iplist",
	}

	secRuleDetails = response.SecRule{
		Uri:         "https://api-z999.compute.us0.oraclecloud.com/secrule/Compute-acme/jack.jones@example.com/es_to_videoservers_stream",
		Application: secRuleParams.Application,
		Name:        secRuleParams.Name,
		Dst_list:    secRuleParams.Dst_list,
		Src_list:    secRuleParams.Src_list,
		Disabled:    secRuleParams.Disabled,
		Action:      secRuleParams.Action,
	}

	allsecrules = response.AllSecRules{
		Result: []response.SecRule{
			secRuleDetails,
		},
	}
)

func (cl clientTest) TestSecRuleResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllSecRules(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecRule(secRuleParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecRule(secRuleParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSecRule(secRuleParams, secRuleParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.SecRuleDetails(secRuleParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecRuleResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecRule(secRuleParams)
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
			err = client.DeleteSecRule(secRuleParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecRules(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecRuleDetails(secRuleParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateSecRule(secRuleParams, secRuleParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestSecRuleResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecRuleDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure rule name"), gc.Equals, true)

	_, err = client.CreateSecRule(api.SecRuleParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure rule name"), gc.Equals, true)

	err = client.DeleteSecRule("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure rule name"), gc.Equals, true)

	_, err = client.UpdateSecRule(api.SecRuleParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure rule name"), gc.Equals, true)

}

func (cl clientTest) TestSecRuleDetials(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secRuleDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecRuleDetails(secRuleParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secRuleDetails)
}

func (cl clientTest) TestAllSecRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsecrules),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecRules(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsecrules)
}

func (cl clientTest) TestDeleteSecRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecRule(secRuleParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secRuleDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecRuleParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, secRuleParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecRule(secRuleParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secRuleDetails)
}

func (cl clientTest) TestUpdateSecRule(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secRuleDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.SecRuleParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, secRuleParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateSecRule(secRuleParams, secRuleParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secRuleDetails)
}
