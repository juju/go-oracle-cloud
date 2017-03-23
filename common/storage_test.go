// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package common_test

import (
	"strings"

	"github.com/juju/go-oracle-cloud/common"

	gc "gopkg.in/check.v1"
)

func (commTest) TestIndexValidity(c *gc.C) {

	for i := common.Index(1); i < common.Index(11); i++ {
		c.Assert(i.Validate(), gc.IsNil)
	}

	i := common.Index(0)
	c.Assert(i.Validate(), gc.NotNil)

	i = common.Index(11)
	c.Assert(i.Validate(), gc.NotNil)
}

func (commTest) TestStoragePoolValidity(c *gc.C) {
	var s common.StoragePool
	c.Assert(s.Validate(), gc.NotNil)
}

func (commTest) TestNewStorageSize(c *gc.C) {
	size := common.NewStorageSize(1024, common.T)
	ok := strings.Compare(string(size), "1024T")
	c.Assert(ok, gc.Equals, 0)
}
