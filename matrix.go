package tinygraph

import "errors"

// MatrixType is log 2 of the cell size
type MatrixType uint8

const (
	// Bit is a single-bit cell
	Bit MatrixType = iota
	// TwoBit is a two-bit cell
	TwoBit
	// FourBit is a four-bit cell
	FourBit
	// Byte is an eight-bit cell
	Byte
	// SixteenBit is a sixteen-bit cell
	SixteenBit
	// ThirtyTwoBit is a thirty-two-bit cell
	ThirtyTwoBit
	// Long is a sixty-four-bit cell
	Long
)

const (
	// WordSize is the size of the word we will be using to store the matrix
	WordSize = uint8(64)
	// WordSizeExp is log 2 of the word size
	WordSizeExp = uint8(6)
)

var (
	// ErrOutOfBounds is returned when a coordinate outside the bounds of the
	// matrix is requested or set
	ErrOutOfBounds        = errors.New("Bit requested is outside the matrix bounds")
	_              Matrix = &ArrayMatrix{}
)

// Matrix is a 2-dimensional square matrix.
type Matrix interface {
	Set(i, j uint32) error
	Get(i, j uint32) (uint8, error)
	Transpose() Matrix
}

// ArrayMatrix is an implementation of Matrix that stores cells as
// a 1-dimensional array of uint64s
type ArrayMatrix struct {
	Words       []uint64
	Size        uint32
	LastIndex   uint32
	WordsPerRow uint32
	MType       MatrixType
}

// NewArrayMatrix creates a new matrix with a given cell size and given dimensions
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

// GetWordIndex returns the index of the word that contains the coordinate specified
func (m *ArrayMatrix) GetWordIndex(i, j uint32) uint32 {
	return (j * m.WordsPerRow) + (i >> 6)
}

// Set sets the principal bit of the cell at the coordinates requested
func (m *ArrayMatrix) Set(i, j uint32) error {
	if i > m.LastIndex || j > m.LastIndex {
		return ErrOutOfBounds
	}
	//idx := m.GetWordIndex(i)
	offset := j*m.WordsPerRow + (i / 64)
	m.Words[offset] |= 1 << (i & 0x1f)
	return nil
}

// Get gets the principal bit of the cell at the coordinates requested
func (m *ArrayMatrix) Get(i, j uint32) (uint8, error) {
	if i > m.LastIndex || j > m.LastIndex {
		return 0, ErrOutOfBounds
	}
	offset := m.GetWordIndex(i, j)
	return uint8((m.Words[offset] >> (i & 0x1f)) & 1), nil
}

// Transpose returns a view of the matrix with the axes transposed
func (m *ArrayMatrix) Transpose() Matrix {
	return &Transposed{m}
}
