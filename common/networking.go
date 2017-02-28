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

// IPPool type describing the
// parent pool of an ip association
type IPPool string

const (
	// PublicIPPool standard ip pool for the oracle provider
	PublicIPPool IPPool = "/oracle/public/ippool"
)

func NewIPPool(name IPPool, prefix IPPrefixType) IPPool {
	return IPPool(prefix) + name
}

type IPPrefixType string

const (
	IPReservationType IPPrefixType = "ipreservation:"
	IPPoolType        IPPrefixType = "ippool:"
)
