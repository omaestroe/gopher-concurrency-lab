package experiment

import (
	"sync"

	"github.com/omaestroe/gopher-concurrency-lab/internal/domain"
)

type Concurrent struct {
	Jobs     []domain.Job
	Workload Workload
}

func (c Concurrent) Run() []JobResult {
	results := make([]JobResult, len(c.Jobs))
	var wg sync.WaitGroup
	for i, job := range c.Jobs {
		wg.Add(1)
		go func(i int, job domain.Job) {
			defer wg.Done()
			result := c.Workload.DoWork()
			results[i] = JobResult{
				JobID:  job.ID,
				Result: result,
			}
		}(i, job)
	}
	wg.Wait()
	return results
}
