package domain

type WorkerStatus string

const (
	WorkerIdle       WorkerStatus = "IDLE"
	WorkerProcessing WorkerStatus = "PROCESSING"
	WorkerWaiting    WorkerStatus = "WAITING"
	WorkerStopped    WorkerStatus = "STOPPED"
)

// 🐹 Worker represents a goroutine responsible for processing jobs.
type Worker struct {
	ID     string
	Status WorkerStatus
	JobID  *string
}
