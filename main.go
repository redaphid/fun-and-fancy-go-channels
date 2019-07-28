package main

import (
	"errors"
	"log"
	"time"

	"github.com/fatih/color"
)

const timeout = 500 * time.Millisecond

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
				color.HiRed("No speakers detected. Adding message back into mic and shutting down")
				b.microphone <- msg
				return
			}
			for i := range b.speakers {
				color.HiBlack("........found speaker #%v", i)
				err := b.BlastThatSpeaker(b.speakers[i], msg, i)
				if err != nil {
					color.HiBlack("....Trying to remove troublesome speaker.")
					b.speakers = append(b.speakers[:i], b.speakers[i+1:]...)
				}
			}
		}
	}()
}

func (b *Broadcaster) BlastThatSpeaker(s chan string, msg string, speakerNum int) error {
	select {
	case s <- msg:
		color.HiBlack("........speaker #%v went through with message: %s", speakerNum, msg)
		return nil
	case <-time.After(timeout):
		color.HiRed("........speaker #%v appears to be broken. message: %s", speakerNum, msg)
		return errors.New("Speaker is broken, and therefore not blastable.")
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
func (b *Broadcaster) Shout(msg string) bool {
	color.Magenta("%s shouts: %s", b.name, msg)
	select {
	case b.microphone <- msg:
		color.Magenta("%s ...and someone heard me!", b.name)
		return true
	case <-time.After(timeout):
		color.Magenta("%s: ...and nobody heard me. I guess I'm just talking to myself, as usual :(.", b.name)
		return false
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
		color.Yellow("%s: I guess I should stop listening after %v messages", l.name, l.HeardMessages)
		l.letsListen = false
	}
}

func (l *Listener) complainAboutBoredom() {
	color.Blue("%s got bored after %v messages", l.name, l.HeardMessages)
}

func main() {
	log.Printf("Go away. I do nothing now.")
}
