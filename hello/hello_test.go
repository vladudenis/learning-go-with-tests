package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // Tell the test suite that this is a helper method
	// This way, if the test fails, the line number will be reported in our function call

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Jacques", "French")
		want := "Bonjour, Jacques"
		assertCorrectMessage(t, got, want)
	})
}
