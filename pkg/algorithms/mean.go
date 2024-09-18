package algorithms

type MeanAlgorithm struct{}

func (m *MeanAlgorithm) Initialize(params map[string]interface{}) error {
	return nil
}

func (m *MeanAlgorithm) Compute(data []float64) ([]float64, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(len(data))
	return []float64{mean}, nil
}

func (m *MeanAlgorithm) Update(params map[string]interface{}) error {
	return nil
}
