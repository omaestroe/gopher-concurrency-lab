package experiment

import (
	"testing"
	"time"

	"github.com/omaestroe/gopher-concurrency-lab/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRunSequential(t *testing.T) {
	tests := []struct {
		name           string
		iterations     int
		jobs           []domain.Job
		expectedResult uint64
	}{
		{
			name:       "one million five jobs",
			iterations: 1_000_000,
			jobs: []domain.Job{
				{ID: "job1"},
				{ID: "job2"},
				{ID: "job3"},
				{ID: "job4"},
				{ID: "job5"},
			},
			expectedResult: getExpectedResult(1_000_000),
		},
		{
			name:       "ten million five jobs",
			iterations: 10_000_000,
			jobs: []domain.Job{
				{ID: "job1"},
				{ID: "job2"},
				{ID: "job3"},
				{ID: "job4"},
				{ID: "job5"},
			},
			expectedResult: getExpectedResult(10_000_000),
		},
		{
			name:       "one hundred million five jobs",
			iterations: 100_000_000,
			jobs: []domain.Job{
				{ID: "job1"},
				{ID: "job2"},
				{ID: "job3"},
				{ID: "job4"},
				{ID: "job5"},
			},
			expectedResult: getExpectedResult(100_000_000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workload := Workload{
				Iterations: tt.iterations,
			}
			sequential := Sequential{
				Jobs:     tt.jobs,
				Workload: workload,
			}

			start := time.Now()
			results := sequential.Run()
			duration := time.Since(start)

			t.Logf("Run %d sequential jobs took %v", len(tt.jobs), duration)
			assert.Len(t, results, len(tt.jobs))
			for i, result := range results {
				assert.Equal(t, tt.jobs[i].ID, result.JobID)
				assert.Equal(t, tt.expectedResult, result.Result)
			}
		})
	}
}
