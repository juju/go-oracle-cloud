package common

import "errors"

// VcableID is the vcable it of the instance that
// is associated with the ip reservation.
type VcableID string

// Validate checks if the VcableID provided is empty or not
func (v VcableID) Validate() (err error) {
	if v == "" {
		return errors.New("go-oracle-cloud: Empty vcable id")
	}

	return nil
}
