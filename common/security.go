package common

import "errors"

type SecRule string

func (s SecRule) Validate() (err error) {
	if s == "" {
		return errors.New("go-oracle-cloud: Empty secure rule permission")
	}

	return nil
}

const (
	DefaultSecRule SecRule = "PERMIT"
)
