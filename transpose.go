package tinygraph

type Transposed struct {
	Matrix Matrix
}

// Verify Matric interface implementation
var _ Matrix = &Transposed{}

func (t *Transposed) Set(i, j uint64) error {
	return t.Matrix.Set(j, i)
}

func (t *Transposed) Unset(i, j uint64) error {
	return t.Matrix.Unset(j, i)
}

func (t *Transposed) SetBit(i, j, k uint64) error {
	return t.Matrix.SetBit(j, i, k)
}

func (t *Transposed) UnsetBit(i, j, k uint64) error {
	return t.Matrix.UnsetBit(j, i, k)
}

func (t *Transposed) Get(i, j uint64) (uint64, error) {
	return t.Matrix.Get(j, i)
}

func (t *Transposed) Transpose() Matrix {
	return t.Matrix
}
