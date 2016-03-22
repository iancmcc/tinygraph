package tinygraph_test

import (
	"fmt"

	. "github.com/iancmcc/tinygraph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array Matrix", func() {

	var (
		m           *ArrayMatrix
		size        uint32
		cellsize    MatrixType
		wordsPerRow uint32
	)

	AssertMatrixChecks := func() {

		BeforeEach(func() {
			m = NewArrayMatrix(cellsize, size).(*ArrayMatrix)
			wordsPerRow = ((size * (1 << cellsize)) + 63) / 64
		})

		It(fmt.Sprintf("should use %d words per row", wordsPerRow), func() {
			Ω(m.WordsPerRow).Should(BeEquivalentTo(wordsPerRow))
		})

		It("should refuse to set a cell in an out-of-bounds column", func() {
			err := m.Set(size-1, size)
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(ErrOutOfBounds))
		})

		It("should refuse to get a cell in an out-of-bounds column", func() {
			_, err := m.Get(size-1, size)
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(ErrOutOfBounds))
		})

		It("should refuse to set a cell in an out-of-bounds row", func() {
			err := m.Set(size, size-1)
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(ErrOutOfBounds))
		})

		It("should refuse to get a cell in an out-of-bounds row", func() {
			_, err := m.Get(size, size-1)
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(ErrOutOfBounds))
		})

		It("should not refuse to set in-bounds cells", func() {
			err := m.Set(0, 0)
			Ω(err).ShouldNot(HaveOccurred())

			err = m.Set(size-1, size-1)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("should return the correct word index", func() {
			m.Set(0, 0)
			Ω(m.Get(0, 0)).Should(BeEquivalentTo(1))
			//[1 0 0 0]

			m.Set(size-1, size-1)
			//[1 0 2 0]
			Ω(m.Get(size-1, size-1)).Should(BeEquivalentTo(1))
		})

		It("should be able to be transposed", func() {
			m.Set(0, 1)
			Ω(m.Transpose().Get(1, 0)).Should(BeEquivalentTo(1))
			Ω(m.Transpose().Get(0, 1)).Should(BeEquivalentTo(0))

			m.Transpose().Set(0, 1)
			Ω(m.Get(1, 0)).Should(BeEquivalentTo(1))

			Ω(m.Transpose().Transpose()).Should(Equal(m))
		})

		It("should set extra bits", func() {
			m.Set(0, 1)
			m.SetBit(0, 1, 1)
			Ω(m.Get(0, 1)).Should(BeEquivalentTo(3))
		})
	}

	AssertForAllMatrixTypes := func() {
		Context("with a single-bit cell size", func() {
			BeforeEach(func() {
				cellsize = Bit
			})
			AssertMatrixChecks()
		})

		Context("with a two-bit cell size", func() {
			BeforeEach(func() {
				cellsize = TwoBit
			})
			AssertMatrixChecks()
		})

		Context("with a four-bit cell size", func() {
			BeforeEach(func() {
				cellsize = FourBit
			})
			AssertMatrixChecks()
		})

		Context("with an eight-bit cell size", func() {
			BeforeEach(func() {
				cellsize = Byte
			})
			AssertMatrixChecks()
		})

		Context("with a sixteen-bit cell size", func() {
			BeforeEach(func() {
				cellsize = SixteenBit
			})
			AssertMatrixChecks()
		})

		Context("with a thirty-two-bit cell size", func() {
			BeforeEach(func() {
				cellsize = ThirtyTwoBit
			})
			AssertMatrixChecks()
		})

		Context("with a sixty-four-bit cell size", func() {
			BeforeEach(func() {
				cellsize = Long
			})
			AssertMatrixChecks()
		})

	}

	Context("with a 2x2 matrix", func() {
		BeforeEach(func() {
			size = 2
		})
		AssertForAllMatrixTypes()
	})

	Context("with an 8x8 matrix", func() {
		BeforeEach(func() {
			size = 8
		})
		AssertForAllMatrixTypes()
	})

	/*
		Context("with a 100x100 matrix", func() {
			BeforeEach(func() {
				size = 100
			})
			AssertForAllMatrixTypes()
		})

		Context("with a 1024x1024 matrix", func() {
			BeforeEach(func() {
				size = 1024
			})
			AssertForAllMatrixTypes()
		})
	*/

})
