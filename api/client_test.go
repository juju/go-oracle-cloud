package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/hoenirvili/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

type clientTest struct {
	server *httptest.Server
}

var _ = gc.Suite(&clientTest{})

// NewConfig returns a predefined config for testing
// the client interactions
func (cli clientTest) NewConfig(c *gc.C) api.Config {
	return api.Config{
		Username: "oracleusername@oracle.com",
		Password: "Password123",
		Identify: "myIdentify",
		Endpoint: "http://localhost",
	}
}

type testRecorder func(c *gc.C, record *httptest.ResponseRecorder)

// TODO(write test server)
func (cli *clientTest) NewServer() {
	cli.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/oracle-compute-v3+json")
		w.WriteHeader(http.StatusOK)
		fmt.Printf("%+v\n", r)
		fmt.Fprintf(w, "hi")
	}))
}

func (cli *clientTest) Clear() {
	cli.server.Close()
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
	cli, err = api.NewClient(cl.NewConfig(c))
	c.Assert(err, gc.IsNil)
	c.Assert(cli, gc.NotNil)
}
