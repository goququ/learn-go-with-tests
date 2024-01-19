package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	msg := Hello("")
	fmt.Println(msg)
}
