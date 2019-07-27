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

func (b *broadcaster) Megaphone() <-chan string {
	return b.megaphone
}

//Shout makes the broadcaster shout
func (b *broadcaster) Shout(msg string) {
	b.megaphone <- msg
}

//BirthBroadcaster births a new...broadcaster
func BirthBroadcaster(name string) *broadcaster {
	log.Printf("%s, a broadcaster, announces it's own birth", name)
	return &broadcaster{
		name:      name,
		megaphone: make(chan string),
	}
}

func BirthListener(name string) *listener {
	log.Printf("%s, a listener, politely enters the room.", name)
	return &listener{
		name: name,
	}
}

type listener struct {
	name string
}

func (l *listener) Listen(b *broadcaster) {
	log.Printf("%s: Alright, I'm listening", l.name)
	for msg := range <-b.Megaphone() {
		log.Printf("%s: I heard %s!", l.name, msg)
	}
}

const watchCount = 1

func main() {
	fmt.Println("Starting up")
	aaron := BirthBroadcaster("aaron")
	larry := BirthListener("larry")
	go larry.Listen(aaron)
	aaron.Shout("I am Aaron!")
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
