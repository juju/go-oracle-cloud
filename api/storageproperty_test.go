// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"net/http"
	"strings"

	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
	gc "gopkg.in/check.v1"
)

var (
	storagePropertyDetails = response.StorageProperty{
		Description: "Default storageproperty for all StoragePools and StorageVolumes",
		Name:        "/oracle/public/storage/default",
		Uri:         "https://api-z999.compute.us0.oraclecloud.com/property/storage/oracle/public/storage/default",
	}

	allstorageproperties = response.AllStorageProperties{
		Result: []response.StorageProperty{
			storagePropertyDetails,
		},
	}

	allstoragepropertynames = response.DirectoryNames{
		Result: []string{
			"/oracle/public/",
			"/oracle/private/",
		},
	}
)

func (cl clientTest) TestStoragePropertyWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StoragePropertyDetails(storagePropertyDetails.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStorageProperties(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllStoragePropertyNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestStoragePropertyWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.AllStorageProperties(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.StoragePropertyDetails(storagePropertyDetails.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllStoragePropertyNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestStoragePropertyResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.StoragePropertyDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty storage property name"), gc.Equals, true)
}

func (cl clientTest) TestStoragePropertyDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &storagePropertyDetails),
		check: c,
	})
	defer ts.Close()

	resp, err := client.StoragePropertyDetails(storagePropertyDetails.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, storagePropertyDetails)
}

func (cl clientTest) TestAllStorageProperties(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstorageproperties),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStorageProperties(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstorageproperties)
}

func (cl clientTest) TestAllStoragePropertyNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allstoragepropertynames),
		check: c,
	})
	defer ts.Close()

	resp, err := client.AllStoragePropertyNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allstoragepropertynames)
}
