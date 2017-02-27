// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// ImageListDetails retrieves details of the specified image list.
// You can also use this request to retrieve details of all the available
// image list entries in the specified image list.
func (c Client) ImageListDetails(
	name string,
) (resp response.ImageList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-api: Empty image list name")
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

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

// AllImageList retrieves details of all the available
// image lists in the specified container.
func (c Client) AllImageList() (resp response.AllImageList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

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

// AllImageListNames retrieves the names of objects and
// subcontainers that you can access in the specified container.
func (c Client) AllImageListNames() (resp response.DirectoryNames, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		directory: true,
		client:    &c.http,
		cookie:    c.cookie,
		verb:      "GET",
		url:       url,
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateImageList Adds an image list to Oracle Compute Cloud Service.
func (c Client) CreateImageList(
	def int,
	description string,
	name string,
) (resp response.ImageList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	params := struct {
		Def         int    `json:"default"`
		Description string `json:"description"`
		Name        string `json:"name"`
	}{
		Def:         def,
		Description: description,
		Name:        name,
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		body:   &params,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteImageList deletes an image list
// You can't delete system-provided image lists
// that are available in the /oracle/public container.
func (c Client) DeleteImageList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-api: Empty image list name")
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
	}); err != nil {
		return err
	}

	return nil
}

// UpdateImageList updates the description of an image list.
// You can also update the default image list entry to be used
// while launching instances using the specified image list.
func (c Client) UpdateImageList(
	currentName string,
	newName string,
	description string,
	def int,
) (resp response.ImageList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty curret image list name",
		)
	}

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Def         int    `json:"default"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name"`
	}{
		Def:         def,
		Description: description,
		Name:        newName,
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, newName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "PUT",
		url:    url,
		body:   &params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
