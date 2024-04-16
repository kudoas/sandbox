//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func BuildEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
