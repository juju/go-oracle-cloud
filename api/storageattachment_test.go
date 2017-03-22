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
	storageAttachmentParams = api.StorageAttachmentParams{
		Index:               common.Index(1),
		Instance_name:       "/Compute-acme/jack.jones@example.com/vol1",
		Storage_volume_name: "/Compute-acme/jack.jones@example.com/a6462ba5-5933-41a1-b853-fcfcb421cb07/5fd18f4a-2ac2-4548-a0cf-57774c024742",
	}

	storageAttachmentDetails = response.StorageAttachment{
		Index:               storageAttachmentParams.Index,
		Account:             nil,
		Storage_volume_name: "/Compute-acme/jack.jones@example.com/vol1",
		Hypervisor:          nil,
		Uri:                 "https://api-z999.compute.us0.oraclecloud.com/storage/attachment/Compute-acme/jack.jones@example.com/a6462ba5-5933-41a1-b853-fcfcb421cb07/5fd18f4a-2ac2-4548-a0cf-57774c024742/a7fb4550-df19-497c-a19f-44fc176e1fc2",
		Instance_name:       "/Compute-acme/jack.jones@example.com/a6462ba5-5933-41a1-b853-fcfcb421cb07/5fd18f4a-2ac2-4548-a0cf-57774c024742",
		State:               common.StateAttaching,
		Readonly:            false,
		Name:                "/Compute-acme/jack.jones@example.com/a6462ba5-5933-41a1-b853-fcfcb421cb07/5fd18f4a-2ac2-4548-a0cf-57774c024742/a7fb4550-df19-497c-a19f-44fc176e1fc2",
	}

	allstorageattachments = response.AllStorageAttachments{
		Result: []response.StorageAttachment{
			storageAttachmentDetails,
		},
	}
)

func (cl clientTest) TestStorageAttachmentWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageAttachmentDetails(storageAttachmentDetails.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStorageAttachments(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.CreateStorageAttachment(storageAttachmentParams)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteStorageAttachment(storageAttachmentDetails.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestStorageAttachmentWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateStorageAttachment(storageAttachmentParams)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllStorageAttachments(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.StorageAttachmentDetails(storageAttachmentDetails.Name)
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
			err = client.DeleteStorageAttachment(storageAttachmentDetails.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		ts.Close()
	}
}

func (cl clientTest) TestStorageAttachmentResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StorageAttachmentDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage attachment instance name"), gc.Equals, true)

	_, err = client.CreateStorageAttachment(api.StorageAttachmentParams{})
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage attachment instance name"), gc.Equals, true)

	err = client.DeleteStorageAttachment("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage attachment instance name"), gc.Equals, true)
}

func (cl clientTest) TestStorageAttachmentDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageAttachmentDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.StorageAttachmentDetails(storageAttachmentDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageAttachmentDetails)
}

func (cl clientTest) TestAllStorageAttachments(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstorageattachments),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStorageAttachments(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstorageattachments)
}

func (cl clientTest) TestDeleteStorageAttachment(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteStorageAttachment(storageAttachmentDetails.Name)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestCreateStorageAttachment(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storageAttachmentDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			var req api.StorageAttachmentParams
			err := enc.NewDecoder(r.Body).Decode(&req)

			c.Assert(err, gc.IsNil)
			c.Assert(req, gc.DeepEquals, storageAttachmentParams)
		},
	})
	defer ts.Close()

	resp, err := client.CreateStorageAttachment(storageAttachmentParams)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storageAttachmentDetails)
}
