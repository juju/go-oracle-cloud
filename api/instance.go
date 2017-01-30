package api

type shape uint8

const (
	OC1 shape = iota
	OC2
	OC3
	OC4
	OC5
	OC6
	OC7
)

// shapes a map of shape mappings
var shapes = map[shape]string{
	OC1: "oc1",
	OC2: "oc2",
	OC3: "oc3",
	OC4: "oc4",
	OC5: "oc5",
	OC6: "oc6",
	OC7: "oc7",
}

type InstanceParams struct {
	// every instance in oracle cloud has a predefined shape
	// in order to create a virtual instance we need to specify the
	// computing power shape
	Shape shape
	// the virtual image that will be used
	// in order to init the instance
	Imagelist string
	// name of the instance
	Name  string
	Label string
	// names of the ssh keys
	SSHKeys []string
}

// TODO
func (c Client) CreateInstance(paramas InstanceParams) (resp resp.Instance, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	shape, ok := shapes[params.Shape]
	if !ok {
		return resp, ErrUndefinedShape
	}

	instanceArgs := struct {
		Shape     string   `json:"shape"`
		Imagelist string   `json:"imagelist"`
		Name      string   `json:"name"`
		SSHKeys   []string `json:"sshkeys"`
	}{
		Shape: shap,
	}

	return resp, nil
}
