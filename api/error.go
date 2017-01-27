package api

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

var (
	ErrAlreadyAuth = errors.New("The client is already authenticated")
	ErrNotAuth     = errors.New("The client is not authenticated")
)

func decodeErr(body io.ReadCloser) *response.Error {
	e := &response.Error{}
	_ = json.NewDecoder(body).Decode(e)
	return e
}
