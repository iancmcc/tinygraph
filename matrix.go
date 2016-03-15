package tinygraph

import "errors"

// MatrixType is log 2 of the datum size
type MatrixType uint8

const (
	Bit MatrixType = iota
	TwoBit
	FourBit
	Byte
	SixteenBit
	ThirtyTwoBit
	Long
)

const (
	WordSize    = uint8(64)
	WordSizeExp = uint8(6)
)

var ErrOutOfBounds = errors.New("Bit requested is outside the matrix bounds")

type Matrix interface {
	Set(i, j uint) error
	Get(i, j uint) uint8
	Transpose() Matrix
}

type ArrayMatrix struct {
	words []uint64
	size  uint64
	mtype MatrixType
}

func NewArrayMatrix(mtype MatrixType, size uint64) Matrix {
	numwords := (size << mtype >> WordSizeExp) + 1
	return &ArrayMatrix{
		words: make([]uint64, numwords),
		size:  size,
		mtype: mtype,
	}
}

func (m *ArrayMatrix) Set(i, j uint) error {
	return nil
}

var _ Matrix = &ArrayMatrix{}

func (m *ArrayMatrix) Get(i, j uint) uint8 {
	var result uint8
	return result
}

func (m *ArrayMatrix) Transpose() Matrix {
	return &Transposed{m}
}
