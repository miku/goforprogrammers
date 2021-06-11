package main

import "fmt"

func hello() string {
	return "hello"
}

func g(f func() string) {
	fmt.Printf("f returned %s\n", f())
}

type stringFunc func() string

func main() {

	var myFunc stringFunc = hello
	fmt.Println(myFunc)
	// hello()
	// v := hello
	// fmt.Printf("%T, %#v\n", v, v)
	// g(hello)
}
