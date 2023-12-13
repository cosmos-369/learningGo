package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Rick")
	got := buffer.String()
	want := "Hello, Rick"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
