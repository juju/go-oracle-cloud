// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

import (
	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/response"
)

var (
	orchestrationParams = api.OrchestrationParams{}

	orchestrationDetails = response.Orchestration{}

	allorchestrations = response.AllOrchestrations{
		Result: []response.Orchestration{
			orchestrationDetails,
		},
	}
)

//TODO(sgiulitti) unit tests
