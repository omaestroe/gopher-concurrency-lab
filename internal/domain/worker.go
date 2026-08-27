package domain

//🐹 Worker

type WorkerStatus string

const (
	WorkerIdle       WorkerStatus = "IDLE"
	WorkerProcessing WorkerStatus = "PROCESSING"
	WorkerWaiting    WorkerStatus = "WAITING"
	WorkerStopped    WorkerStatus = "STOPPED"
)

type Worker struct {
	ID     string
	Status WorkerStatus
	JobID  *string
}
