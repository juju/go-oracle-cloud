// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package response

// BackupConfiguration you can schedule backups to be taken
// automatically at defined intervals.
// Scheduling a backup creates a snapshot of the specified
// storage volume and the snapshot is stored in
// the associated Oracle Storage Cloud Service instance.
type BackupConfiguration struct {
	// Uri is the Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// RunAsUser represents any actions on this
	// model will be performed as this user.
	RunAsUser string `json:"runAsUser,omitempty"`

	// Name is the name of the backup configuration
	Name string `json:"name,omitempty"`

	// Enabled flag for:
	// when true, backups will automatically
	// be generated based on the interval.
	Enabled bool `json:"enabled,omitempty"`

	// BackupRetentionCount represents how many backups to retain
	BackupRetentionCount uint32 `json:"backupRetentionCount,omitempty"`

	// Scheduled time for next backup execution
	NextScheduledRun string `json:"nextScheduledRun,omitempty"`

	// Interval represents the interval in the backup configuration.
	// There are two kinds of Intervals. Each Interval has its own JSON format.
	// Your Interval field should look like one of the following:
	//
	// "interval":{
	//   "Hourly":{
	//     "hourlyInterval":2
	//	 }
	// }
	//
	//
	// {"DailyWeekly":
	// 	{
	//	  "daysOfWeek":["MONDAY"],
	//	  "timeOfDay":"03:15",
	// 	  "userTimeZone":"America/Los_Angeles"
	//  }
	// }
	// Days of the week is any day of the week
	// fully capitalized (MONDAY, TUESDAY, etc).
	// The user time zone is any IANA user timezone.
	//For example user time zones see List of IANA time zones.
	Interval interface{} `json:"Interval,omitempty"`

	// VolumeUri is the complete URI of the storage
	// volume that you want to backup.
	VolumeUri string `json:"volumeUri,omitempty"`

	// Description of this Backup Configuration
	Description string `json:"description,omitempty"`

	// TagId is the ID used to tag other cloud resources
	TagId string `json:"tagId,omitempty"`
}
