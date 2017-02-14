package response

type Acl struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	EnableFlag  bool     `json:"enableFlag"`
	Tags        []string `json:"tags,omitempty"`
	Uri         string   `json:"uri"`
}

type AllAcl struct {
	Result []Acl `json:"result"`
}
