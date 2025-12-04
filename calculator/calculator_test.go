// calculator_test.go
package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator operations", func() {

	Context("Addition", func() {
		It("adds two numbers correctly", func() {
			Expect(Add(2, 3)).To(Equal(5.0))
		})
	})

	Context("Subtraction", func() {
		It("subtracts two numbers correctly", func() {
			Expect(Subtract(5, 3)).To(Equal(2.0))
		})
	})

	Context("Multiplication", func() {
		It("multiplies two numbers correctly", func() {
			Expect(Multiply(4, 3)).To(Equal(12.0))
		})
	})

	Context("Division", func() {
		It("divides two numbers correctly", func() {
			result, err := Divide(10, 2)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(5.0))
		})

		It("returns an error when dividing by zero", func() {
			_, err := Divide(10, 0)
			Expect(err).ToNot(BeNil())
		})
	})

})
