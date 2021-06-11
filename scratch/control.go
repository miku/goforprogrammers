package main

import (
	"fmt"
	"log"
)

func main() {

	var k int32
	for {
		k++
		if k == 2_147_483_647 {
			break
		}
	}
	log.Printf("%d %d", k, k+1)

	// A slice.
	var abc = []string{"a", "b", "c"}

	for i, v := range abc {
		fmt.Printf("%d %v\n", i, v)
	}

	for _, v := range abc {
		fmt.Printf("%v\n", v)
	}

	// index-only
	for v := range abc {
		fmt.Printf("%v\n", v)
	}

	if 2 > 1 {
		log.Println("yay")
	}
}
