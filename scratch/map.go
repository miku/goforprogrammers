package main

import (
	"fmt"
)

func main() {

	var mapping = make(map[string]int)

	mapping["k"] = 123 // panic: assignment to entry in nil map

	stats := map[string]int{
		"ok":     120,
		"failed": 2,
	}
	fmt.Println(stats)       // --> map[failed:2 ok:120]
	stats["ok"]++            // modify entry
	stats["new"] = 1         // add entry
	fmt.Println(stats)       // --> map[failed:2 new:1 ok:121]
	delete(stats, "new")     // delete an entry
	fmt.Println(stats)       // --> map[failed:2 ok:121]
	fmt.Println(stats["ok"]) // --> 121
}
