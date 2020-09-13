package hdflibwallet_test

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHdflibwallet(t *testing.T) {
	RegisterFailHandler(Fail)
	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "Hdflibwallet Suite")
}
