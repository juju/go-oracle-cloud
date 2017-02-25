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
	//ErrAlreadyAuth error used when you use more than
	// once the Authentication method of the client
	ErrAlreadyAuth = errors.New("go-oracle-cloud: The client is already authenticated")
	// ErrNotAuth error returned by the client if the Authentication method is not used
	ErrNotAuth = errors.New("go-oracle-cloud: The client is not authenticated")
)

// dumpApiError used in the callback request custom handlers
func dumpApiError(body io.Reader) string {
	var e response.Error
	// skip the error in decoding part
	json.NewDecoder(body).Decode(&e)
	return e.Message
}
