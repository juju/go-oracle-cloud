package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

func (c Client) ImageListDetail(name string) (resp response.ImageList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-api: Empty image name")
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "imagelist", c.identify, c.username, name)

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

func (c Client) AllImageList() (resp response.AllImageList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "imagelist", c.identify, c.username)

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

func (c Client) AllImageListNames() (resp response.AllImageListNames, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "imagelist", c.identify, c.username)

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

func (c Client) AddImageList(def uint64, description string, name string) (resp response.ImageList, err error) {

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
		Name:        fmt.Sprintf("/Compute-%s/%s/%s", c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/",
		c.endpoint, "imagelist", c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		body:   &image,
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusCreated {
				return fmt.Errorf("go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
			return nil
		},
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) DeleteImageList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-api: Empty image name")
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "imagelist", c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusNoContent {
				return fmt.Errorf("go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)

			}
			return nil
		},
	}); err != nil {
		return err
	}

	return nil
}

func (c Client) UpdateImageList(target, name, description string, def uint64) (err error) {

	if !c.isAuth() {
		return ErrNotAuth
	}

	if target == "" {
		return errors.New("go-oracle-cloud: Cannot update image list because of empty target")
	}

	image := struct {
		Def         uint64 `json:"default"`
		Description string `json:"description"`
		Name        string `json:"name"`
	}{
		Def:         def,
		Description: description,
		Name:        fmt.Sprintf("/Compute-%s/%s/%s", c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "imagelist", c.identify, c.username, target)

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

func (c Client) AllImageListEntries(entryName string) (resp response.AllImageListEntries, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if entryName == "" {
		return resp, errors.New("go-oracle-api: Empty image list entry name")
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s",
		c.endpoint, "imagelist", c.identify, c.username, entryName)

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

type AddImagelistParams struct {
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Version       uint64                 `json:"version"`
	MachineImages []string               `json:"machineimages"`
	Uri           string                 `json:"uri"`
}

//TODO
func (c Client) AddImageListEntries(name string, params AddImagelistParams) (resp response.AllImageListEntries, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-api: Empty image list entry name")
	}

	url := fmt.Sprintf("%s/%s/Compute-%s/%s/%s/%s/",
		c.endpoint, "imagelist", c.identify, c.username, name, "entry")

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "POST",
		url:    url,
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusCreated {
				return fmt.Errorf("go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
			return nil
		},
		resp: &resp,
		body: &params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
