// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package response

type ImageList struct {
	Default     uint64    `json:"default"`
	Description string    `json:"description,omitempty"`
	Entries     []Entries `json:"entries,omitempty"`
	Uri         string    `json:"uri"`
	Name        string    `json:"name"`
}

type Entries struct {
	Attributes    attr     `json:"attributes,omitempty"`
	Version       uint64   `json:"version"`
	MachineImages []string `json:"machineimages"`
	Uri           string   `json:"uri"`
}

type attr struct {
	//TODO(sgiulitti) make a special type for userdata
	Userdata        map[string]interface{} `json:"userdata,omitempty"`
	MinimumDiskSize string                 `json:"minimumdisksize,omitempty"`
	DefaultShape    string                 `json:"defaultshape,omitempty"`
	SupportedShapes string                 `json:"supportedShapes,omitempty"`
}

type AllImageList struct {
	Result []ImageList `json:"result"`
}

type ImageListEntry struct {
	Attributes    attr      `json:"attributes,omitempty"`
	ImageList     ImageList `json:"imagelist"`
	Version       uint64    `json:"version"`
	Machineimages []string  `json:"machineimages"`
	Uri           string    `json:"uri"`
}
