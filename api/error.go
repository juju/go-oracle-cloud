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

	errForbidden = &ErrForbidden{
		message: "go-oracle-cloud: The server understood the request but refuses to authorize it",
	}
	errPaymentRequired = &ErrPaymentRequired{
		message: "go-oracle-cloud: The status code is reserved for future use",
	}

	errMethodNotAllowed = &ErrMethodNotAllowed{
		message: "go-oracle-cloud: The method received in the request-line is known by the origin server but not supported by the target resource",
	}

	errNotAcceptable = &ErrNotAcceptable{
		message: "go-oracle-cloud: The target resource does not have a current representation that would be acceptable to the user agent",
	}

	errRequestTimeout = &ErrRequestTimeout{
		message: "go-oracle-cloud: The server did not recive complete request within the time that it was prepared to wait",
	}

	errGone = &ErrGone{
		message: "go-oracle-cloud: The request could not be completed due to a conflict with the current state of the target resource",
	}
	errLengthRequired = &ErrLengthRequired{
		message: "go-oracle-cloud: The server refuses to accept the request without a defined Content-Length",
	}

	errRequestEntityTooLarge = &ErrRequestEntityTooLarge{
		message: "go-oracle-cloud: the server is refusing to process a request because the request payload is larger than the server is willing or able to process",
	}
)

// ErrRequestEntityTooLarge indicates that the server is
// refusing to process a request because the request payload is larger
// than the server is willing or able to process
type ErrRequestEntityTooLarge struct{ message string }

// Error returns the internal error in a string format
func (e ErrRequestEntityTooLarge) Error() string {
	return e.message
}

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrRequestEntityTooLarge) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrLengthRequired indicates that the server
// refuses to accept the request without a defined Content-Length
type ErrLengthRequired struct{ message string }

// Error returns the internal error in a string format
func (e ErrLengthRequired) Error() string {
	return e.message
}

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrLengthRequired) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

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
	e.message = e.message + " Raw: " + string(body)
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
	e.message = e.message + " Raw: " + string(body)
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
	e.message = e.message + " Raw: " + string(body)
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
	e.message = e.message + " Raw: " + string(body)
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
	e.message = e.message + " Raw: " + string(body)
	return e
}

// dumpApiError used in the callback request custom handlers
func dumpApiError(resp *http.Response) error {
	body, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf(
		"go-oracle-cloud: Error api response %d Raw: %s",
		resp.StatusCode, string(body),
	)
}

// ErrForbidden error that says the server understood the request,
// but is refusing to fulfill it. Authorization will not help
// and the request SHOULD NOT be repeated
type ErrForbidden struct{ message string }

// Error returns type that implements the error interface
func (e ErrForbidden) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrForbidden) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrPaymentRequired error that says the status code
// is reserved for future use
type ErrPaymentRequired struct{ message string }

// Error returns type that implements the error interface
func (e ErrPaymentRequired) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrPaymentRequired) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrStatusGone that the request could not be completed due
// to a conflict with the current state of the target resource
type ErrStatusGone struct{ message string }

// Error returns the internal error in a string format
func (e ErrStatusGone) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e ErrStatusGone) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrRequestTimeout indicates that the server did
// not receive a complete request message within the time that it was
// prepared to wait
type ErrRequestTimeout struct{ message string }

// Error returns the internal error in a string format
func (e ErrRequestTimeout) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrRequestTimeout) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e

}

// ErrNotAcceptable indicates that the target
// resource does not have a current representation that would be
// acceptable to the user agent
type ErrNotAcceptable struct{ message string }

// Error returns the internal error in a string format
func (e ErrNotAcceptable) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e ErrNotAcceptable) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrMethodNotAllowed indicates that the method
// received in the request-line is known by the origin server but not
// supported by the target resource
type ErrMethodNotAllowed struct{ message string }

// Error returns the internal error in a string format
func (e ErrMethodNotAllowed) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e *ErrMethodNotAllowed) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e
}

// ErrGone indicates that the request could not be completed due
// to a conflict with the current state
// of the target resource"
type ErrGone struct{ message string }

// Error returns the internal error in a string format
func (e ErrGone) Error() string { return e.message }

// DumpApiError returns the error in a error format from a given reader source
func (e ErrGone) DumpApiError(r io.Reader) error {
	body, _ := ioutil.ReadAll(r)
	e.message = e.message + " Raw: " + string(body)
	return e

}

// IsForbidden returns true if the error is
// ErrForbidden type
func IsForbidden(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrForbidden)
	return ok
}

// IsPaymentRequired returns true if the error is
// ErrPaymentRequired type
func IsPaymentRequired(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrPaymentRequired)
	return ok
}

// IsMethodNotAllowed returns true if the error is
// ErrMethodNotAllowed type
func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrMethodNotAllowed)
	return ok
}

// IsNotAcceptable returns true if the error is
// ErrNotAcceptable type
func IsNotAcceptable(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrNotAcceptable)
	return ok
}

// IsLengthRequired returns true if the error is
// ErrLengthRequired type
func IsLengthRequired(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrLengthRequired)
	return ok
}

// IsRequestEntityTooLarge returns true if the underlying
// error is a type of ErrRequestEntityTooLarge
func IsRequestEntityTooLarge(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*ErrRequestEntityTooLarge)
	return ok
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
