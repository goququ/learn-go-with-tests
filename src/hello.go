package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func main() {
	msg := Hello()
	fmt.Println(msg)
}
