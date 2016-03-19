package tinygraph_test

import (
	. "github.com/iancmcc/tinygraph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array Matrix", func() {

	var (
		m        *ArrayMatrix
		size     uint32
		cellsize MatrixType
	)
	Context("with a single-bit cell size", func() {

		cellsize = Bit

		Context("with an 8x8 matrix", func() {

			size = 8

			BeforeEach(func() {
				m = NewArrayMatrix(cellsize, size).(*ArrayMatrix)
			})

			It("should use 1 word per row", func() {
				Ω(m.WordsPerRow).Should(BeEquivalentTo(1))
			})

			It("should refuse to set a cell in an out-of-bounds column", func() {
				err := m.Set(7, 8)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(Equal(ErrOutOfBounds))
			})

			It("should refuse to set a cell in an out-of-bounds row", func() {
				err := m.Set(8, 7)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(Equal(ErrOutOfBounds))
			})

			It("should not refuse to set in-bounds cells", func() {
				err := m.Set(0, 0)
				Ω(err).ShouldNot(HaveOccurred())

				err = m.Set(7, 7)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
