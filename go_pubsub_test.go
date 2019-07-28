package main_test

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "pubsub"
)

var _ = Describe("GoPubsub", func() {
	testLog := color.New(color.FgBlack).Add(color.BgWhite)
	var aaron *Broadcaster
	BeforeEach(func() {
		testLog.Print("And so it begins")
		aaron = BirthBroadcaster("aaron")
	})

	AfterEach(func() {
		testLog.Print("And so it ends")
	})

	Describe("When broadcasting to one person", func() {
		var larry *Listener
		BeforeEach(func() {
			larry = BirthListener("larry")
			go larry.Listen(aaron)
		})
		Describe("When broadcasting 5 times", func() {
			BeforeEach(func() {
				for i := 0; i < 5; i++ {
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				<-time.After(500 * time.Millisecond)
			})
			It("Should have Larry hear 5 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(5))
			})
		})
		Describe("When broadcasting 10 times, but says 'bye!' on message #3", func() {
			BeforeEach(func(done Done) {
				i := 0
				for {
					if i == 2 {
						aaron.Shout("bye!")
						continue
					}
					isAaronOk := aaron.Shout(fmt.Sprintf("hey x%v", i))
					color.Green("Aaron is fine.")
					if !isAaronOk {
						break
					}
					i = i + 1
				}
				close(done)
			}, 1)
			It("Should have Larry hear 3ish messages.", func() {
				Expect(larry.HeardMessages).To(BeNumerically("<", 6))
				Expect(larry.HeardMessages).To(BeNumerically(">", 1))
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
			BeforeEach(func() {
				for i := 0; i < 5; i++ {
					aaron.Shout(fmt.Sprintf("hey x%v", i))
				}
				<-time.After(500 * time.Millisecond)
			})
			It("Should have Larry hear 5 messages.", func() {
				Expect(larry.HeardMessages).To(Equal(5))
			})
			It("Should have Leonard hear 5 messages.", func() {
				Expect(leonard.HeardMessages).To(Equal(5))
			})
		})
	})
})
