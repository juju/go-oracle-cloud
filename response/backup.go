package response

// TODO(fix backup configuration)
type BackupConfiguration struct {
	Uri                  string   `json:"uri,omitempty"`
	RunAsUser            string   `json:"runAsUser,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Enabled              bool     `json:"enabled,omitempty"`
	BackupRetentionCount uint64   `json:"backupRetentionCount,omitempty"`
	NextScheduledRun     string   `json:"nextScheduledRun,omitempty"`
	Interval             Interval `json:"Interval,omitempty"`
	VolumeUri            string   `json:"volumeUri,omitempty"`
	Description          string   `json:"description,omitempty"`
	TagId                string   `json:"tagId,omitempty"`
}

type Interval struct {
	Hourly string `json:"Hourly"`
}

type Hourly struct {
	HourlyInterval uint64 `json:"hourlyInterval"`
}

type AllBackupConfiguration struct {
	Result []BackupConfiguration `json:"result"`
}
