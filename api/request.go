package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// treatStatus will be used as a callback to custom check the response
// if the client decides the response contains some error codes
// it will make the Request return that error
type treatStatus func(resp *http.Response) error

// paramsRequest used to fill up the params for the request function
type paramsRequest struct {
	// use this client to do the request
	client *http.Client
	// include in the request the cookie
	cookie *http.Cookie
	// url this will not be checked so, make sure it is valid
	url string
	// verb contains an http verb POST, GET, PUT, PATH
	verb string
	// body will contain the http body request if any
	// for Get request this should be leaved nil
	body interface{}
	// treat will be used as a callback to
	// check the response and save extract things
	treat treatStatus
	// resp will contains a json object where the
	// request function will decode
	resp interface{}
}

func request(cfg paramsRequest) (err error) {
	var buf io.Reader

	// if we have a body that meas we must
	if cfg.body != nil {
		raw, err := json.Marshal(cfg.body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(raw)
	}

	req, err := http.NewRequest(cfg.verb, cfg.url, buf)
	if err != nil {
		return err
	}

	if cfg.cookie != nil {
		// add the session cookie
		req.AddCookie(cfg.cookie)
	}

	req.Header.Add("User-Agent", "go-oracle-cloud v1.0")
	switch strings.ToUpper(cfg.verb) {
	case "POST", "DELETE", "PUT":
		req.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	case "GET":
		req.Header.Add("Accept", "application/oracle-compute-v3+json")
	}

	resp, err := cfg.client.Do(req)
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			err = errClose
		}
	}()
	if err != nil {
		return err
	}

	if cfg.treat != nil {
		if err = cfg.treat(resp); err != nil {
			return err
		}
	}

	if cfg.resp != nil {
		// decode the JSON from the response if any
		if err = json.NewDecoder(resp.Body).Decode(cfg.resp); err != nil {
			return err
		}
	}

	return nil
}
