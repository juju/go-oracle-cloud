package response

type SSH struct {
	Enabled bool   `json:"enabled"`
	Uri     string `json:"uri"`
	Key     string `json:"key"`
	Name    string `json:"name"`
}

type AllSSH struct {
	Result []SSH `json:"result"`
}

type AllSSHNames struct {
	Result []string `json:"result"`
}
