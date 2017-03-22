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
	storageVolumeParams = api.StorageVolumeParams{
		Name:       "/Compute-acme/jack.jones@example.com/vol1",
		Properties: []common.StoragePool{common.DefaultPool},
		Size:       common.NewStorageSize(10, common.G),
	}

	storageVolumeDetails = response.StorageVolume{
		Status:            "Initializing",
		Account:           "/Compute-acme/default",
		Writecache:        false,
		Managed:           true,
		Description:       nil,
		Tags:              nil,
		Bootable:          true,
		Hypervisor:        nil,
		Quota:             nil,
		Uri:               "https://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol2",
		Snapshot:          nil,
		Status_detail:     "The storage volume is currently being initialized.",
		Imagelist_entry:   1,
		Storage_pool:      "/Compute-acme/storagepool/iscsi/thruput_1",
		Machineimage_name: "/Compute-acme/jack.jones@example.com/linux-oel-6.7-x86_64",
		Status_timestamp:  "2016-04-18T14:24:05Z",
		Shared:            false,
		Imagelist:         "/Compute-acme/jack.jones@example.com/ol6",
		Size:              21474836480,
		Properties:        []common.StoragePool{common.DefaultPool},
		Name:              "/Compute-acme/jack.jones@example.com/vol2",
	}

	allstoragevolumes = response.AllStorageVolumes{
		Result: []response.StorageVolume{
			storageVolumeDetails,
		},
	}
)

func (cl clientTest) TestStorageVolumeWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageVolumeDetails(storageVolumeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStorageVolumes(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateStorageVolume(storageVolumeParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateStorageVolume(storageVolumeParams, storageVolumeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteStorageVolume(storageVolumeParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

func (cl clientTest) TestStorageVolumeWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateStorageVolume(storageVolumeParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllStorageVolumes(nil)
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
			err = client.DeleteStorageVolume(storageVolumeParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.StorageVolumeDetails(storageVolumeParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateStorageVolume(storageVolumeParams, storageVolumeParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestStorageVolumeResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageVolumeDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage volume name"), gc.Equals, true)

	_, err = client.CreateStorageVolume(api.StorageVolumeParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage volume name"), gc.Equals, true)

	err = client.DeleteStorageVolume("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage volume name"), gc.Equals, true)

	_, err = client.UpdateStorageVolume(api.StorageVolumeParams{}, "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage volume name"), gc.Equals, true)

}

func (cl clientTest) TestStorageVolumeDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageVolumeDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.StorageVolumeDetails(storageVolumeDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageVolumeDetails)
}

func (cl clientTest) TestAllStorageVolumes(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstoragevolumes),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStorageVolumes(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstoragevolumes)
}

func (cl clientTest) TestDeleteStorageVolume(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteStorageVolume(storageVolumeParams.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateStorageVolume(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageVolumeDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.StorageVolumeParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, storageVolumeParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateStorageVolume(storageVolumeParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageVolumeDetails)
}

func (cl clientTest) TestUpdateStorageVolume(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageVolumeDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.StorageVolumeParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, storageVolumeParams)
		},
	})
	defer ts.Close()

	resp, err := client.UpdateStorageVolume(storageVolumeParams, storageVolumeParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageVolumeDetails)
}
