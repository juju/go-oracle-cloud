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
	ErrAlreadyAuth    = errors.New("go-oracle-cloud: The client is already authenticated")
	ErrNotAuth        = errors.New("go-oracle-cloud: The client is not authenticated")
	ErrUndefinedShape = errors.New("go-oracle-cloud: The shape specified is undefined")
)

func dumpApiError(body io.Reader) string {
	var e response.Error
	json.NewDecoder(body).Decode(&e) // skip error
	return e.Message
}
