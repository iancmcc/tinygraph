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

var (
	ErrOutOfBounds        = errors.New("Bit requested is outside the matrix bounds")
	_              Matrix = &ArrayMatrix{}
)

type Matrix interface {
	Set(i, j uint32) error
	Get(i, j uint32) uint8
	Transpose() Matrix
}

type ArrayMatrix struct {
	Words       []uint64
	Size        uint32
	LastIndex   uint32
	WordsPerRow uint32
	MType       MatrixType
}

func NewArrayMatrix(mtype MatrixType, size uint32) Matrix {
	matrix := &ArrayMatrix{
		Size:      size,
		LastIndex: size - 1,
		MType:     mtype,
	}
	// calculate the number of words required to store a square matrix of size
	// rows with size mtype cells per row.  Each row begins with a new uint64,
	// to avoid inefficient rebuilding of a row that starts in the middle of
	// a uint64.
	// Shift to get number of bits per row
	bitsPerRow := size << mtype
	// Ceiling division to get number of words per row
	matrix.WordsPerRow = ((bitsPerRow - 1) >> WordSizeExp) + 1
	// Now multiply by size to get total number of words required to store the
	// matrix
	matrix.Words = make([]uint64, size*matrix.WordsPerRow)
	return matrix
}

func (m *ArrayMatrix) GetWord(i uint32) uint32 {
	return uint32(i << m.MType >> WordSizeExp)
}

func (m *ArrayMatrix) Set(i, j uint32) error {
	if i > m.LastIndex || j > m.LastIndex {
		return ErrOutOfBounds
	}
	m.GetWord(i)
	return nil
}

func (m *ArrayMatrix) Get(i, j uint32) uint8 {
	var result uint8
	return result
}

func (m *ArrayMatrix) Transpose() Matrix {
	return &Transposed{m}
}
