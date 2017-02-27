// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

var (
	//ErrAlreadyAuth error used when you use more than
	// once the Authentication method of the client
	ErrAlreadyAuth = errors.New("go-oracle-cloud: The client is already authenticated")
	// ErrNotAuth error returned by the client if the Authentication method is not used
	ErrNotAuth = errors.New("go-oracle-cloud: The client is not authenticated")

	ErrBadRequest     = errors.New("go-oracle-cloud: The request given is invalid")
	ErrUnathorized    = errors.New("go-oracle-cloud: Client does not have authorisation for this resource")
	ErrInternalApi    = errors.New("go-oracle-cloud: The cloud has encountered an error")
	ErrNotFound       = errors.New("go-oracle-cloud: The resource you requested is not found")
	ErrStatusConflict = errors.New("go-oracle-cloud: Some association isn't right, please double check the request")
)

// dumpApiError used in the callback request custom handlers
func dumpApiError(resp *http.Response) error {
	var e response.Error
	// skip the error in decoding part
	json.NewDecoder(resp.Body).Decode(&e)
	return fmt.Errorf(
		"go-oracle-cloud: Error api response %d %s", resp.Status, e.Message,
	)
}
