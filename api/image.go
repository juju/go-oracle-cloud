package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// ImageListDetails retrieves details of the specified image list.
// You can also use this request to retrieve details of all the available
// image list entries in the specified image list.
func (c Client) ImageListDetails(name string) (resp response.ImageList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-api: Empty image name")
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
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
		treat:  defaultTreat,
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
		treat:     defaultTreat,
		resp:      &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateImageList Adds an image list to Oracle Compute Cloud Service.
func (c Client) CreateImageList(def uint64, description string, name string) (resp response.ImageList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	image := struct {
		Def         uint64 `json:"default"`
		Description string `json:"description"`
		Name        string `json:"name"`
	}{
		Def:         def,
		Description: description,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		body:   &image,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteImageList deletes an image list
func (c Client) DeleteImageList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-api: Empty image name")
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat:  defaultDeleteTreat,
	}); err != nil {
		return err
	}

	return nil
}

// UpdateImageList updates the description of an image list.
// You can also update the default image list entry to be used
// while launching instances using the specified image list.
func (c Client) UpdateImageList(target, name, description string, def uint64) (err error) {

	if !c.isAuth() {
		return ErrNotAuth
	}

	if target == "" {
		return errors.New("go-oracle-cloud: Cannot update image list because of empty target")
	}

	image := struct {
		Def         uint64 `json:"default"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name"`
	}{
		Def:         def,
		Description: description,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, target)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "PUT",
		url:    url,
		body:   &image,
		treat:  defaultTreat,
	}); err != nil {
		return err
	}

	return nil
}

// ImageListEntry Retrieves details of the specified image list entry.
func (c Client) ImageListEntry(name string, version uint64) (resp response.Entries, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Cannot retrive entry from empty image list name",
		)
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s/entry/%d",
		c.endpoint, c.identify, c.username, name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteImageListEntry deletes an Image List Entry
func (c Client) DeleteImageListEntry(name string, version uint64) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New(
			"go-oracle-cloud: Cannot retrive entry from empty image list name",
		)
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s/entry/%d",
		c.endpoint, c.identify, c.username, name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "DELETE",
		url:    url,
		treat:  defaultDeleteTreat,
	}); err != nil {
		return err
	}

	return nil
}

// AddImageListEntry adds an Image List Entry
func (c Client) AddImageListEntry(
	name string,
	version uint64,
	attributes map[string]interface{},
	machineImages []string,
) (resp response.Entries, err error) {

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

	n := len(machineImages)
	// suppose we have only machine images names
	// so we must make them oracle cloud api complaint
	// when we are passing them into the post body
	for i := 0; i < n; i++ {
		machineImages[i] = fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, machineImages[i])
	}

	entry := struct {
		Attributes    map[string]interface{} `json:"attributes"`
		MachineImages []string               `json:"machineImages"`
		Version       uint64                 `json:"version"`
	}{
		Attributes:    attributes,
		MachineImages: machineImages,
		Version:       version,
	}

	url := fmt.Sprintf("%s/imagelist/Compute-%s/%s/%s/entry/%d",
		c.endpoint, c.identify, c.username, name, version)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		treat:  defaultPostTreat,
		resp:   &resp,
		body:   &entry,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
