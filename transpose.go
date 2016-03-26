package tinygraph

type TransposedArrayMatrix struct {
	Matrix *ArrayMatrix
}

// Verify Matric interface implementation
var _ Matrix = &TransposedArrayMatrix{}

func (t *TransposedArrayMatrix) Set(i, j uint64) error {
	return t.Matrix.Set(j, i)
}

func (t *TransposedArrayMatrix) Unset(i, j uint64) error {
	return t.Matrix.Unset(j, i)
}

func (t *TransposedArrayMatrix) SetBit(i, j, k uint64) error {
	return t.Matrix.SetBit(j, i, k)
}

func (t *TransposedArrayMatrix) UnsetBit(i, j, k uint64) error {
	return t.Matrix.UnsetBit(j, i, k)
}

func (t *TransposedArrayMatrix) Get(i, j uint64) (uint64, error) {
	return t.Matrix.Get(j, i)
}

func (t *TransposedArrayMatrix) Replace(i, j, k uint64) error {
	return t.Matrix.Replace(j, i, k)
}

func (t *TransposedArrayMatrix) Clear(i, j uint64) error {
	return t.Matrix.Clear(j, i)
}

func (t *TransposedArrayMatrix) Transpose() Matrix {
	return t.Matrix
}

func (t *TransposedArrayMatrix) Copy() Matrix {
	n := t.Matrix.Copy().(*ArrayMatrix)
	return n.Transpose()
}
