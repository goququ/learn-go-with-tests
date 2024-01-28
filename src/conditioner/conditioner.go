package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CompareSign string

const (
	Bigger  = ">="
	Smaller = "<="
)

func getData() string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err.Error())
	}

	return string(data)
}

type PersonWish map[string]int

func prepareData(data string) []string {
	parts := strings.Split(data, "\n")
	parts = parts[1:]

	return parts
}

func checkTemp(data []string) string {
	result := ""
	defaultTmpWindow := [2]int{15, 30}
	var tmpWindow [2]int

	for _, row := range data {
		parts := strings.Split(row, " ")

		// If it's number of persons
		if len(parts) < 2 {
			copy(tmpWindow[:], defaultTmpWindow[:])

			if result != "" {
				result += "\n"
			}
			continue
		}

		value, _ := strconv.Atoi(parts[1])

		if parts[0] == Bigger {
			if value >= tmpWindow[0] {
				tmpWindow[0] = value
			}
		}

		if parts[0] == Smaller {
			if value <= tmpWindow[1] {
				tmpWindow[1] = value
			}
		}

		if tmpWindow[0] <= tmpWindow[1] {
			result += fmt.Sprintf("%d\n", tmpWindow[0])
		} else {
			result += "-1\n"
		}
	}

	return result
}

func checkAllTmp(s string) string {
	preparedData := prepareData(s)
	result := checkTemp(preparedData)
	return result
}

func main() {
	data := getData()
	result := checkAllTmp(data)

	fmt.Print(result)
}
