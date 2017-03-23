// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package common_test

import (
	"github.com/juju/go-oracle-cloud/common"

	gc "gopkg.in/check.v1"
)

type commTest struct{}

var _ = gc.Suite(&commTest{})

func (commTest) TestNewInterval(c *gc.C) {
	value := 20
	interval := common.NewInterval(value)
	c.Assert(interval.Hourly.HourlyInterval, gc.DeepEquals, value)

	err := interval.Validate()
	c.Assert(err, gc.IsNil)
}

func (commTest) TestNewDailyWeekly(c *gc.C) {
	weeks := []common.Week{
		common.Sunday,
		common.Monday,
	}
	time := "03:15"
	timezone := "America/Los_angeles"

	daily := common.NewDailyWeekly(weeks, time, timezone)

	c.Assert(daily.DaysOfWeek, gc.DeepEquals, weeks)
	c.Assert(daily.TimeOfDay, gc.DeepEquals, time)
	c.Assert(daily.UserTimeZone, gc.DeepEquals, timezone)

	err := daily.Validate()
	c.Assert(err, gc.IsNil)
}
