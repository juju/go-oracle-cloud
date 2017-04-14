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
	imageListParams = struct {
		Def         int
		Description string
		Name        string
	}{
		Def:         1,
		Description: "ol 6",
		Name:        "/Compute-acme/jack.jones@example.com/ol6",
	}

	imageListDetails = response.ImageList{
		Default:     1,
		Description: &imageListParams.Description,
		Uri:         "https://api-z999.compute.us0.oraclecloud.com/imagelist/Compute-acme/admin/ol6",
		Name:        "/Compute-acme/jack.jones@example.com/ol6",
	}

	allimagelists = response.AllImageLists{
		Result: []response.ImageList{
			imageListDetails,
		},
	}

	imagelistnames = response.DirectoryNames{
		Result: []string{
			"/Compute-acme/jack.jones@example.com",
			"/Compute-acme/jack.someoher@example.com",
		},
	}
)

func (cl clientTest) TestImageListResourceWithNoAuthentication(c *gc.C) {
	ts, client := cl.StartTestServer(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.CreateImageList(
		imageListParams.Def,
		imageListParams.Description,
		imageListParams.Name,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	err = client.DeleteImageList(imageListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.UpdateImageList(
		imageListParams.Name,
		imageListParams.Name,
		imageListParams.Description,
		imageListParams.Def,
	)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.ImageListDetails(imageListParams.Name)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.AllImageLists(nil)
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)

	_, err = client.ImageListNames()
	c.Assert(err, gc.NotNil)
	c.Assert(api.IsNotAuth(err), gc.Equals, true)
}

func (cl clientTest) TestImageListWithErrors(c *gc.C) {
	for key, val := range httpStatusErrors {
		ts, client := cl.StartTestServerAuth(httpParams{
			manualHeaderStatus: true,
			check:              c,
			body:               createResponse(c, errAPI),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(key)
			},
		})

		_, err := client.CreateImageList(
			imageListParams.Def,
			imageListParams.Description,
			imageListParams.Name,
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
			err = client.DeleteImageList(imageListParams.Name)
			c.Assert(err, gc.NotNil)
			c.Assert(val(err), gc.Equals, true)
			c.Assert(
				strings.Contains(err.Error(), errAPI.Message),
				gc.Equals,
				true)
		}

		_, err = client.ImageListDetails(imageListParams.Name)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.AllImageLists(nil)
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.ImageListNames()
		c.Assert(err, gc.NotNil)
		c.Assert(val(err), gc.Equals, true)
		c.Assert(
			strings.Contains(err.Error(), errAPI.Message),
			gc.Equals,
			true)

		_, err = client.UpdateImageList(
			imageListParams.Name,
			imageListParams.Name,
			imageListParams.Description,
			imageListParams.Def,
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

func (cl clientTest) TestImageListResourceWithEmptyName(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	_, err := client.ImageListDetails("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list name"), gc.Equals, true)

	_, err = client.CreateImageList(0, "", "")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list name"), gc.Equals, true)

	err = client.DeleteImageList("")
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list name"), gc.Equals, true)

	_, err = client.UpdateImageList("", "", "", 0)
	c.Assert(err, gc.NotNil)
	c.Assert(strings.Contains(err.Error(), "Empty image list current name"), gc.Equals, true)
}

func (cl clientTest) TestCreateImageList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &imageListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Def         int    `json:"default"`
				Description string `json:"description,omitempty"`
				Name        string `json:"name"`
			}{}

			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req.Def, gc.DeepEquals, imageListParams.Def)
			c.Assert(req.Description, gc.DeepEquals, imageListParams.Description)
			c.Assert(req.Name, gc.DeepEquals, imageListParams.Name)
		},
		u: &unmarshaler{
			raw:  imageListDetailsRaw,
			into: &response.ImageList{},
		},
	})
	defer ts.Close()

	resp, err := client.CreateImageList(
		imageListParams.Def,
		imageListParams.Description,
		imageListParams.Name,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, imageListDetails)
}

func (cl clientTest) TestUpdateImageList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &imageListDetails),
		check: c,
		handler: func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Def         int    `json:"default"`
				Description string `json:"description,omitempty"`
				Name        string `json:"name"`
			}{}

			err := enc.NewDecoder(r.Body).Decode(&req)
			c.Assert(err, gc.IsNil)
			c.Assert(req.Def, gc.DeepEquals, imageListParams.Def)
			c.Assert(req.Description, gc.DeepEquals, imageListParams.Description)
			c.Assert(req.Name, gc.DeepEquals, imageListParams.Name)
		},
		u: &unmarshaler{
			raw:  imageListDetailsRaw,
			into: &response.ImageList{},
		},
	})
	defer ts.Close()

	resp, err := client.UpdateImageList(
		imageListParams.Name,
		imageListParams.Name,
		imageListParams.Description,
		imageListParams.Def,
	)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, imageListDetails)
}

func (cl clientTest) TestImageListDetails(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &imageListDetails),
		check: c,
		u: &unmarshaler{
			raw:  imageListDetailsRaw,
			into: &response.ImageList{},
		},
	})
	defer ts.Close()

	resp, err := client.ImageListDetails(imageListParams.Name)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, imageListDetails)

}

func (cl clientTest) TestAllImageLists(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &allimagelists),
		check: c,
		u: &unmarshaler{
			raw:  allImageListDetailsRaw,
			into: &response.AllImageLists{},
		},
	})
	defer ts.Close()

	resp, err := client.AllImageLists(nil)
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, allimagelists)
}

func (cl clientTest) TestImageListNames(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		body:  createResponse(c, &imagelistnames),
		check: c,
		u: &unmarshaler{
			raw:  imageListNamesRaw,
			into: &response.DirectoryNames{},
		},
	})
	defer ts.Close()

	resp, err := client.ImageListNames()
	c.Assert(err, gc.IsNil)
	c.Assert(resp, gc.DeepEquals, imagelistnames)
}

func (cl clientTest) TestDeleteImageList(c *gc.C) {
	ts, client := cl.StartTestServerAuth(httpParams{
		check: c,
	})
	defer ts.Close()

	err := client.DeleteImageList(imageListParams.Name)
	c.Assert(err, gc.IsNil)
}
