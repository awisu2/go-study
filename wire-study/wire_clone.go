//go:build wireinject

package main

import "github.com/google/wire"

// 引数は型があっていれば良いみたい
// NewMessage need `message string`
func InitializeEventClone(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
