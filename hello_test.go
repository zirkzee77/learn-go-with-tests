package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
  
  t.Run("saying hello to people in spanish", func(t *testing.T) {
    got := Hello("Raul", "spanish")
    want := "Hola, Raul"
    assertCorrectMessage(t, got, want)
  })

  t.Run("saying hello to people in french", func(t *testing.T) {
    got := Hello("Emmanuel", "french")
    want := "Bonjour, Emmanuel"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
