package response

import "github.com/hoenirvili/go-oracle-cloud/common"

type SecRule struct {
	Action      common.SecRule `json:"action"`
	Application string         `json:"application"`
	Description string         `json:"description,omitempty"`
	Disabled    bool           `json:"disabled"`
	Dst_list    string         `json:"dst_list"`
	Name        string         `json:"name"`
	src_list    string         `json:"src_list"`
	Uri         string         `json:"uri"`
}

type AllSecRules struct {
	Result []SecRule `json:"result,omitempty"`
}
