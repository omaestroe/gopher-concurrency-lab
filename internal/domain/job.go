package domain

import "time"

type JobStatus string

const (
	JobCreated    JobStatus = "CREATED"
	JobWaiting    JobStatus = "WAITING"
	JobProcessing JobStatus = "PROCESSING"
	JobCompleted  JobStatus = "COMPLETED"
	JobFailed     JobStatus = "FAILED"
	JobCancelled  JobStatus = "CANCELLED"
)

type Job struct {
	ID          string
	Status      JobStatus
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
}
