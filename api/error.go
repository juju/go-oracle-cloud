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

	errNotAuthorized = &ErrNotAuthorized{
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

type ErrNotAuthorized struct{ message string }

func (e ErrNotAuthorized) Error() string { return e.message }

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

// IsAlreadyAuth returns true if the error
// indicates an already already authenticated client
func IsAlreadyAuth(err error) bool {
	_, ok := err.(*ErrAlreadyAuth)
	return ok
}

// IsNotAuth returns true if the error
// indicates that the client is not authenticate
func IsNotAuth(err error) bool {
	_, ok := err.(*ErrNotAuth)
	return ok
}

// IsNotFound returns true if the error
// indicates an http not found 404
func IsNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	return ok
}

// IsBadRequest returns true if the error
// indicates an http bad request 400
func IsBadRequest(err error) bool {
	_, ok := err.(*ErrBadRequest)
	return ok
}

// IsNotAuthorized returns true if the error
// indicates an http unauthorized 401
func IsNotAuthorized(err error) bool {
	_, ok := err.(*ErrNotAuthorized)
	return ok
}

// IsInternalApi returns true if the error
// indicated an http internal service error 500
func IsInternalApi(err error) bool {
	_, ok := err.(*ErrInternalApi)
	return ok
}

// IsStatusConflict returns true if the error
// indicates an http conflict error 401
func IsStatusConflict(err error) bool {
	_, ok := err.(*ErrStatusConflict)
	return ok
}
