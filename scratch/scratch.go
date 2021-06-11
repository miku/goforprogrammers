package main

import (
	"fmt"
)

func main() {
	var b bool
	fmt.Println(b) // false

	// all unicode letter can be identifiers

	世界 := "hello world"
	fmt.Println(世界)

	multiline := `
	Hello
	World
	`

	fmt.Println(multiline)

	// arrays
	// var twoString [2]string // twoString declared but not used

	var stringSlice []string

	fmt.Println(len(stringSlice)) // 0

	stringSlice = append(stringSlice, "a")
	stringSlice = append(stringSlice, "b")
	stringSlice = append(stringSlice, "c")

	fmt.Println(stringSlice[0])
	fmt.Println(stringSlice[1])
	sub := stringSlice[1:3]
	sub[1] = "new"
	fmt.Println(stringSlice)

	// copy

	// v := &stringSlice

	// stringSlice[0] = "1" // "panic": index out of range [0] with length 0; "recover"

	// panic("we stop here") // panic: we stop here

	// twoString[0] = "1"
	// twoString[1] = "2"

	// OOM
	// data := make([]string, 1_000_000_000) // signal: killed
	// log.Println(data)

	// var data []string // nimmt nur "slice" (~24b)
	// for i := 0; i < 1000000; i++ {
	// 	data = append(data, "")
	// }
}
