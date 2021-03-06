// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/awisu2/go-study/wire-study"
)

// Injectors from wire.go:

func InitializeEvent(msg string) main.Event {
	message := main.NewMessage(msg)
	greeter := main.NewGreeter(message)
	event := main.NewEvent(greeter)
	return event
}
