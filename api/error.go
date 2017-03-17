// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	// errNotAuth error that represents that the client it not authenticated
	errNotAuth = &ErrNotAuth{
		message: "go-oracle-cloud: Client is not authenticated",
	}

	// errBadRequest error meaning the given request used
	// is invalid or constructed bad
	errBadRequest = &ErrBadRequest{
		message: "go-oracle-cloud: The request given is invalid",
	}

	// errNotAuthorized error meaning that the client does not
	// have authorization to access that specific resource
	errNotAuthorized = &ErrNotAuthorized{
		message: "go-oracle-cloud: Client does not have authorization to access this resource",
	}

	// errInternalAPI error meaning the oracle rest server has encountered
	// an internal error
	errInternalApi = &ErrInternalApi{
		message: "go-oracle-cloud: Oracle infrstracutre has encountered an internal error",
	}

	// errStatusConflict error meaning that the current association or relation could not be
	// created
	errStatusConflict = &ErrStatusConflict{
		message: "go-oracle-cloud: Some association isn't right or object already",
	}

	// errNotFound error meaning that the resource the client want's to access is not found
	errNotFound = &ErrNotFound{
		message: "go-oracle-cloud: The resource you requested is not found",
	}
)

// ErrorDumper interface that represents the
// ability to dump the error in a go error format
// from a given reader
type ErrorDumper interface {
	DumpApiError(r io.Reader) error
}

// ErrNotAuth error type that implements the error interface
type ErrNotAuth struct{ message string }

// Error returns the internal error in a string format
func (e ErrNotAuth) Error() string { return e.message }

// errNotAuthorized error type that implements the error and
// ErrorDumper interfaces
type ErrNotAuthorized struct{ message string }

// Error returns the internal error in a string format
func (e ErrNotAuthorized) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrNotAuthorized) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e

}

// ErrBadRequest error type that implements the error and
// ErrorDumper interfaces
type ErrBadRequest struct{ message string }

// Error returns the internal error in a string format
func (e ErrBadRequest) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrBadRequest) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

// ErrInternalApi error type that implements the error and
// ErrorDumper interfaces
type ErrInternalApi struct{ message string }

// Error returns the internal error in a string format
func (e ErrInternalApi) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrInternalApi) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

// ErrNotFound error type that implements the error and
// ErrorDumper interfaces
type ErrNotFound struct{ message string }

// Error returns the internal error in a string format
func (e ErrNotFound) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrNotFound) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

// ErrStatusConflict error type that implements the error and
// ErrorDumper interfaces
type ErrStatusConflict struct{ message string }

// Error returns the internal error in a string format
func (e ErrStatusConflict) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrStatusConflict) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

// dumpApiError used in the callback request custom handlers
func dumpApiError(resp *http.Response) error {
	body, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf(
		"go-oracle-cloud: Error api response %d ,Raw: %s",
		resp.StatusCode, string(body),
	)
}

// IsNotAuth returns true if the error
// indicates that the client is not authenticate
func IsNotAuth(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrNotAuth)
	return ok
}

// IsNotFound returns true if the error
// indicates an http not found 404
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrNotFound)
	return ok
}

// IsBadRequest returns true if the error
// indicates an http bad request 400
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrBadRequest)
	return ok
}

// IsNotAuthorized returns true if the error
// indicates an http unauthorized 401
func IsNotAuthorized(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrNotAuthorized)
	return ok
}

// IsInternalApi returns true if the error
// indicated an http internal service error 500
func IsInternalApi(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrInternalApi)
	return ok
}

// IsStatusConflict returns true if the error
// indicates an http conflict error 401
func IsStatusConflict(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrStatusConflict)
	return ok
}
