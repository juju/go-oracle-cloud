package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// Creates a security IP list.
// Note that, after creating a security IP list, you can add additional
// IP addresses to the list by using the CreateIpSecList again with just
// the additional IP addresses.
func (c Client) CreateIpSecList(
	description string,
	name string,
	secipentries []string,
) (resp response.IpSecList, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty secure ip list name",
		)
	}

	params := struct {
		Description  string   `json:"description,omitempty"`
		Name         string   `json:"name"`
		Secipentries []string `json:"secipentries"`
	}{
		Description: description,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, name),
		Secipentries: secipentries,
	}

	url := fmt.Sprintf("%s/seciplist/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// Deletes the specified security IP list. No response is returned.
// You can't delete system-provided security application that are
// available in the /oracle/public container.
func (c Client) DeleteIpSecList(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty sec ip list name")
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
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

// IpSecListDetail retrieves information about the specified security IP list.
// You can use this request to verify whether CreateIpSecList
// and UpdateIpSecList operations were completed successfully.
func (c Client) IpSecListDetail(name string) (resp response.IpSecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty sec ip list name",
		)
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) AllIpSecList() (resp response.AllIpSecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates IP addresses and description of the specified security IP list.
// Note that this command replaces the values in the secipentries and
// description fields with the new values that you specify.
// To add one or more IP addresses to the existing list, run
// the add seciplist command and specify just the additional IP addresses.
func (c Client) UpdateIpSecList(
	description string,
	currentName string,
	newName string,
	secipentries []string,
) (resp response.IpSecList, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty current ip sec list name",
		)
	}

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Description  string   `json:"description,omitempty"`
		Name         string   `json:"name"`
		Secipentries []string `json:"secipentries"`
	}{
		Description: description,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
		Secipentries: secipentries,
	}

	url := fmt.Sprintf("%s/seciplist/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "PUT",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil

}
