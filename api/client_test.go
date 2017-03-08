// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api_test

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"

	"github.com/hoenirvili/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

type clientTest struct{}

var _ = gc.Suite(&clientTest{})

// findHostAndPortFromURL extracts the
// host and port part of a full url
// (like http://address:port/path#fragment)
func findHostAndPortFromURL(rawurl string) (string, int, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", 0, err
	}
	host, port, err := net.SplitHostPort(url.Host)
	if err != nil {
		return "", 0, err
	}
	iport, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, err
	}
	return host, iport, nil
}

type httpParams struct {
	// body for providing the test function
	// to send content back to the api client
	body []byte

	// check is the assert library
	check *gc.C

	// in contains the post/put body
	in interface{}

	// handler for manual testing the
	// http.ResponseWriter and http.Reqeust
	handler http.HandlerFunc

	// handler manual the header
	// and header status
	manualHeaderStatus bool
}

const (
	// content-type values for json requests and directory request
	json = "application/oracle-compute-v3+json"
	dir  = "application/oracle-compute-v3+directory+json"
)

const (
	username = "gooraclecloudadmin@oracle.com"
	password = "some1337eleetpasswd"
	identify = "qraffd"
)

func hdr(value string, w http.ResponseWriter, r *http.Request) (err error) {
	switch value {
	case json:
		w.Header().Set("Accept", json)
		w.Header().Set("Content-Type", json)
		return nil
	case dir:
		w.Header().Set("Accept", dir)
		w.Header().Set("Content-Type", dir)
		return nil
	default:
		return fmt.Errorf("oracle api does not support this header value %s", value)
	}
}

func handlerHeaderStatus(w http.ResponseWriter, r *http.Request) (err error) {
	value := r.Header.Get("Content-Type")
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		return hdr(value, w, r)
	case http.MethodGet, http.MethodPut:
		w.WriteHeader(http.StatusOK)
		return hdr(value, w, r)
	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)
		return hdr(value, w, r)
	default:
		return fmt.Errorf("oracle api does not support this method %s", r.Method)
	}
}

// StartTestServer will start an httptest server on a random port
// with the given httpParams and then return the oracle client implementation
func (cli clientTest) StartTestServer(
	params httpParams,
) (*httptest.Server, *api.Client) {
	// create a new http server for testing
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if !params.manualHeaderStatus {
				// handler the header if it's somthing wrong then
				// handler the status based on the method
				err := handlerHeaderStatus(w, r)
				params.check.Assert(err, gc.IsNil)
			}

			// treat the w and *r here in a custom way
			if params.handler != nil {
				params.handler(w, r)
			}

			// if the caller is expecting the
			// server to send some response,
			// we should include that response from body
			// and send it back
			m := len(params.body)
			if m > 0 {
				// write the response
				n, err := w.Write(params.body)
				params.check.Assert(err, gc.IsNil)
				params.check.Assert(m, gc.Equals, n)
			}

			defer func() {
				err := r.Body.Close()
				params.check.Assert(err, gc.IsNil)
			}()
		}))

	// find the host and port from the http server url
	host, port, err := findHostAndPortFromURL(ts.URL)
	// always test if the finding is valid
	params.check.Assert(err, gc.IsNil)

	// create a new config
	cfg := api.Config{
		Username: username,
		Password: password,
		Identify: identify,
		Endpoint: fmt.Sprintf("http://%s:%d", host, port),
	}

	// create a new client based on the config
	client, err := api.NewClient(cfg)
	params.check.Assert(err, gc.IsNil)
	params.check.Assert(client, gc.NotNil)

	// return the new started server and the oracle client
	return ts, client
}

func (cl clientTest) TestNewClient(c *gc.C) {
	cli, err := api.NewClient(api.Config{})
	c.Assert(err, gc.NotNil)
	c.Assert(cli, gc.IsNil)
	c.Assert(err.Error(), gc.DeepEquals,
		"go-oracle-cloud: Empty identify endpoint name")

	cli, err = api.NewClient(api.Config{
		Identify: "someIdentify",
	})
	c.Assert(err, gc.NotNil)
	c.Assert(cli, gc.IsNil)
	c.Assert(err.Error(), gc.DeepEquals,
		"go-oracle-cloud: Empty client username")

	cli, err = api.NewClient(api.Config{
		Identify: "someIdentify",
		Username: "sometestuser",
	})
	c.Assert(err, gc.NotNil)
	c.Assert(cli, gc.IsNil)
	c.Assert(err.Error(), gc.DeepEquals,
		"go-oracle-cloud: Empty client password")

	cli, err = api.NewClient(api.Config{
		Identify: "someIdentify",
		Username: "sometestuser",
		Password: "providesomepasswrd",
	})
	c.Assert(err, gc.NotNil)
	c.Assert(cli, gc.IsNil)
	c.Assert(err.Error(), gc.DeepEquals,
		"go-oracle-cloud: Empty endpoint url basepath")

	cli, err = api.NewClient(api.Config{
		Identify: "someIdentify",
		Username: "sometestuser",
		Password: "providesomepasswrd",
		Endpoint: "s",
	})

	c.Assert(err, gc.NotNil)
	c.Assert(cli, gc.IsNil)
	c.Assert(err.Error(), gc.DeepEquals,
		"go-oracle-cloud: The endpoint provided is invalid")

	// provide some valid configuration
	cfg := api.Config{
		Username: "OracleAdmin@oracle.com",
		Password: "GreatSuperPassword",
		Identify: "qtqqd",
		Endpoint: "http://edz.api55.oracle.com",
	}
	cli, err = api.NewClient(cfg)
	c.Assert(err, gc.IsNil)
	c.Assert(cli, gc.NotNil)
	c.Assert(cli.Identify(), gc.DeepEquals, cfg.Identify)
	c.Assert(cli.Username(), gc.DeepEquals, cfg.Username)
	c.Assert(cli.Password(), gc.DeepEquals, cfg.Password)

	name := cli.ComposeName("someName")
	c.Assert(name, gc.DeepEquals,
		fmt.Sprintf("/Compute-%s/%s/%s",
			cli.Identify(), cli.Username(), "someName"))
}
