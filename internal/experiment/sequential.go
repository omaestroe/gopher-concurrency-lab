package experiment

import "github.com/omaestroe/gopher-concurrency-lab/internal/domain"

type Sequential struct {
	Jobs     []domain.Job
	Workload Workload
}

func (s Sequential) Run() []JobResult {
	results := make([]JobResult, len(s.Jobs))

	for i, job := range s.Jobs {
		results[i] = JobResult{
			JobID:  job.ID,
			Result: s.Workload.DoWork(),
		}
	}
	return results
}
