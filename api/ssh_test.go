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
	sshParams = struct {
		Name    string
		Key     string
		Enabled bool
	}{
		Key:     "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDzU21CEj6JsqIMQAYwNbmZ5P2BVxA...",
		Name:    "/Compute-acme/jack.jones@example.com/key1",
		Enabled: true,
	}

	sshKeyDetails = response.SSH{
		Name:    sshParams.Name,
		Uri:     "https://api-z999.compute.us0.oraclecloud.com/sshkey/Compute-acme/jack.jones@example.com/key1",
		Enabled: sshParams.Enabled,
	}

	allsshkeys = response.AllSSH{
		Result: []response.SSH{
			sshKeyDetails,
		},
	}

	sshKeyNames = response.AllSSHNames{
		Result: []string{
			"/Compute-acme/jack.jones@example.com/adminkey",
			"/Compute-acme/jack.jones@example.com/permkey",
			"/Compute-acme/jack.jones@example.com/tempkey",
		},
	}
)

func (cl clientTest) TestSSHResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SSHKeyDetails(sshParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllSSHKeys(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateSHHKey(sshParams.Name, sshParams.Key, sshParams.Enabled)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateSSHKey(
		sshParams.Name,
		sshParams.Name,
		sshParams.Key,
		sshParams.Enabled,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteSSHKey(sshParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestSSHWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateSHHKey(sshParams.Name, sshParams.Key, sshParams.Enabled)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllSSHKeys(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.SSHKeyDetails(sshParams.Name)
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
			err = client.DeleteSSHKey(sshParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.UpdateSSHKey(
			sshParams.Name,
			sshParams.Name,
			sshParams.Key,
			sshParams.Enabled,
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

func (cl clientTest) TestSSHResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.SSHKeyDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ssh key name"), gc.Equals, true)

	_, err = client.CreateSHHKey("", "", false)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ssh key name"), gc.Equals, true)

	err = client.DeleteSSHKey("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ssh key name"), gc.Equals, true)

	_, err = client.UpdateSSHKey("", "", "", false)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty ssh key name"), gc.Equals, true)

}

func (cl clientTest) TestSSHDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &sshKeyDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.SSHKeyDetails(sshParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, sshKeyDetails)
}

func (cl clientTest) TestAllSSHKeys(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allsshkeys),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSSHKeys(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allsshkeys)
}

func (cl clientTest) TestSSHKeysNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &sshKeyNames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllSSHKeyNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, sshKeyNames)
}

func (cl clientTest) TestDeleteSSHKey(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteSSHKey(sshParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateSSHKey(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &sshKeyDetails),
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name    string
				Key     string
				Enabled bool
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, sshParams.Name)
			c.Assert(req.Key, gc.DeepEquals, sshParams.Key)
			c.Assert(req.Enabled, gc.DeepEquals, sshParams.Enabled)

		},
	})
	defer ts.Close()

	resp, err := client.CreateSHHKey(sshParams.Name, sshParams.Key, sshParams.Enabled)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, sshKeyDetails)

}

func (cl clientTest) TestUpdateSSHKey(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &sshKeyDetails),
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Name    string
				Key     string
				Enabled bool
			}{}
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req.Name, gc.DeepEquals, sshParams.Name)
			c.Assert(req.Key, gc.DeepEquals, sshParams.Key)
			c.Assert(req.Enabled, gc.DeepEquals, sshParams.Enabled)

		},
	})
	defer ts.Close()

	resp, err := client.UpdateSSHKey(sshParams.Name, sshParams.Name, sshParams.Key, sshParams.Enabled)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, sshKeyDetails)

}
