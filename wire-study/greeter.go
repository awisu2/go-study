package main

import "fmt"

type Message struct {
	Word string
}

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg.Word)
}

func NewMessage(word string) Message {
	return Message{
		Word: word,
	}
}

func NewMessage2() Message {
	return Message{
		Word: "not empty!",
	}
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
