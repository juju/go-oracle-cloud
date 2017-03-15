// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"net/http"
	"strings"

	enc "encoding/json"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/common"
	"github.com/juju/go-oracle-cloud/response"
	gc "gopkg.in/check.v1"
)

var (
	backupConfigurationParams = api.BackupConfigurationParams{
		Name:                 "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		Enabled:              false,
		BackupRetentionCount: 2,
		VolumeUri:            "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
		Interval:             common.NewInterval(2),
	}
)

func (cl clientTest) TestBackupConfigratuionWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})

	defer ts.Close()

	_, err := client.CreateBackupConfiguration(backupConfigurationParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateBackupConfiguration(backupConfigurationParams, backupConfigurationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteBackupConfiguration(backupConfigurationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.BackupConfigurationDetails(backupConfigurationParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllBackupConfigurations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestBackupConfigratuionWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateBackupConfiguration(backupConfigurationParams)
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
			err = client.DeleteBackupConfiguration(backupConfigurationParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllBackupConfigurations(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.BackupConfigurationDetails(backupConfigurationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateBackupConfiguration(backupConfigurationParams, backupConfigurationParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}

}

func (cl clientTest) TestBackupConfigurationResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.BackupConfigurationDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup configuration name"), gc.Equals, true)

	_, err = client.CreateBackupConfiguration(api.BackupConfigurationParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup configuration name"), gc.Equals, true)

	err = client.DeleteBackupConfiguration("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup configuration name"), gc.Equals, true)

	_, err = client.UpdateBackupConfiguration(backupConfigurationParams, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup configuration current name"), gc.Equals, true)
}

var (
	backupConfigurationDetails = response.BackupConfiguration{
		Uri:                  "https://api-z999.compute.us0.oraclecloud.com/backupservice/v1/configuration/Compute-acme/jack.jones@example.com/backupConfigVol1",
		RunAsUser:            "/Compute-acme/jack.jones@example.com",
		Name:                 "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		Enabled:              false,
		BackupRetentionCount: 2,
		NextScheduledRun:     "2016-08-19T05:10:44.859Z",
		Interval:             common.NewInterval(2),
	}

	allbackupconfigurations = []response.BackupConfiguration{
		backupConfigurationDetails,
	}
)

func (cl clientTest) TestCreateBackupConfiguration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &backupConfigurationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.BackupConfigurationParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, backupConfigurationParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateBackupConfiguration(backupConfigurationParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, backupConfigurationDetails)
}

func (cl clientTest) TestDeleteBackupConfiguration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteBackupConfiguration(backupConfigurationParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestBackupConfigurationDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &backupConfigurationDetails),
	})
	defer ts.Close()

	resp, err := client.BackupConfigurationDetails(backupConfigurationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, backupConfigurationDetails)
}

func (cl clientTest) TestAllBackupConfigurations(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &allbackupconfigurations),
	})
	defer ts.Close()

	resp, err := client.AllBackupConfigurations(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allbackupconfigurations)
}

func (cl clientTest) TestUpdateBackupConfiguration(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &backupConfigurationDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.BackupConfigurationParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, backupConfigurationParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateBackupConfiguration(backupConfigurationParams, backupConfigurationParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, backupConfigurationDetails)
}
