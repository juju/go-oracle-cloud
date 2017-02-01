package response

type Vnc struct {
	Name        string   `json:"name"`
	Uri         string   `json:"uri"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags"`
	MacAddress  string   `json:"macAddress"`
	TransitFlag bool     `json:"transitFlag"`
}

type AllVnc struct {
	Result []Vnc `json:"result"`
}
