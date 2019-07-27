package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pubsub"
)

var _ = Describe("GoPubsub", func() {
	Describe("When broadcasting to one person", func() {
		var (
			aaron *Broadcaster
			larry *Listener
		)
		BeforeEach(func() {
			aaron = BirthBroadcaster("aaron")
			larry = BirthListener("larry")
			go larry.Listen(aaron)
		})
		Describe("When broadcasting 5 times", func() {
			BeforeEach(func() {
				for i := 0; i < 5; i++ {
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				// log.Println("Done broadcasting.")
			})
			It("Should have Larry hear 5 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(5))
			})
		})
	})
})
