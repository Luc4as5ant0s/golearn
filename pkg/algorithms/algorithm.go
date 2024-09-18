package algorithms

type Algorithm interface {
	Initialize(params map[string]interface{}) error
	Compute(data []float64) ([]float64, error)
	Update(params map[string]interface{}) error
}