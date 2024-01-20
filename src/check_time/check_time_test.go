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
		got := checkAllDates(string(input))

		if string(got) != string(want) {
			t.Errorf("got %v,\n\n want %v", got, string(want))
		}

		linesTotalWant := strings.Count(string(want), "\n")
		linesTotalGot := strings.Count(got, "\n")
		if linesTotalGot != linesTotalWant {
			t.Errorf("got %v lines in output, want %v lines", linesTotalGot, linesTotalWant)
		}
	})
}
