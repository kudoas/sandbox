package main

import (
	"errors"
	"log"
	"time"
)

type Message string

func ProvideMessage() Message {
	return Message("Hello")
}

// Message dependency
type Greeter struct {
	Message Message
	Grumpy  bool
}

func ProvideGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// Greeter dependency
type Event struct {
	Greeter Greeter
}

func ProvideEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Println(msg)
}

func main() {
	// m := NewMessage()
	// g := NewGreeter(m)
	// e, err := NewEvent(g)
	e, err := BuildEvent()
	if err != nil {
		log.Fatalf("failed to create event: %v", err)
	}
	e.Start()
}
