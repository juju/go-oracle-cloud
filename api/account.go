package api

import (
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

func (c Client) AccountDetails(name string) (resp response.AccountDetails, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	// build the url for the api endpoint
	url := fmt.Sprintf("%s/%s/Compute-%s/%s", c.endpoint, "account", c.identify, name)

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

func (c Client) Account(name string) (resp response.Account, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	// build the url for the api endpoint
	url := fmt.Sprintf("%s/%s/", c.endpoint, "account")
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
