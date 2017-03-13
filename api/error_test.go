package api_test

import (
	"github.com/juju/go-oracle-cloud/api"

	gc "gopkg.in/check.v1"
)

func (cl clientTest) TestIsNotAuth(c *gc.C) {
	err := &api.ErrNotAuth{}
	ok := api.IsNotAuth(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestIsNotFound(c *gc.C) {
	err := &api.ErrNotFound{}
	ok := api.IsNotFound(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestIsBadRequest(c *gc.C) {
	err := &api.ErrBadRequest{}
	ok := api.IsBadRequest(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestIsNotAuthorized(c *gc.C) {
	err := &api.ErrNotAuthorized{}
	ok := api.IsNotAuthorized(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestIsInternalApi(c *gc.C) {
	err := &api.ErrInternalApi{}
	ok := api.IsInternalApi(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestIsStatusConflic(c *gc.C) {
	err := &api.ErrStatusConflict{}
	ok := api.IsStatusConflict(err)
	c.Assert(ok, gc.Equals, true)
}

func (cl clientTest) TestAllEmptyErrors(c *gc.C) {
	funcs := []func(error) bool{
		api.IsNotAuth,
		api.IsNotFound,
		api.IsBadRequest,
		api.IsNotAuthorized,
		api.IsInternalApi,
		api.IsStatusConflict,
	}

	var ok bool
	for _, f := range funcs {
		ok = f(nil)
		c.Assert(ok, gc.Equals, false)
	}
}
