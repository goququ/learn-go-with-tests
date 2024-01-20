package main

import (
	"os"
	"strings"
	"testing"
)

func TestCheckShips(t *testing.T) {
	t.Run("check ships", func(t *testing.T) {
		input, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		want, err := os.ReadFile("output.txt")
		if err != nil {
			t.Fatal(err)
		}
		got := checkShips(string(input))

		if string(got) != string(want) {
			t.Errorf("got %q, want %q", got, want)
		}

		linesTotal := strings.Count(got, "\n")
		if linesTotal != 1000 {
			t.Errorf("got %q lines in output, want %q lines", linesTotal, 1000)
		}
	})
}
