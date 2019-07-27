package main

import (
	"fmt"
	"log"
)

//Broadcaster is a guy who shouts into his megaphone
type broadcaster struct {
	name      string
	megaphone chan string
}

//Shout makes the broadcaster shout
func (b *broadcaster) Shout(msg string) {
	b.Megaphone <- msg
}

//BirthBroadcaster births a new...broadcaster
func BirthBroadcaster(name string) *broadcaster {
	log.Printf("%s, a broadcaster, announces it's own birth", name)
	return &broadcaster{
		megaphone: make(chan string),
	}
}

//Listener likes to listen to broadcaster, for some reason.
type Listener struct{}

const watchCount = 1

func main() {
	fmt.Println("Starting up")
	aaron := BirthBroadcaster("aaron")
}

func example() {

	fmt.Println("Starting up")
	a := make(chan bool)
	b := make(chan bool)
	go func() {
		for {
			select {
			case res := <-a:
				log.Printf("a: %v", res)
			case res := <-b:
				log.Printf("b: %v", res)
			default:
				log.Println("We're done!")
				return
			}
			log.Println("Looping")
		}
		log.Println("exited loop")
	}()
	a <- false
	b <- true
}
