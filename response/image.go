package response

type Image struct {
	Default     uint64    `json:"default"`
	Description string    `json:"description,omitempty"`
	Entries     []Entries `json:"entries",omitempty`
	Uri         string    `json:"uri"`
	Name        string    `json:"name"`
}

type AllImageNames struct {
	Result []string `json:"result"`
}

type Entries struct {
	Attributes    attr     `json:"attributes,omitempty"`
	Version       uint64   `json:"version"`
	MachineImages []string `json:"machineimages"`
	Uri           string   `json:"uri"`
}

type attr struct {
	//TODO (make a special type for userdata)
	Userdata        interface{} `json:"userdata,omitempty"`
	MinimumDiskSize string      `json:"minimumdisksize,omitempty"`
	DefaultShape    string      `json:"defaultshape,omitempty"`
	SupportedShapes string      `json:"supportedshape,omitempty"`
}

type AllImage struct {
	Result []Image `json:"result"`
}
