// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package common_test

import (
	"github.com/juju/go-oracle-cloud/common"

	gc "gopkg.in/check.v1"
)

func (commTest) TestSecRuleActionValidity(c *gc.C) {
	var s common.SecRuleAction
	c.Assert(s.Validate(), gc.NotNil)
}
