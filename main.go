package main

import (
	"fmt"
	"log"
)

//Broadcaster is a guy who shouts into his megaphone
type Broadcaster struct {
}

//Megaphone is the thing Broadcaster uses to yell.
func (b *Broadcaster) Megaphone(msg string) {

}

//Listener likes to listen to broadcaster, for some reason.
type Listener struct {
}

const watchCount = 1

func main() {

	fmt.Println("Starting up")
	a := make(chan bool)
	b := make(chan bool)
	select {
	case res := <-a:
		log.Printf("a: %v", res)
	case <-b:
		log.Printf("b: %v", res)
	}
	a <- true
	b <- false
}
