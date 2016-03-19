package tinygraph

type Transposed struct {
	Matrix Matrix
}

// Verify Matric interface implementation
var _ Matrix = &Transposed{}

func (t *Transposed) Set(i, j uint32) error {
	return t.Matrix.Set(j, i)
}

func (t *Transposed) Get(i, j uint32) uint8 {
	return t.Matrix.Get(j, i)
}

func (t *Transposed) Transpose() Matrix {
	return t.Matrix
}
