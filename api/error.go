// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	errNotAuth = &ErrNotAuth{
		message: "go-oracle-cloud: Client is not authenticated",
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

type ErrorDumper interface {
	DumpApiError(r io.Reader) error
}

type ErrNotAuth struct{ message string }

func (e ErrNotAuth) Error() string { return e.message }

type ErrNotAuthorized struct{ message string }

func (e ErrNotAuthorized) Error() string { return e.message }
func (e *ErrNotAuthorized) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e

}

type ErrBadRequest struct{ message string }

func (e ErrBadRequest) Error() string { return e.message }
func (e *ErrBadRequest) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

type ErrInternalApi struct{ message string }

func (e ErrInternalApi) Error() string { return e.message }
func (e *ErrInternalApi) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

type ErrNotFound struct{ message string }

func (e ErrNotFound) Error() string { return e.message }
func (e *ErrNotFound) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " ,Raw: " + string(body)
	return e
}

type ErrStatusConflict struct{ message string }

func (e ErrStatusConflict) Error() string { return e.message }
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
