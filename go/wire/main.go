package main

import "log"

type Message string

func NewMessage() Message {
	return Message("Hello")
}

// Message dependency
type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

// Greeter dependency
type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Println(msg)
}

func main() {
	m := NewMessage()
	g := NewGreeter(m)
	e := NewEvent(g)

	e.Start()
}
