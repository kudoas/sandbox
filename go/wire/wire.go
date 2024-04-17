//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func BuildEvent() (Event, error) {
	wire.Build(Set)
	return Event{}, nil
}
