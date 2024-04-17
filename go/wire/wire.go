//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func BuildEvent() (Event, error) {
	wire.Build(ProvideEvent, ProvideGreeter, ProvideMessage)
	return Event{}, nil
}
