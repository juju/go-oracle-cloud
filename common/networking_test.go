// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package common_test

import (
	"strings"

	"github.com/juju/go-oracle-cloud/common"

	gc "gopkg.in/check.v1"
)

func (commTest) TestNicType(c *gc.C) {
	nic := common.Nic{
		Vethernet: "random",
	}

	t := nic.GetType()
	c.Assert(t, gc.DeepEquals, common.VEthernet)

	nic = common.Nic{
		Vethernet: "",
	}

	t = nic.GetType()
	c.Assert(t, gc.DeepEquals, common.VNic)
}

func (commTest) TestNewIpPool(c *gc.C) {
	t := common.NewIPPool(common.PublicIPPool, common.IPReservationType)
	p := strings.Split(string(t), ":")
	c.Assert(p, gc.NotNil)
	c.Assert(p[0], gc.Equals, "ipreservation")
	c.Assert(common.IPPool(p[1]), gc.Equals, common.PublicIPPool)
}

func (commTest) TestProtocolValidity(c *gc.C) {
	var p common.Protocol
	c.Assert(p.Validate(), gc.NotNil)
}
