package experiment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoWork(t *testing.T) {
	tests := []struct {
		name           string
		iterations     int
		expectedResult uint64
	}{
		{
			name:           "one million",
			iterations:     1_000_000,
			expectedResult: getExpectedResult(1_000_000),
		},
		{
			name:           "ten million",
			iterations:     10_000_000,
			expectedResult: getExpectedResult(10_000_000),
		},
		{
			name:           "one hundred million",
			iterations:     100_000_000,
			expectedResult: getExpectedResult(100_000_000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workload := Workload{
				Iterations: tt.iterations,
			}

			start := time.Now()
			result := workload.DoWork()
			duration := time.Since(start)

			t.Logf("DoWork() took %v", duration)

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func getExpectedResult(iterations int) uint64 {
	n := uint64(iterations)

	return n * (n + 1) * (2*n + 1) / 6
}
