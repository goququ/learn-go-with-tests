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
		got := checkAllTmp(string(input))

		if string(got) != string(want) {
			t.Errorf("got \n\n%v\n\n want \n\n%v", got, string(want))
		}

		linesTotalWant := strings.Count(string(want), "\n")
		linesTotalGot := strings.Count(got, "\n")
		if linesTotalGot != linesTotalWant {
			t.Errorf("got \n\n%v lines in output, want %v lines", linesTotalGot, linesTotalWant)
		}
	})
}
