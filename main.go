package main

import (
	"log"
	"time"

	"github.com/fatih/color"
)

const timeout = 100 * time.Millisecond

//Broadcaster is a guy who shouts into his megaphone
type Broadcaster struct {
	name       string
	speakers   []chan string
	microphone chan string
	started    bool
	// mics []
}

//WireupSoundSystem is the insane engineer who decided to do this the hard way.
func (b *Broadcaster) WireupSoundSystem() {
	go func() {
		for msg := range b.microphone {
			color.HiBlack("........internal system: %s", msg)
			if len(b.speakers) == 0 {
				color.HiRed("No speakers detected. Adding message back into mic")
				b.microphone <- msg
				return
			}
			for i := range b.speakers {
				color.HiBlack("........found speaker #%v", i)
				go b.BlastThatSpeaker(b.speakers[i], msg, i)
			}
		}
	}()
}

func (b *Broadcaster) BlastThatSpeaker(s chan string, msg string, speakerNum int) {
	select {
	case s <- msg:
		color.HiBlack("........speaker #%v went through with message: %s", speakerNum, msg)
	case <-time.After(timeout):
		color.HiRed("........speaker #%v appears to be broken. message: %s", speakerNum, msg)
		// close(s)
	}
}

func (b *Broadcaster) StartSoundSystem() {
	if len(b.speakers) == 0 {
		// b.StopSoundSystem()
		return
	}
	b.WireupSoundSystem()
}

//Megaphone returns a stream of messages.
func (b *Broadcaster) Megaphone() <-chan string {
	newSpeaker := make(chan string)
	b.speakers = append(b.speakers, newSpeaker)
	b.StartSoundSystem()
	return newSpeaker
}

//Shout makes the Broadcaster shout
func (b *Broadcaster) Shout(msg string) {
	color.Magenta("%s shouts: %s", b.name, msg)
	select {
	case b.microphone <- msg:
		color.Magenta("%s ...and someone heard me!", b.name)
	case <-time.After(timeout):
		color.Magenta("%s: ...and nobody heard me. I guess I'm just talking to myself, as usual :(.", b.name)
		close(b.microphone)
	}
}

//BirthBroadcaster births a new...Broadcaster
func BirthBroadcaster(name string) *Broadcaster {
	color.Magenta("%s, a Broadcaster, announces it's own birth", name)
	b := Broadcaster{
		name:       name,
		speakers:   make([]chan string, 0),
		microphone: make(chan string),
	}
	b.StartSoundSystem()
	return &b
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
