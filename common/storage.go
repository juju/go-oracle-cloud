package common

import "errors"

// Index is an index storage number
// that is in the range of 1-10
// this index determines the device name by which the volume
// is exposed to the instance
type Index int

// Validate validates the index provided if it's
// compliant with the oracle cloud storage index system
func (i Index) Validate() (err error) {
	if i < 1 || i > 10 {
		return errors.New(
			"go-oracle-cloud: Invalid storage index number",
		)
	}
	return nil
}

// StateStorage type that holds the state of the storage
// in motion
type StateStorage string

const (
	// StateAttaching describes the storage attachment is in
	// the process of attaching to the instance.
	StateAttaching StateStorage = "attaching"

	// StateAttached describes the storage attachment is
	// attached to the instance.
	StateAttached StateStorage = "attached"

	// StateDetached describes the storage attachment is
	// in the process of detaching from the instance.
	StateDetached StateStorage = "detached"

	// StateUnavailable tells that the storage attachment is unavailable.
	StateUnavailable StateStorage = "unabailable"

	//StateUnknown descibes the state of the storage attachment is not known.
	StateUnknown StateStorage = "unknown"
)
