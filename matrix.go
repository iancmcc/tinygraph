package tinygraph

import (
	"errors"

	"github.com/willf/bitset"
)

type MatrixType uint64

const (
	Bit MatrixType = 1 << iota
	HalfNybble
	Nybble
	Byte
	HalfWord
	Word
	Long
)

var ErrOutOfBounds = errors.New("Bit requested is outside the matrix bounds")

type Matrix interface {
	initialize()
	Set(x, y uint) error
	Unset(x, y uint)
	Get(x, y uint) uint64
}

func NewMatrix(size uint, cellsize MatrixType) (matrix Matrix) {
	switch cellsize {
	case Bit:
		matrix = &BitMatrix{size: size}
	}
	matrix.initialize()
	return
}

type BitMatrix struct {
	size uint
	rows []*bitset.BitSet
}

func (m *BitMatrix) initialize() {
	var (
		i    uint
		rows []*bitset.BitSet
	)
	rows = make([]*bitset.BitSet, m.size)
	for i = 0; i < m.size; i++ {
		rows[i] = bitset.New(m.size)
	}
}

func (m *BitMatrix) Set(x, y uint) (err error) {
	if x > m.size || y > m.size {
		err = ErrOutOfBounds
	}
	m.rows[x].Set(y)
	return
}

func (m *BitMatrix) Unset(x, y uint) {

}

func (m *BitMatrix) Get(x, y uint) uint64 {
	result := m.rows[x].Get(y)
	return result
}
