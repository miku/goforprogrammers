package main

import (
	"fmt"
)

//prints to stdout and puts an int on channel
func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	//send a value on channel
	ch <- 2
}

func main() {
	// make a channel. You need to use the make function to create channels.
	// channels can also be buffered where you can specify size. eg: ch := make(chan int, 2)

	ch := make(chan int)

	// inline goroutine. Define a function and then call it.
	// write on a channel when done
	go func() {
		fmt.Println("Hello inline")
		//send a value on channel
		// time.Sleep(100 * time.Millisecond)
		ch <- 1
	}()

	// Start function as goroutine
	go printHello(ch)
	fmt.Println("Hello from main")

	//get first value from channel.
	//and assign to a variable to use this value later
	//here that is to print it
	i := <-ch
	fmt.Println("Recieved ", i)

	// get the second value from channel
	// do not assign it to a variable because we dont want to use that
	<-ch
}
