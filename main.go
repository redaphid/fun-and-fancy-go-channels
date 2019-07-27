package main

import (
	"log"
	"time"

	"github.com/fatih/color"
)

//Broadcaster is a guy who shouts into his megaphone
type Broadcaster struct {
	name      string
	megaphone chan string
}

//Megaphone returns a stream of messages.
func (b *Broadcaster) Megaphone() <-chan string {
	return b.megaphone
}

//Shout makes the Broadcaster shout
func (b *Broadcaster) Shout(msg string) {
	color.Magenta("%s shouts: %s", b.name, msg)
	b.megaphone <- msg
}

//BirthBroadcaster births a new...Broadcaster
func BirthBroadcaster(name string) *Broadcaster {
	color.Magenta("%s, a Broadcaster, announces it's own birth", name)
	return &Broadcaster{
		name:      name,
		megaphone: make(chan string),
	}
}

//BirthListener gives birth to a listener
func BirthListener(name string) *Listener {
	color.Green("%s, a Listener, politely enters the room.", name)
	return &Listener{
		name: name,
	}
}

//Listener is a meek but friendly thing that takes what Broadcaster says to heart
type Listener struct {
	name          string
	HeardMessages int
	letsListen    bool
}

//Listen listens to the Broadcaster
func (l *Listener) Listen(b *Broadcaster) {
	l.letsListen = true
	color.Green("%s: Alright, I'm listening", l.name)
	megaphone := b.Megaphone()
	for l.letsListen {
		log.Printf("beginning of loop")
		select {
		case msg := <-megaphone:
			l.HeardMessages++
			l.talkAboutIt(msg)
		case <-time.After(1 * time.Second):
			l.complainAboutBoredom()
			return
		}
	}
	color.Green("%s, a Listener, politely leaves the room.", l.name)
}

func (l *Listener) talkAboutIt(msg string) {
	color.Green("%s: msg#%v: %s", l.name, l.HeardMessages, msg)
	if msg == "bye!" {
		color.Yellow("%s: I guess I should stop listening.", l.name)
		l.letsListen = false
	}
}

func (l *Listener) complainAboutBoredom() {
	color.Blue("%s got bored after %v messages", l.name, l.HeardMessages)
}

func main() {
	log.Printf("Go away. I do nothing now.")
}
