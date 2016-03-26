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
		size        uint64
		cellsize    MatrixType
		wordsPerRow uint64
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

		It("should set the principal bit correctly", func() {
			m.Set(0, 0)
			Ω(m.Get(0, 0)).Should(BeEquivalentTo(1))
			var i, j uint64
			for i = 0; i < size; i += 4 {
				for j = 0; j < size; j += 4 {
					m.Set(i, j)
					Ω(m.Get(i, j)).Should(BeEquivalentTo(1))
				}
			}
		})

		It("should unset the principal bit correctly", func() {
			m.Set(0, 0)
			Ω(m.Get(0, 0)).Should(BeEquivalentTo(1))
			var i, j uint64
			for i = 0; i < size; i += 4 {
				for j = 0; j < size; j += 4 {
					m.Set(i, j)
					m.Unset(i, j)
					Ω(m.Get(i, j)).Should(BeEquivalentTo(0))
				}
			}
		})

		It("should be able to be transposed", func() {
			var i, j uint64
			for i = 0; i < size; i += 5 {
				for j = 0; j < size; j += 5 {
					m.Set(i, j)
					Ω(m.Transpose().Get(j, i)).Should(BeEquivalentTo(1))
				}
			}
		})

		It("should be able to set bits on a transposed matrix", func() {
			var i, j uint64
			for i = 0; i < size; i += 7 {
				for j = 0; j < size; j += 7 {
					m.Transpose().Set(i, j)
					Ω(m.Get(j, i)).Should(BeEquivalentTo(1))
				}
			}
		})

		It("should set extra bits", func() {
			m.Set(0, 1)
			var i uint64
			for i = 1; i < (1 << m.MType); i++ {
				m.SetBit(0, 1, i)
				var val uint64
				if i == 63 {
					val = ^uint64(0)
				} else {
					val = (1 << (i + 1)) - 1
				}
				Ω(m.Get(0, 1)).Should(BeEquivalentTo(val))
			}
		})

		It("should unset extra bits", func() {
			m.Set(0, 1)
			var i uint64
			for i = 1; i < (1 << m.MType); i++ {
				m.SetBit(0, 1, i)
				var val uint64
				if i == 63 {
					val = ^uint64(0)
				} else {
					val = (1 << (i + 1)) - 1
				}
				Ω(m.Get(0, 1)).Should(BeEquivalentTo(val))
			}
			for i = (1 << m.MType) - 1; i > 0; i-- {
				m.UnsetBit(0, 1, i)
				var val uint64
				if i == 63 {
					val = ^uint64(0)
				} else {
					val = (1 << (i + 1)) - 1
				}
				Ω(m.Get(0, 1)).Should(BeEquivalentTo(val >> 1))
			}
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

})
