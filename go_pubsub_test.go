package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pubsub"
)

var _ = Describe("GoPubsub", func() {
	Expect(Broadcaster).NotTo(BeNil())
})
