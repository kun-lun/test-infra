package apis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kun-lun/test-infra/pkg/apis"
)

var _ = Describe("TestInfra", func() {
	Describe("BuildSampleManifest", func() {
		Context("Everything OK", func() {
			It("should produce sample manifest correctly", func() {
				testInfra := TestInfra{}
				sampleManifest := testInfra.BuildSampleManifest()
				Expect(sampleManifest).NotTo(BeNil())
			})
		})
	})

	Describe("PrepareForDeploymentCmd", func() {
		Context("Everything OK", func() {
			It("should prepare the folder correctly", func() {
				testInfra := TestInfra{}
				sampleManifest := testInfra.PrepareForDeploymentCmd()
				Expect(sampleManifest).NotTo(BeNil())
			})
		})
	})
})
