package response

type VnicSet struct {
	// AppliedAcls is a list of ACLs applied to the VNICs in the set.
	AppliedAcls []string `json:"appliedAcls,omitempty"`

	// Description of the object
	Description string `json:"description,omitempty"`

	// Name is the name of the vnic set
	Name string `json:"name"`

	// Tags associated with the object.
	Tags []string `json:"tags"`

	// List of VNICs associated with this VNIC set
	Vnics []string `json:"vnics"`

	// Uri is the Uniform Resource Identifier
	Uri string `json:"uri"`
}

type AllVnicSets struct {
	Result []VnicSet `json:"result,omitempty"`
}
