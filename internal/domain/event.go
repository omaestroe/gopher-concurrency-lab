package domain

import "time"

type EventType string

const (
	ExperimentStarted  EventType = "EXPERIMENT_STARTED"
	ExperimentFinished EventType = "EXPERIMENT_FINISHED"
)

type Event struct {
	ID        string
	Type      EventType
	Timestamp time.Time
	JobID     *string
	WorkerID  *string
}
