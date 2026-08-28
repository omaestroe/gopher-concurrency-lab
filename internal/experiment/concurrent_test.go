package experiment

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/omaestroe/gopher-concurrency-lab/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRunConcurrent(t *testing.T) {
	tests := []struct {
		name           string
		iterations     int
		jobs           []domain.Job
		expectedResult uint64
	}{
		{
			name:           "1M - 5 jobs",
			iterations:     1_000_000,
			jobs:           createJobs(5),
			expectedResult: getExpectedResult(1_000_000),
		},
		{
			name:           "10M - 5 jobs",
			iterations:     10_000_000,
			jobs:           createJobs(5),
			expectedResult: getExpectedResult(10_000_000),
		},
		{
			name:           "100M - 5 jobs",
			iterations:     100_000_000,
			jobs:           createJobs(5),
			expectedResult: getExpectedResult(100_000_000),
		},
		{
			name:           "100M - 7 jobs",
			iterations:     100_000_000,
			jobs:           createJobs(7),
			expectedResult: getExpectedResult(100_000_000),
		},
		{
			name:           "100M - 12 jobs",
			iterations:     100_000_000,
			jobs:           createJobs(12),
			expectedResult: getExpectedResult(100_000_000),
		},
		{
			name:           "100M - 16 jobs",
			iterations:     100_000_000,
			jobs:           createJobs(16),
			expectedResult: getExpectedResult(100_000_000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workload := Workload{
				Iterations: tt.iterations,
			}
			concurrent := Concurrent{
				Jobs:     tt.jobs,
				Workload: workload,
			}

			start := time.Now()
			results := concurrent.Run()
			duration := time.Since(start)

			t.Logf("Run %d concurrent jobs took %v", len(tt.jobs), duration)
			assert.Len(t, results, len(tt.jobs))
			for i, result := range results {
				assert.Equal(t, tt.jobs[i].ID, result.JobID)
				assert.Equal(t, tt.expectedResult, result.Result)
			}
		})
	}
}

func BenchmarkRunConcurrent(b *testing.B) {
	tests := []struct {
		name       string
		iterations int
		jobs       []domain.Job
	}{
		{
			name:       "100M - 1 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(1),
		},
		{
			name:       "100M - 2 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(2),
		},
		{
			name:       "100M - 4 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(4),
		},
		{
			name:       "100M - 8 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(8),
		},
		{
			name:       "100M - 16 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(16),
		},
		{
			name:       "100M - 32 jobs",
			iterations: 100_000_000,
			jobs:       createJobs(32),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			workload := Workload{
				Iterations: tt.iterations,
			}
			concurrent := Concurrent{
				Jobs:     tt.jobs,
				Workload: workload,
			}

			for b.Loop() {
				concurrent.Run()
			}
		})
	}
}

func BenchmarkRunConcurrentGOMAXPROCS(b *testing.B) {
	tests := []struct {
		name     string
		maxProcs int
	}{
		{
			name:     "GOMAXPROCS_1",
			maxProcs: 1,
		},
		{
			name:     "GOMAXPROCS_2",
			maxProcs: 2,
		},
		{
			name:     "GOMAXPROCS_4",
			maxProcs: 4,
		},
		{
			name:     "GOMAXPROCS_8",
			maxProcs: 8,
		},
	}

	jobs := createJobs(8)

	workload := Workload{
		Iterations: 100_000_000,
	}

	concurrent := Concurrent{
		Jobs:     jobs,
		Workload: workload,
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			previous := runtime.GOMAXPROCS(tt.maxProcs)
			defer runtime.GOMAXPROCS(previous)

			for b.Loop() {
				concurrent.Run()
			}
		})
	}
}

// Helper functions
func createJobs(count int) []domain.Job {
	jobs := make([]domain.Job, count)

	for i := range count {
		jobs[i] = domain.Job{
			ID: fmt.Sprintf("job%d", i+1),
		}
	}

	return jobs
}

/*func TestRuntimeInfo(t *testing.T) {
	t.Logf("NumCPU: %d", runtime.NumCPU())
	t.Logf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))
}*/
