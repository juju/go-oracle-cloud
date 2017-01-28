package api

import (
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

func (c Client) Account(name string) (resp response.Account, err error) {
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
		body:   nil,
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
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
