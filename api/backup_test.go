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

func (cl clientTest) TestBackupWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})

	defer ts.Close()

	_, err := client.CreateBackup(backupParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteBackup(backupParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.BackupDetails(backupParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllBackups(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestBackupWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateBackup(backupParams)
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
			err = client.DeleteBackup(backupParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.AllBackups(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.BackupDetails(backupParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}

}

func (cl clientTest) TestBackupResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.BackupDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup name"), gc.Equals, true)

	_, err = client.CreateBackup(api.BackupParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup name"), gc.Equals, true)

	err = client.DeleteBackup("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty backup name"), gc.Equals, true)

}

var (
	backupParams = api.BackupParams{
		BackupConfigurationName: "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		Name:        "/Compute-acme/jack.jones@example.com/BACKUP-A",
		Description: "",
	}

	backupDetails = response.Backup{
		Uri:       "https://api-z999.compute.us0.oraclecloud.com/backupservice/v1/backup/Compute-acme/jack.jones@example.com/BACKUP-A",
		RunAsUser: "/Compute-acme/jack.jones@example.com",
		Name:      "/Compute-acme/jack.jones@example.com/BACKUP-A",
		BackupConfigurationName: "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		VolumeUri:               "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
		ErrorMessage:            nil,
		DetailedErrorMessage:    nil,
		State:                   common.Submitted,
		Description:             nil,
		Bootable:                false,
		Shared:                  false,
		SnapshotUri:             nil,
		SnapshotSize:            nil,
		TagID:                   "0d22fec6-fc3c-4987-8021-8f0cf49b8737",
	}

	allbackups = response.AllBackups{
		Result: []response.Backup{backupDetails},
	}
)

func (cl clientTest) TestCreateBackup(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &backupDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.BackupParams
			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, backupParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateBackup(backupParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, backupDetails)
}

func (cl clientTest) TestDeleteBackup(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteBackup(backupParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestBackupDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &backupDetails),
	})
	defer ts.Close()

	resp, err := client.BackupDetails(backupParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, backupDetails)
}

func (cl clientTest) TestAllBackups(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &allbackups),
	})
	defer ts.Close()

	resp, err := client.AllBackups(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allbackups)
}
