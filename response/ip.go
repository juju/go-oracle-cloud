package response

// IpAssociation IP address association
// to associate an IP address reservation
type IpAssociation struct {
	Name                 string   `json:"name"`
	Uri                  string   `json:"uri"`
	Description          string   `json:"description,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	Vnic                 string   `json:"vnic"`
	IpAddressReservation string   `json:"ipAddressReservation"`
}

type AllIpAssociation struct {
	Result []IpAssociation `json:"result,omitempty"`
}
