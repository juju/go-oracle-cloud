// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

var (
	ErrAlreadyAuth = errors.New("go-oracle-cloud: The client is already authenticated")
	ErrNotAuth     = errors.New("go-oracle-cloud: The client is not authenticated")
)

// dumpApiError used in the callback request custom handlers
func dumpApiError(body io.Reader) string {
	var e response.Error
	json.NewDecoder(body).Decode(&e) // skip error
	return e.Message
}
