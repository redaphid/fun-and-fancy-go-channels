package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pubsub"
)

var _ = Describe("GoPubsub", func() {
	var aaron *Broadcaster
	BeforeEach(func() {
		aaron = BirthBroadcaster("aaron")
	})

	Describe("When broadcasting to one person", func() {
		var larry *Listener
		BeforeEach(func() {
			larry = BirthListener("larry")
			go larry.Listen(aaron)
		})
		Describe("When broadcasting 5 times", func() {
			BeforeEach(func(done Done) {
				for i := 0; i < 5; i++ {
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				done <- true
			}, 1)
			It("Should have Larry hear 5 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(5))
			})
		})
		Describe("When broadcasting 10 times, but says 'bye!' on message #3", func() {
			BeforeEach(func(done Done) {
				for i := 0; i < 10; i++ {
					if i == 2 {
						aaron.Shout("bye!")
						continue
					}
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				done <- true
			}, 1)
			It("Should have Larry hear 3 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(3))
			})
		})
	})

	Describe("When broadcasting to 2 people", func() {
		var larry *Listener
		var leonard *Listener
		BeforeEach(func() {
			larry = BirthListener("larry")
			leonard = BirthListener("leonard")
			go larry.Listen(aaron)
			go leonard.Listen(aaron)
		})
		Describe("When broadcasting 5 times", func() {
			BeforeEach(func(done Done) {
				for i := 0; i < 5; i++ {
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				done <- true
			}, 1)
			It("Should have Larry hear 5 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(5))
			})
		})
	})
})
