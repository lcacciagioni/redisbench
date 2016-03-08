package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRedisbench(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redisbench Suite")
}
