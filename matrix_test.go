package tinygraph_test

import (
	. "github.com/iancmcc/tinygraph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boolean Matrix", func() {

	var (
		m        Matrix
		cellsize MatrixType = Bit
	)

	BeforeEach(func() {
		m = NewMatrix(5, cellsize)
	})

	Describe("Setting a value in a bit matrix", func() {
		Context("With a large enough matrix", func() {
			It("should be able to set the bit", func() {
				By("setting a bit within bounds")
				m.Set(2)
				Expect(m.Get(2)).To(Equal(uint64(1)))
			})
		})
	})

})
