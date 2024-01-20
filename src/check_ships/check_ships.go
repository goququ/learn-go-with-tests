package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Answer string

const (
	Yes Answer = "YES"
	No  Answer = "NO"
)

func getData() string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err.Error())
	}

	return string(data)
}

func prepareData(data string) [][]int {
	parts := strings.Split(data, "\n")
	parts = parts[1:]
	result := [][]int{}

	for _, raw := range parts {
		if raw == "" {
			continue
		}

		parsedRaw := make([]int, 10)

		for i, v := range strings.Split(raw, " ") {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Print(err.Error())
			}
			parsedRaw[i] = num
		}

		result = append(result, parsedRaw)
	}

	return result
}

func checkRow(row []int) Answer {
	chips := []int{0, 4, 3, 2, 1}

	for _, chip := range row {
		if chips[chip] == 0 {
			return No
		}

		chips[chip] -= 1
	}

	return Yes
}

func checkShips(data string) string {
	preparedData := prepareData(data)
	result := ""

	for _, row := range preparedData {
		isValid := checkRow(row)
		result += string(isValid) + "\n"
	}

	return result
}

func main() {
	data := getData()
	result := checkShips(data)

	fmt.Print(result)
}
