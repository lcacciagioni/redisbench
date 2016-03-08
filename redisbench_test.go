package main_test

import (
	"math/rand"

	. "github.com/lcacciagioni/redisbench"
	"github.com/rafaeljusto/redigomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redisbench", func() {
	c := redigomock.NewConn()
	numMsg := rand.Intn(10000)
	min := 1
	max := 1000
	Describe("With NodeStressString", func() {
		result1 := NodeStressString(c, numMsg, min, max)
		result2 := NodeStressString(c, numMsg, min, max)
		It("Must return a result", func() {
			Expect(result1).ToNot(Equal(StressResult{}))
			Expect(result2).ToNot(Equal(StressResult{}))
		})
		Context("Given 2 runs results must not be equal", func() {
			It("Must return 2 differen results", func() {
				Expect(result1).ToNot(Equal(result2))
			})
		})
	})
	Describe("With NodeStressBytes", func() {
		result1 := NodeStressBytes(c, numMsg, min, max)
		result2 := NodeStressBytes(c, numMsg, min, max)
		It("Must return a result", func() {
			Expect(result1).ToNot(Equal(StressResult{}))
			Expect(result2).ToNot(Equal(StressResult{}))
		})
		Context("Given 2 runs results must not be equal", func() {
			It("Must return 2 differen results", func() {
				Expect(result1).ToNot(Equal(result2))
			})
		})
	})
})
