package main_test

import (
	. "github.com/lcacciagioni/redisbench"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/rand"
)

var _ = Describe("Randomizers", func() {
	// Here we add 10 just because any random string is uselless with a smaller size
	size1 := rand.Intn(10000) + 10
	size2 := rand.Intn(10000) + 10
	min := 1
	max := 10000
	Describe("With RandStringBytesMaskImprSrc", func() {
		It("Should return a string of the specified size", func() {
			Expect(len(RandStringBytesMaskImprSrc(size1))).To(Equal(size1))
		})
		Context("In 2 secuential runs", func() {
			It("Must not return the same value", func() {
				Expect(RandStringBytesMaskImprSrc(size1)).ToNot(Equal(RandStringBytesMaskImprSrc(size1)))
			})
		})
	})
	Describe("With RandStringBytesMaskImprSrcBytes", func() {
		It("Should return a []]byte of the specified size", func() {
			Expect(len(RandStringBytesMaskImprSrcBytes(size2))).To(Equal(size2))
		})
		Context("In 2 secuential runs", func() {
			It("Must not return the same value", func() {
				Expect(RandStringBytesMaskImprSrcBytes(size2)).ToNot(Equal(RandStringBytesMaskImprSrc(size2)))
			})
		})
	})
	Describe("With RandRangeString", func() {
		It("Must Return a string not bigger than max and not smaller than min", func() {
			Expect(len(RandRangeString(min, max))).To(BeNumerically(">=", min))
			Expect(len(RandRangeString(min, max))).To(BeNumerically("<=", max))
		})
		Context("In two secuential runs", func() {
			It("must not return the same value", func() {
				Expect(RandRangeString(min, max)).NotTo(Equal(RandRangeString(min, max)))
			})
		})
	})
	Describe("With RandRangeBytes", func() {
		It("Must Return a string not bigger than max and not smaller than min", func() {
			Expect(len(RandRangeBytes(min, max))).To(BeNumerically(">=", min))
			Expect(len(RandRangeBytes(min, max))).To(BeNumerically("<=", max))
		})
		Context("In two secuential runs", func() {
			It("must not return the same value", func() {
				Expect(RandRangeBytes(min, max)).NotTo(Equal(RandRangeBytes(min, max)))
			})
		})
	})
})
