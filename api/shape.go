package api

import (
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

func (c Client) ShapeDetails(name Shape) (resp response.Shape, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	s, ok := shapes[name]
	if !ok {
		return resp, ErrUndefinedShape
	}

	url := fmt.Sprintf("%s/%s/%s", c.endpoint, "shape", s)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) AllShapeDetails() (resp response.AllShapes, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/%s/", c.endpoint, "shape")

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}
	return resp, nil
}
