package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func getData() string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err.Error())
	}

	return string(data)
}

func prepareData(data string) []string {
	parts := strings.Split(data, "\n")
	parts = parts[1:]

	return parts
}

func parseNumbersRow(row string) string {
	result := ""

	// Регулярное выражение для проверки формата номера
	re := regexp.MustCompile(`([A-Z]\d[A-Z]{2})|([A-Z]\d{2}[A-Z]{2})`)

	// Проверяем, соответствует ли строка формату номера
	matches := re.FindAllString(row, -1)

	// Если строка не соответствует формату номера, выводим "-"
	if len(matches) == 0 || strings.Join(matches, "") != row {
		result = "-"
	} else {
		// В противном случае выводим все совпадения, разделенные пробелами
		result = strings.Join(matches, " ")
	}

	return result
}

func checkAutoNumbers(data string) string {
	preparedData := prepareData(data)
	result := ""

	for _, row := range preparedData {
		if row == "" {
			continue
		}
		numers := parseNumbersRow(row)
		result += numers + "\n"
	}

	return result
}

func main() {
	data := getData()
	result := checkAutoNumbers(data)

	fmt.Print(result)
}
