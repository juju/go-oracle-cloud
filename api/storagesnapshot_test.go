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
	storageSnapshotParams = api.StorageSnapshotParams{
		Volume:      "/Compute-acme/jack.jones@example.com/vol1",
		Description: "Remote snapshot of vol1",
		Name:        "/Compute-acme/jack.jones@example.com/vol1/dc14ed57a74c072fd2ec71ed152021eadc10bc868ae17c759f6233ee7317a15d",
	}

	account                = "10737418240"
	storageSnapshotDetails = response.StorageSnapshot{
		Status:           "initializing",
		Start_timestamp:  "2016-07-15T05:03:21Z",
		Name:             "/Compute-acme/jack.jones@example.com/vol1/dc14ed57a74c072fd2ec71ed152021eadc10bc868ae17c759f6233ee7317a15d",
		Description:      "Remote snapshot of vol1",
		Account:          &account,
		Uri:              "https://api-z999.compute.us0.oraclecloud.com/storage/snapshot/Compute-acme/jack.jones@example.com/vol1/dc14ed57a74c072fd2ec71ed152021eadc10bc868ae17c759f6233ee7317a15d",
		Volume:           "/Compute-acme/jack.jones@example.com/vol1",
		Snapshot_id:      "dc14ed57a74c072fd2ec71ed152021eadc10bc868ae17c759f6233ee7317a15d",
		Status_detail:    "The storage snapshot is currently being initialized.",
		Property:         "/oracle/public/storage/snapshot/default",
		Status_timestamp: "2016-07-15T05:03:21Z",
		Size:             10737418240,
	}

	allstoragesnapshots = response.AllStorageSnapshots{
		Result: []response.StorageSnapshot{
			storageSnapshotDetails,
		},
	}

	allstoragesnapshotnames = response.DirectoryNames{
		Result: []string{
			"/Compute-acme/jack.jones@example.com/vol1/d1235ce8155909f0d601cf26c18c95c2786cfb07ece17d8d02676d121eea6485",
			"/Compute-acme/jack.jones@example.com/vol1/snapshot2",
			"/Compute-acme/jack.jones@example.com/vol1/dc14ed57a74c072fd2ec71ed152021eadc10bc868ae17c759f6233ee7317a15d",
		},
	}
)

func (cl clientTest) TestStorageSnapshotWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageSnapshotDetails(storageSnapshotParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStorageSnapshots(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateStorageSnapshot(storageSnapshotParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStorageSnapshotNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteStorageSnapshot(storageSnapshotParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestStorageSnapshotWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateStorageSnapshot(storageSnapshotParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllStorageSnapshots(nil)
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
			err = client.DeleteStorageSnapshot(storageSnapshotParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.StorageSnapshotDetails(storageSnapshotDetails.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllStorageSnapshotNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestStorageSnapshotResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageSnapshotDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage snapshot name"), gc.Equals, true)

	_, err = client.CreateStorageSnapshot(api.StorageSnapshotParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage snapshot name"), gc.Equals, true)
	err = client.DeleteStorageSnapshot("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage snapshot name"), gc.Equals, true)

}

func (cl clientTest) TestStorageSnapshotDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageSnapshotDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.StorageSnapshotDetails(storageSnapshotDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageSnapshotDetails)
}

func (cl clientTest) TestAllStorageSnapshotNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstoragesnapshotnames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStorageSnapshotNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstoragesnapshotnames)
}

func (cl clientTest) TestAllStorageSnapshots(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstoragesnapshots),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStorageSnapshots(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstoragesnapshots)
}

func (cl clientTest) TestDeleteStorageSnapshot(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteStorageSnapshot(storageSnapshotParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateStorageSnapshot(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageSnapshotDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.StorageSnapshotParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, storageSnapshotParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateStorageSnapshot(storageSnapshotParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageSnapshotDetails)
}
