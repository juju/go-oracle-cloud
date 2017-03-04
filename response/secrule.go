package response

import "github.com/hoenirvili/go-oracle-cloud/common"

type SecRule struct {
	Action      common.SecRuleAction `json:"action"`
	Application string               `json:"application"`
	Description string               `json:"description,omitempty"`
	Disabled    bool                 `json:"disabled"`
	Dst_is_ip   string               `json:"dst_is_ip"`
	Dst_list    string               `json:"dst_list"`
	Name        string               `json:"name"`
	Id          string               `json:"id"`
	Src_is_ip   string               `json:"src_is_ip"`
	Src_list    string               `json:"src_list"`
	Uri         string               `json:"uri"`
}

type AllSecRules struct {
	Result []SecRule `json:"result,omitempty"`
}
