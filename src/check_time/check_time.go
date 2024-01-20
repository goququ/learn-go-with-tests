package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
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

func formatDayOrMonth(d string) string {
	if len(d) == 1 {
		return "0" + d
	}
	return d
}

func prepareData(data string) []string {
	parts := strings.Split(data, "\n")
	parts = parts[1:]
	result := []string{}

	for _, date := range parts {
		if date == "" {
			continue
		}

		dateParts := strings.Split(date, " ")
		finalDate := dateParts[2] + "-" + formatDayOrMonth(dateParts[1]) + "-" + formatDayOrMonth(dateParts[0])

		result = append(result, finalDate)
	}

	return result
}

func checkDate(d string) Answer {
	_, err := time.Parse("2006-01-02", d)
	if err != nil {
		return No
	}
	return Yes
}

func checkAllDates(data string) string {
	preparedData := prepareData(data)
	result := ""

	for _, row := range preparedData {
		isValid := checkDate(row)
		result += string(isValid) + "\n"
	}

	return result
}

func main() {
	data := getData()
	result := checkAllDates(data)

	fmt.Print(result)
}
