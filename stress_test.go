package main_test

import (
	. "github.com/lcacciagioni/redisbench"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stress", func() {
	Describe("With StressCluster", func() {
		cluster := []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"}
		It("Must fail when no redis is there", func() {
			Expect(StressCluster(cluster, 1, 10, 100)).To(HaveOccurred())
		})
	})
	Describe("With StressNode", func() {
		It("Must fail when no redis is there", func() {
			Expect(StressNode("127.0.0.1:6379", 1, 10, 100)).To(HaveOccurred())
		})
	})
})
