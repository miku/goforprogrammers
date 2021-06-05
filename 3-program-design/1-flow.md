# Design Flow

Go is object-oriented, but takes a different approach.

* no type hierarchy (and much less emphasis on inheritance, which does not really exist)
* typically, you do not start with your interfaces (as in other languages)

First ideas:

* What is this program going to do?
* And once some program flow emerges, we may introduce new types.

> Whenever someone from Java or C++ or C# comes to Go, they look for "class",
> find "struct", and stop looking. #golang

-- https://twitter.com/rob_pike/status/942528032887029760


Notably:

* you can add methods (behavior) to any type

[embedmd]:# (../x/addmethod/main.go)
```go
package main

import "fmt"

// We cannot define methods directly on int, but on a derived type.
// 	// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go
type xint int

func (v xint) isEven() bool {
	return v%2 == 0
}

func main() {
	var x, y xint = 2, 3
	fmt.Printf("%v %v\n", x, x.isEven())
	fmt.Printf("%v %v\n", y, y.isEven())
}
```
