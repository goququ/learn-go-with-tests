package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CompareSign string

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

func compressRow(numbers []string) string {
	if len(numbers) == 1 {
		sequance := fmt.Sprintf("%s %d", numbers[0], 0)
		return fmt.Sprintf("%d\n%s", 2, sequance)
	}

	var result []string
	start, _ := strconv.Atoi(numbers[0])
	prev := start

	for i := 1; i < len(numbers); i++ {
		curr, _ := strconv.Atoi(numbers[i])

		if ((curr-prev) != 1 && start <= curr) || ((curr-prev) != -1 && start >= curr) {
			result = append(result, strconv.Itoa(start))
			result = append(result, strconv.Itoa(prev-start))
			start = curr
		}

		prev = curr

		if i == len(numbers)-1 {
			result = append(result, strconv.Itoa(start))
			result = append(result, strconv.Itoa(curr-start))
		}
	}

	sequance := strings.Join(result, " ")

	return fmt.Sprintf("%d\n%s", len(result), sequance)
}

func compressAllRows(data []string) string {
	result := ""

	for index, row := range data {
		rowSrtingSlice := strings.Split(row, " ")

		// this is count of elements in next row
		if index%2 == 0 || len(rowSrtingSlice) == 0 {
			continue
		}

		result += fmt.Sprintf("%s\n", compressRow(rowSrtingSlice))
	}

	return result
}

func checkAllTmp(s string) string {
	preparedData := prepareData(s)
	result := compressAllRows(preparedData)

	return result
}

func main() {
	data := getData()
	result := checkAllTmp(data)

	fmt.Print(result)
}
