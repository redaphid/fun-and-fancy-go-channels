package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
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
	color.Red("%s, a broadcaster, announces it's own birth", name)
	return &broadcaster{
		name:      name,
		megaphone: make(chan string),
	}
}

func BirthListener(name string) *listener {
	color.Green("%s, a listener, politely enters the room.", name)
	return &listener{
		name: name,
	}
}

type listener struct {
	name string
}

func (l *listener) Listen(b *broadcaster) {
	color.Green("%s: Alright, I'm listening", l.name)
	megaphone := b.Megaphone()
	i := 0
	for msg := range megaphone {
		i++
		color.Blue("%s: msg#%v: %s", l.name, i, msg)
		if msg == "bye!" {
			color.Yellow("%s: I guess I should stop listening.", l.name)
			break
		}
	}
	log.Printf("%s: Goodbye, world!.", l.name)
}

const watchCount = 1
const shoutTimes = 5

func main() {
	fmt.Println("Starting up")
	aaron := BirthBroadcaster("aaron")
	larry := BirthListener("larry")
	go larry.Listen(aaron)
	for i := 0; i < shoutTimes; i++ {
		aaron.Shout(fmt.Sprintf("hey x%v", i))
	}
	aaron.Shout("bye!")
	time.Sleep(100 * time.Millisecond)
}
