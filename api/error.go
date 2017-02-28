// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

var (
	errNotAuth = &ErrNotAuth{
		message: "go-oracle-cloud: Client is not authenticated",
	}

	errAlreadyAuth = &ErrAlreadyAuth{
		message: "go-oracle-cloud: Client is already authenticated",
	}

	errBadRequest = &ErrBadRequest{
		message: "go-oracle-cloud: The request given is invalid",
	}

	errNotAuthorized = &ErrNotUnauthorized{
		message: "go-oracle-cloud: Client does not have authorization to access this resource",
	}

	errInternalApi = &ErrInternalApi{
		message: "go-oracle-cloud: Oracle infrstracutre has encountered an internal error",
	}

	errStatusConflict = &ErrStatusConflict{
		message: "go-oracle-cloud: Some association isn't right or object already",
	}

	errNotFound = &ErrNotFound{
		message: "go-oracle-cloud: The resource you requested is not found",
	}
)

type ErrAlreadyAuth struct{ message string }

func (e ErrAlreadyAuth) Error() string { return e.message }

type ErrNotAuth struct{ message string }

func (e ErrNotAuth) Error() string { return e.message }

type ErrNotUnauthorized struct{ message string }

func (e ErrNotUnauthorized) Error() string { return e.message }

type ErrBadRequest struct{ message string }

func (e ErrBadRequest) Error() string { return e.message }

type ErrInternalApi struct{ message string }

func (e ErrInternalApi) Error() string { return e.message }

type ErrNotFound struct{ message string }

func (e ErrNotFound) Error() string { return e.message }

type ErrStatusConflict struct{ message string }

func (e ErrStatusConflict) Error() string { return e.message }

// dumpApiError used in the callback request custom handlers
func dumpApiError(resp *http.Response) error {
	var e response.Error
	// skip the error in decoding part
	json.NewDecoder(resp.Body).Decode(&e)
	return fmt.Errorf(
		"go-oracle-cloud: Error api response %d %s", resp.Status, e.Message,
	)
}

func IsAlreadyAuth(err error) bool {
	_, ok := err.(*ErrAlreadyAuth)
	return ok
}

func IsNotAuth(err error) bool {
	_, ok := err.(*ErrNotAuth)
	return ok
}

func IsNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	return ok
}

func IsBadRequest(err error) bool {
	_, ok := err.(*ErrBadRequest)
	return ok
}

func IsNotUnauthorized(err error) bool {
	_, ok := err.(*ErrNotUnauthorized)
	return ok
}

func IsInternalApi(err error) bool {
	_, ok := err.(*ErrInternalApi)
	return ok
}

func IsStatusConflict(err error) bool {
	_, ok := err.(*ErrStatusConflict)
	return ok
}
