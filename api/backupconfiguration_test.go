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
	interval = map[string]interface{}{
		"Hourly": map[string]interface{}{
			"hourlyInterval": 2,
		},
	}

	backupParams = api.BackupConfigurationParams{
		Name:                 "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		Enabled:              false,
		BackupRetentionCount: 2,
		VolumeUri:            "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
		Interval:             interval,
	}
)

func (cl clientTest) TestBackupConfigratuionWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})

	defer ts.Close()

	_, err := client.CreateBackupConfiguration(backupParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateBackupConfiguration(backupParams, backupParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteBackupConfiguration(backupParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.BackupConfigurationDetails(backupParams.Name)
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

		_, err := client.CreateBackupConfiguration(backupParams)
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
			err = client.DeleteBackupConfiguration(backupParams.Name)
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

		_, err = client.BackupConfigurationDetails(backupParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateBackupConfiguration(backupParams, backupParams.Name)
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

	_, err = client.UpdateBackupConfiguration(backupParams, "")
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
		Interval:             common.NewInterval(1),
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
			c.Assert(req.Name, gc.DeepEquals, backupParams.Name)
			c.Assert(req.Enabled, gc.DeepEquals, backupParams.Enabled)
			c.Assert(req.BackupRetentionCount,
				gc.DeepEquals, backupParams.BackupRetentionCount)
			c.Assert(req.VolumeUri, gc.DeepEquals, backupParams.VolumeUri)
			//TODO
		},
	})
	defer ts.Close()

	resp, err := client.CreateBackupConfiguration(backupParams)
	c.Assert(err, gc.IsNil)
	//TODO(sgiulitti) fix this
	_ = resp
}