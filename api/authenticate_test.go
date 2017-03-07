// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api_test

import (
	"net/http"

	gc "gopkg.in/check.v1"
)

func (cl clientTest) TestAuthentication(c *gc.C) {
	auth := map[string]string{}

	cl.Start(httpParams{
		marshall: &auth,
		status:   http.StatusOK,
		handler:  func(w http.ResponseWriter, r *http.Request) {},
	})

	// TODO(shit isn't working..??!?!!")
	err := cl.c.Authenticate()
	c.Assert(err, gc.NotNil)
	//cl.Stop()
}
