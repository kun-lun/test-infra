package apis_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Apis Suite")
}
