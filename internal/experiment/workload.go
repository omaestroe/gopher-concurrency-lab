package experiment

type Workload struct {
	Iterations int
}

func (workload Workload) DoWork() uint64 {
	var result uint64
	for i := 1; i <= workload.Iterations; i++ {
		result += uint64(i) * uint64(i)
	}
	return result
}
