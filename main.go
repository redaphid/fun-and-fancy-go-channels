package main

import (
	"fmt"
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

}
