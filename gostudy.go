package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
