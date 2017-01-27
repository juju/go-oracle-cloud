package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func newRequest(verb string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("", url, bytes.NewBuffer(body))
	switch verb {
	case "POST", "PUT":
		req.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	case "GET":
	case "DELETE":
	default:
		return nil, errors.New("Invalid http verb provided")
	}

	req.Method = verb
}
