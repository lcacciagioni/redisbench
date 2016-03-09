package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

// TestRedisbench is the generic handler that ginkgo uses to test
func TestRedisbench(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redisbench Suite")
}
