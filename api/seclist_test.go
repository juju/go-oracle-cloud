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
	secListParams = struct {
		Description          string
		Name                 string
		Outbound_cidr_policy common.SecRuleAction
		Policy               common.SecRuleAction
	}{
		Name:                 "/Compute-acme/jack.jones@example.com/allowed_video_servers",
		Outbound_cidr_policy: common.SecRulePermit,
		Policy:               common.SecRuleDeny,
	}

	secListDetails = response.SecList{
		Account:              "/Compute-acme/default",
		Name:                 "/Compute-acme/jack.jones@example.com/allowed_video_servers",
		Uri:                  "https://api-z999.compute.us0.oraclecloud.com/seclist/Compute-acme/jack.jones@example.com/allowed_video_servers",
		Outbound_cidr_policy: common.SecRulePermit,
		Policy:               common.SecRuleDeny,
	}

	allseclists = response.AllSecLists{
		Result: []response.SecList{
			secListDetails,
		},
	}
)

func (cl clientTest) TestSecListResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.AllSecLists(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSecList(secListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSecList(
		secListParams.Description,
		secListParams.Name,
		secListParams.Outbound_cidr_policy,
		secListParams.Policy,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSecList(
		secListParams.Description,
		secListParams.Name,
		secListParams.Name,
		secListParams.Outbound_cidr_policy,
		secListParams.Policy,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.SecListDetails(secIpListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestSecListResourceWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSecList(
			secListParams.Description,
			secListParams.Name,
			secListParams.Outbound_cidr_policy,
			secListParams.Policy,
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
			err = client.DeleteSecList(secListParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllSecLists(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SecListDetails(secListParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateSecList(
			secListParams.Description,
			secListParams.Name,
			secListParams.Name,
			secListParams.Outbound_cidr_policy,
			secListParams.Policy,
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

func (cl clientTest) TestSecListResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SecListDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure list name"), gc.Equals, true)

	_, err = client.CreateSecList("", "", "", "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure list name"), gc.Equals, true)

	err = client.DeleteSecList("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure list name"), gc.Equals, true)

	_, err = client.UpdateSecList("", "", "", "", "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty secure list name"), gc.Equals, true)

}

func (cl clientTest) TestSecListDetials(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secListDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SecListDetails(secListParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secListDetails)
}

func (cl clientTest) TestAllSecLists(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allseclists),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSecLists(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allseclists)
}

func (cl clientTest) TestDeleteSecList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSecList(secListParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSecList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Description          string
				Name                 string
				Outbound_cidr_policy common.SecRuleAction
				Policy               common.SecRuleAction
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, secListParams.Name)
			c.Assert(req.Description, gc.DeepEquals, secListParams.Description)
			c.Assert(req.Policy, gc.DeepEquals, secListParams.Policy)
			c.Assert(req.Outbound_cidr_policy, gc.DeepEquals,
				secListParams.Outbound_cidr_policy)
		},
	})
	defer ts.Close()

	resp, err := client.CreateSecList(
		secListParams.Description,
		secListParams.Name,
		secListParams.Outbound_cidr_policy,
		secListParams.Policy,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secListDetails)
}

func (cl clientTest) TestUpdateSecList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &secListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Description          string
				Name                 string
				Outbound_cidr_policy common.SecRuleAction
				Policy               common.SecRuleAction
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, secListParams.Name)
			c.Assert(req.Description, gc.DeepEquals, secListParams.Description)
			c.Assert(req.Policy, gc.DeepEquals, secListParams.Policy)
			c.Assert(req.Outbound_cidr_policy, gc.DeepEquals,
				secListParams.Outbound_cidr_policy)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateSecList(
		secListParams.Description,
		secListParams.Name,
		secListParams.Name,
		secListParams.Outbound_cidr_policy,
		secListParams.Policy,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, secListDetails)
}
