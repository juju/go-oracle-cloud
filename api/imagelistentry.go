// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// ImageListEntryDetails retrieves details of the specified image list entry.
func (c Client) ImageListEntryDetails(
	name string,
	version string,
) (resp response.ImageListEntry, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty image list entry name",
		)
	}

	if version == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty image list entry verion",
		)
	}

	url := fmt.Sprintf("%s%s/entry/%s",
		c.endpoints["imagelistentries"], name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteImageListEntry deletes an Image List Entry
func (c Client) DeleteImageListEntry(
	name string,
	version string,
) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New(
			"go-oracle-cloud: Cannot retrive entry from empty image list name",
		)
	}

	if version == "" {
		return errors.New(
			"go-oracle-cloud: Empty image list entry verion",
		)
	}

	url := fmt.Sprintf("%s%s/entry/%s",
		c.endpoints["imagelistentries"], name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "DELETE",
		url:    url,
	}); err != nil {
		return err
	}

	return nil
}

// AddImageListEntry adds an image list entry to Oracle Compute Cloud
// Each machine image in an image list is identified by an image list entry.
func (c Client) AddImageListEntry(
	name string,
	attributes map[string]interface{},
	version int,
	machineImages []string,
) (resp response.ImageListEntryAdd, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Cannot create entry from empty image list name",
		)
	}

	if attributes == nil {
		return resp, errors.New(
			"go-oracle-cloud: Cannot create entry from nil attributes",
		)
	}

	if machineImages == nil {
		return resp, errors.New(
			"go-oracle-cloud: Cannot create entry from nil machineImages",
		)
	}

	params := struct {
		Attributes    map[string]interface{} `json:"attributes"`
		MachineImages []string               `json:"machineImages"`
		Version       int                    `json:"version"`
	}{
		Attributes:    attributes,
		MachineImages: machineImages,
		Version:       version,
	}

	url := fmt.Sprintf("%s%s/entry/%s",
		c.endpoints["imagelistentries"], name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		resp:   &resp,
		body:   &params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
