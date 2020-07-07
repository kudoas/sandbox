package main

import (
	"testing"
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := Abs(-30)
	expected := 30
	if got != expected {
		t.Errorf("Abs(-30) = %d; want 30", got)
	}
}
