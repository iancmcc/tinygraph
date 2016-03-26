package tinygraph

import (
	"errors"
	"fmt"
	"strconv"
)

// MatrixType is log 2 of the cell size
type MatrixType uint64

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
	WordSize = uint64(64)
	// WordSizeMinusOne is word size minus one, for bitwise modulus
	WordSizeMinusOne = uint64(63)
	// WordSizeExp is log 2 of the word size
	WordSizeExp = uint64(6)
	// One is 1
	One = uint64(1)
)

var (
	// ErrOutOfBounds is returned when a coordinate outside the bounds of the
	// matrix is requested or set
	ErrOutOfBounds        = errors.New("Bit requested is outside the matrix bounds")
	_              Matrix = &ArrayMatrix{}
)

// Matrix is a 2-dimensional square matrix.
type Matrix interface {
	Set(i, j uint64) error
	Unset(i, j uint64) error
	SetBit(i, j, k uint64) error
	UnsetBit(i, j, k uint64) error
	Replace(i, j, k uint64) error
	Clear(i, j uint64) error
	Get(i, j uint64) (uint64, error)
	Copy() Matrix
	Transpose() Matrix
}

func New(mtype MatrixType, size uint64) Matrix {
	return NewArrayMatrix(mtype, size)
}

// ArrayMatrix is an implementation of Matrix that stores cells as
// a 1-dimensional array of uint64s
type ArrayMatrix struct {
	Words       []uint64
	Size        uint64
	LastIndex   uint64
	WordsPerRow uint64
	MType       MatrixType
	cellmask    uint64
	cellsize    uint64
}

// NewArrayMatrix creates a new matrix with a given cell size and given dimensions
func NewArrayMatrix(mtype MatrixType, size uint64) Matrix {
	matrix := &ArrayMatrix{
		Size:      size,
		LastIndex: size - 1,
		MType:     mtype,
		cellmask:  (1 << (1 << mtype)) - 1,
		cellsize:  1 << mtype,
	}
	if mtype == Long {
		matrix.cellmask = ^uint64(0)
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
func (m *ArrayMatrix) GetWordIndex(i, j uint64) uint64 {
	return (i * m.WordsPerRow) + (j << m.MType >> WordSizeExp)
}

// Set sets the principal bit of the cell at the coordinates requested
func (m *ArrayMatrix) Set(i, j uint64) error {
	if i > m.LastIndex || j > m.LastIndex {
		return ErrOutOfBounds
	}
	mask := One << (j << m.MType & WordSizeMinusOne)
	m.Words[m.GetWordIndex(i, j)] |= mask
	return nil
}

func (m *ArrayMatrix) Unset(i, j uint64) error {
	if i > m.LastIndex || j > m.LastIndex {
		return ErrOutOfBounds
	}
	mask := One << (j << m.MType & WordSizeMinusOne)
	m.Words[m.GetWordIndex(i, j)] &= ^mask
	return nil
}

func (m *ArrayMatrix) SetBit(i, j, k uint64) error {
	if i > m.LastIndex || j > m.LastIndex || k >= m.cellsize {
		return ErrOutOfBounds
	}
	mask := One << k << (j << m.MType & WordSizeMinusOne)
	m.Words[m.GetWordIndex(i, j)] |= mask
	return nil
}

func (m *ArrayMatrix) Replace(i, j, k uint64) error {
	if i > m.LastIndex || j > m.LastIndex {
		return ErrOutOfBounds
	}
	offset := (j << m.MType & WordSizeMinusOne)
	newcell := k << offset
	mask := m.cellmask << offset
	word := m.Words[m.GetWordIndex(i, j)]
	m.Words[m.GetWordIndex(i, j)] ^= (word ^ newcell) & mask
	return nil
}

func (m *ArrayMatrix) Clear(i, j uint64) error {
	return m.Replace(i, j, 0)
}

func (m *ArrayMatrix) UnsetBit(i, j, k uint64) error {
	if i > m.LastIndex || j > m.LastIndex || k >= m.cellsize {
		return ErrOutOfBounds
	}
	offset := (j << m.MType & WordSizeMinusOne)
	mask := One << k << offset
	cell := (m.cellmask << offset) & ^mask
	m.Words[m.GetWordIndex(i, j)] &= cell
	return nil
}

// Get gets the cell at the coordinates requested
func (m *ArrayMatrix) Get(i, j uint64) (uint64, error) {
	if i > m.LastIndex || j > m.LastIndex {
		return 0, ErrOutOfBounds
	}
	word := m.Words[m.GetWordIndex(i, j)]
	result := word >> (j << m.MType & WordSizeMinusOne) & m.cellmask
	return result, nil
}

func (m *ArrayMatrix) GetRow(i uint64) ([]uint64, error) {
	if i > m.LastIndex {
		return nil, ErrOutOfBounds
	}
	row := make([]uint64, m.WordsPerRow)
	idx := i * m.WordsPerRow
	copy(row, m.Words[idx:idx+m.WordsPerRow])
	return row, nil
}

// Transpose returns a view of the matrix with the axes transposed
func (m *ArrayMatrix) Transpose() Matrix {
	return &TransposedArrayMatrix{m}
}

func (m *ArrayMatrix) Copy() Matrix {
	n := &ArrayMatrix{
		Size:      m.Size,
		LastIndex: m.LastIndex,
		MType:     m.MType,
		cellmask:  m.cellmask,
		cellsize:  m.cellsize,
	}
	n.Words = make([]uint64, len(m.Words))
	copy(n.Words, m.Words)
	return n
}

func logWord(s string, i uint64) {
	fmt.Printf("%s: %s\n", s, strconv.FormatUint(i, 2))
}
