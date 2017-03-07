// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/hoenirvili/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

type clientTest struct {
	srv *httptest.Server
	c   *api.Client
}

var _ = gc.Suite(&clientTest{})

// DefaultConfig returns a predefined config for testing
// the client interactions
func (cli clientTest) DefaultConfig() api.Config {
	return api.Config{
		Username: "oracleusername@oracle.com",
		Password: "Password123",
		Identify: "myIdentify",
		Endpoint: "http://localhost",
	}
}

// NewClient returns a a new client default client for testing
func (cli clientTest) NewClient(c *gc.C) *api.Client {
	cfg := cli.DefaultConfig()
	client, err := api.NewClient(cfg)
	c.Assert(err, gc.IsNil)
	c.Assert(client, gc.NotNil)
	return client
}

// handler_test type used for defining a handler test method
type handler_test func(w http.ResponseWriter, r *http.Request)

type httpParams struct {
	c *gc.C

	// handler is the handler
	handler handler_test
	// the json structure where the decoding will point to
	marshall interface{}
	// http status expected
	status int

	// the url that the server will
	// response to
	url string
}

// Start will start a http server for testing the handler_test
func (cli *clientTest) Start(h httpParams) {

	// if the client is not inited init it with the default clientr
	if cli.c == nil {
		cli.c = cli.NewClient(h.c)
	}

	cli.srv = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				fmt.Println(w)
				fmt.Println(*r)
				if h.status != 0 {
					w.WriteHeader(h.status)
				}

				if h.marshall != nil {
					err := json.NewDecoder(r.Body).Decode(h.marshall)
					h.c.Assert(err, gc.IsNil)
					h.c.Assert(r.Body.Close(), gc.IsNil)
				}

				if h.handler != nil {
					h.handler(w, r)
				}
			},
		),
	)

	h.c.Assert(cli.srv, gc.NotNil)
}

// Stop will close the http test server
func (cli *clientTest) Stop() {
	cli.srv.Close()
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
	cfg := cl.DefaultConfig()
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
