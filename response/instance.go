package response

type LaunchPlan struct {
}

type Instance struct {
}

type AllInstance struct {
	List []List `json:"list"`
}

type AllInstanceNames struct {
	Result []string `json:"result"`
}

type List struct {
	Domain                 string    `json:"domain"`
	Placement_requirements []string  `json:"placement_requirements"`
	Ip                     string    `json:"ip"`
	Site                   string    `json:"site,omitempty"`
	Shape                  string    `json:"shape"`
	Imagelist              string    `json:"imagelist"`
	Atributes              Atributes `json:"attributes"`
}

type Atributes struct {
	Network Network `json:"network"`
}

type Network struct {
}
