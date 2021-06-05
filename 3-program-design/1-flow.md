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

A struct is flatter than a class. It's data first.

[embedmd]:# (../x/rgba/main.go)
```go
package main

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point struct {
	X, Y int
}

// A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
// It is well-formed if Min.X <= Max.X and likewise for Y. Points are always
// well-formed. A rectangle's methods always return well-formed outputs for
// well-formed inputs.
//
// A Rectangle is also an Image whose bounds are the rectangle itself. At
// returns color.Opaque for points in the rectangle and color.Transparent
// otherwise.
type Rectangle struct {
	Min, Max Point
}

// RGBA is an in-memory image whose At method returns color.RGBA values.
type RGBA struct {
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}

func main() {}
```

Favors composition over inheritance. Easy to comprehend, harder to realize.

Possible problems with classic approach:

* too many levels of indirection
* premature abstractions

> Anecdote: On a short tour through some code, someone using a classic
> object-oriented language showed me a code center-piece: An interface with a
> good dozen methods (including data and more UI related names). It seemed more
> like a *bundle* than an real *abstraction*.

