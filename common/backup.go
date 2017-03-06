package common

type BackupState string

const (
	Submitted       BackupState = "SUBMITTED"
	Inprogress      BackupState = "INPROGRESS"
	Completed       BackupState = "COMPLETED"
	Failed          BackupState = "FAILED"
	Canceling       BackupState = "CANCELLING"
	Canceled        BackupState = "CANCELED"
	Timeout         BackupState = "TIMEOUT"
	DeleteSubmitted BackupState = "DELETE_SUBMITTED"
	Deleting        BackupState = "DELETING"
	Deleted         BackupState = "DELETED"
)
