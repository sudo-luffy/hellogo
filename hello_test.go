package main

import "testing"

func assertHello(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Avinash", "")
		want := "Hello, Avinash"
		assertHello(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertHello(t, got, want)
	})
	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Avinash", "Spanish")
		want := "Hola, Avinash"
		assertHello(t, got, want)
	})
}
