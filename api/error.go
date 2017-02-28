// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	errNotAuth = &ErrNotAuth{
		message: "go-oracle-cloud: Client is not authenticated ",
	}

	errAlreadyAuth = &ErrAlreadyAuth{
		message: "go-oracle-cloud: Client is already authenticated ",
	}

	errBadRequest = &ErrBadRequest{
		message: "go-oracle-cloud: The request given is invalid ",
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

type ErrorDumper interface {
	DumpApiError(r io.Reader) error
}

type ErrAlreadyAuth struct{ message string }

func (e ErrAlreadyAuth) Error() string { return e.message }

type ErrNotAuth struct{ message string }

func (e ErrNotAuth) Error() string { return e.message }

type ErrNotAuthorized struct{ message string }

func (e ErrNotAuthorized) Error() string { return e.message }
func (e *ErrNotAuthorized) DumpApiError(r io.Reader) error {
	var suffix string
	json.NewDecoder(r).Decode(&suffix)
	e.message = e.message + ",Raw " + suffix
	return e

}

type ErrBadRequest struct{ message string }

func (e ErrBadRequest) Error() string { return e.message }
func (e *ErrBadRequest) DumpApiError(r io.Reader) error {
	var suffix string
	json.NewDecoder(r).Decode(&suffix)
	e.message = e.message + ",Raw " + suffix
	return e
}

type ErrInternalApi struct{ message string }

func (e ErrInternalApi) Error() string { return e.message }
func (e *ErrInternalApi) DumpApiError(r io.Reader) error {
	var suffix string
	json.NewDecoder(r).Decode(&suffix)
	e.message = e.message + ",Raw " + suffix
	return e
}

type ErrNotFound struct{ message string }

func (e ErrNotFound) Error() string { return e.message }
func (e *ErrNotFound) DumpApiError(r io.Reader) error {
	var suffix string
	json.NewDecoder(r).Decode(&suffix)
	e.message = e.message + ",Raw " + suffix
	return e
}

type ErrStatusConflict struct{ message string }

func (e ErrStatusConflict) Error() string { return e.message }
func (e *ErrStatusConflict) DumpApiError(r io.Reader) error {
	var suffix string
	json.NewDecoder(r).Decode(&suffix)
	e.message = e.message + ",Raw " + suffix
	return e
}

// dumpApiError used in the callback request custom handlers
func dumpApiError(resp *http.Response) error {
	var message string
	// skip the error in decoding part
	json.NewDecoder(resp.Body).Decode(&message)
	return fmt.Errorf(
		"go-oracle-cloud: Error api response %d , Raw: %s", resp.Status, message,
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
