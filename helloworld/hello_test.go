package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joseph", "")
		want := "Hello, Joseph"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say '你好' to people when the second argument specifies Chinese", func(t *testing.T) {
		got := Hello("Joseph", "Chinese")
		want := "你好, Joseph"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Bonjour' to people when the second argument specifies French", func(t *testing.T) {
		got := Hello("Joseph", "French")
		want := "Bonjour, Joseph"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hola' to people when the second argument specifies Spanish", func(t *testing.T) {
		got := Hello("Joseph", "Spanish")
		want := "Hola, Joseph"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
