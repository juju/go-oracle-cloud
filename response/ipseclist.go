package response

type IpSecList struct {
	Description  string   `json:"description,omitempty"`
	Name         string   `json:"name"`
	Secipentries []string `json:"secipentries"`
	Uri          string   `json:"uri"`
}

type AllIpSecList struct {
	Result []IpSecList `json:"result,omitempty"`
}
