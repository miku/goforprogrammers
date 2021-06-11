package main

import "fmt"

func printHello() {
	fmt.Println("Hello from printHello")
}

func main() {
	// Inline goroutine. Define a function inline and then call it.
	go func() { fmt.Println("Hello inline") }()
	// Call a function as goroutine
	go printHello()
	fmt.Println("Hello from main")
}
