package response

type Shape struct {
	Nds_iops_limit uint64 `json:"nds_iops_limit"`
	Ram            uint64 `json:"ram"`
	Uri            string `json:"uri"`
	Root_disk_size uint64 `json:"root_disk_size"`
	Io             uint64 `json:"io"`
	Name           string `json:"name"`
}

type AllShapes struct {
	Result []Shape `json:"result"`
}
