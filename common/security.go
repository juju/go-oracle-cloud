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

type FlowDirection string

const (
	Egress  FlowDirection = "egress"
	Ingress FlowDirection = "ingress"
)
