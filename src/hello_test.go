package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		if got != want {
			t.Errorf("Hello('Chris') = %q, want %q", got, want)
		}
	})

	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("Hello('') = %q, want %q", got, want)
		}
	})
}
