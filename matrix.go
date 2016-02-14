package tinygraph

type MatrixType uint64

const (
	Bit MatrixType = 1 << iota
	HalfNybble
	Nybble
	Byte
	Halfword
	Word
	Long
)

type Matrix interface {
	Set(i uint64)
	Get(i uint64) uint64
}

func NewMatrix(size uint64, cellsize MatrixType) Matrix {
	switch cellsize {
	case Bit:
		return &BitMatrix{}
	}
	return nil
}

type BitMatrix struct {
}

func (m *BitMatrix) Set(i uint64) {
}

func (m *BitMatrix) Get(i uint64) uint64 {
	return 0
}
