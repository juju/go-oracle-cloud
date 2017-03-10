package api_test

import (
	"github.com/juju/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

func (cl clientTest) TestBackupConfigratuionWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})

	defer ts.Close()

	interval := map[string]map[string]int{
		"Hourly": map[string]int{
			"hourlyInterval": 1,
		},
	}

	params := api.BackupConfigurationParams{
		Name:                 "/Compute-acme/jack.jones@example.com/backupConfigVol1",
		Description:          "some random description",
		Enabled:              false,
		BackupRetentionCount: 2,
		VolumeUri:            "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
		Interval:             interval,
	}

	_, err := client.CreateBackupConfiguration(params)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateBackupConfiguration(params, params.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteBackupConfiguration(params.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.BackupConfigurationDetails(params.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllBackupConfigurations(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

// TODO(sgiulitti) continue to write unit tests
