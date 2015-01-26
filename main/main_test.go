package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("exits 1 if no uaa-public-key is provided", func() {
		session := RoutingApi()
		Eventually(session).Should(Exit(1))
	})
})

func RoutingApi(args ...string) *Session {
	path, err := Build("github.com/pivotal-cf-experimental/routing-api/main")
	Expect(err).NotTo(HaveOccurred())

	session, err := Start(exec.Command(path, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	return session
}