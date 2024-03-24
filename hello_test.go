package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Joseph")
	want := "HelloJoseph"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
