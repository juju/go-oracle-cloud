package common

type InstanceState string

const (
	StateQueued       InstanceState = "queued"
	StatePreparing    InstanceState = "preparing"
	StateInitializing InstanceState = "initializing"
	StateStarting     InstanceState = "starting"
	StateRunning      InstanceState = "running"
	StateSuspending   InstanceState = "suspending"
	StateSuspended    InstanceState = "suspended"
	StateStopping     InstanceState = "stopping"
	StateStopped      InstanceState = "stopped"
	StateUnreachable  InstanceState = "unreachable"
	StateError        InstanceState = "error"
)
