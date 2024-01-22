package main

import (
	"os"
	"strings"
	"testing"
)

func TestNutoNumbers(t *testing.T) {
	t.Run("check auto numbers", func(t *testing.T) {
		input, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		want, err := os.ReadFile("output.txt")
		if err != nil {
			t.Fatal(err)
		}
		got := checkAutoNumbers(string(input))

		if string(got) != string(want) {
			t.Errorf("got %q, want %q", got, want)
		}

		linesTotalWant := strings.Count(string(want), "\n")
		linesTotalGot := strings.Count(got, "\n")
		if linesTotalGot != linesTotalWant {
			t.Errorf("got %v lines in output, want %v lines", linesTotalGot, linesTotalWant)
		}
	})
}
