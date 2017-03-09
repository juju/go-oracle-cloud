// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"fmt"

	"github.com/juju/go-oracle-cloud/api"
	gc "gopkg.in/check.v1"
)

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
