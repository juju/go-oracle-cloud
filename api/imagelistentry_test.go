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

func (cl clientTest) TestImageListEntryResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.CreateImageListEntry(
		imageListEntryParams.name,
		nil,
		imageListEntryParams.version,
		imageListEntryParams.machineimages,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteImageListEntry(
		imageListEntryParams.name,
		imageListEntryParams.version,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.ImageListEntryDetails(
		imageListEntryParams.name,
		imageListEntryParams.version,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

}

var (
	imageListEntryParams = struct {
		name          string
		attributes    map[string]interface{}
		machineimages []string
		version       int
	}{
		name:          "/Compute-some/somename@gmail.com/somename",
		version:       2,
		machineimages: []string{"/oracle/public/oel_6.4_2GB_v2"},
	}

	imageListEntryDetails = response.ImageListEntry{
		Attributes:    response.AttributesEntry{},
		Imagelist:     "/Compute-some/somename@gmail.com/somename",
		Version:       2,
		Machineimages: []string{"/oracle/public/oel_6.4_2GB_v2"},
		Uri:           "https://api-z999.compute.us0.oraclecloud.com/imagelist/Compute-acme/jack.jones@example.com/prodimages/entry/2",
	}

	entryDetials = response.ImageListEntryAdd{
		Imagelist: response.ImageList{
			Default:     1,
			Description: nil,
			Entries:     nil,
			Uri:         "imagelist/Compute-acme/jack.jones@example.com/ol6",
			Name:        "/Compute-acme/jack.jones@example.com/ol6",
		},
		Version:       2,
		Machineimages: []string{"/oracle/public/oel_6.4_2GB_v2"},
		Uri:           "https://api-z999.compute.us0.oraclecloud.com/imagelist/Compute-acme/jack.jones@example.com/prodimages/entry/2",
	}
)

func (cl clientTest) TestImageListEntryWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateImageListEntry(
			imageListEntryParams.name,
			nil,
			imageListEntryParams.version,
			imageListEntryParams.machineimages,
		)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		// when we encounter this case that,
		// the delete method is recivng http.StatusNotFound
		// this means for the delete resource point of view to not be
		// an acutal error and it will return nil so we don't need to check this
		if key != http.StatusNotFound {
			err = client.DeleteImageListEntry(imageListEntryParams.name,
				imageListEntryParams.version)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.ImageListEntryDetails(imageListEntryParams.name,
			imageListEntryParams.version)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		ts.Close()
	}
}

func (cl clientTest) TestImageListEntryResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.ImageListEntryDetails("", 0)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry name"), gc.Equals, true)

	_, err = client.ImageListEntryDetails(imageListEntryParams.name, 0)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry version"), gc.Equals, true)

	_, err = client.CreateImageListEntry("", nil, 0, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry name"), gc.Equals, true)

	_, err = client.CreateImageListEntry(
		imageListEntryParams.name,
		nil, 0, nil)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry machine images"), gc.Equals, true)

	_, err = client.CreateImageListEntry(
		imageListEntryParams.name,
		nil, 0, imageListEntryParams.machineimages)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry verion"), gc.Equals, true)

	err = client.DeleteImageListEntry("", 0)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry name"), gc.Equals, true)

	err = client.DeleteImageListEntry(imageListEntryParams.name, 0)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list entry verion"), gc.Equals, true)

}

func (cl clientTest) TestCreateImageListEntry(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &entryDetials),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Attributes    map[string]interface{} `json:"attributes,omitempty"`
				MachineImages []string               `json:"machineImages"`
				Version       int                    `json:"version"`
			}{}

			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req.Version, gc.DeepEquals, imageListEntryParams.version)
			c.Assert(req.Attributes, gc.DeepEquals, imageListEntryParams.attributes)
			c.Assert(req.MachineImages, gc.DeepEquals, imageListEntryParams.machineimages)
		},
		u: &unmarshaler{
			raw:  imageListEntryRaw,
			into: &response.ImageListEntry{},
		},
	})
	defer ts.Close()

	resp, err := client.CreateImageListEntry(
		imageListEntryParams.name,
		imageListEntryParams.attributes,
		imageListEntryParams.version,
		imageListEntryParams.machineimages,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, entryDetials)
}

func (cl clientTest) TestDeleteImageListEntry(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteImageListEntry(
		imageListEntryParams.name,
		imageListEntryParams.version,
	)
	c.Assert(err, gc.IsNil)
}

func (cl clientTest) TestImageListEntryDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
		body:  createResponse(c, &imageListEntryDetails),
		u: &unmarshaler{
			raw:  imageListEntryRaw,
			into: &response.ImageListEntry{},
		},
	})
	defer ts.Close()

	resp, err := client.ImageListEntryDetails(
		imageListEntryParams.name,
		imageListEntryParams.version,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, imageListEntryDetails)
}
