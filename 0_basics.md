# Basics

## Spec

* [ref/spec](https://golang.org/ref/spec)

## Execution Model

* Go is a compiled, strongly typed, garbage-collected language. An executable
  lives in a package main and contains a function main as entry point.

## Variables

* There are four notations to declare and initialize variables.

You'll mostly the following two (type inference):

```
a := 1.17
var b = 1.0
```

You can control the type explicitly. The type follows the identifier (as
opposed to C, Java, ...):

```
var a int         // zero value
var a int8 = 10.0 // explicit type
```

## Types

* boolean, called `bool`, predeclared constants `true` and `false`
* [https://golang.org/ref/spec#Numeric_types](https://golang.org/ref/spec#Numeric_types)

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

There are implementation specific predeclared types for `uint` (32 or 64 bits),
`int` (same as `uint`), `uintptr` large enough to store pointer value

* [proposal: spec: remove complex numbers #19921](https://github.com/golang/go/issues/19921) (Apr 11, 2017)

> Go supports complex numbers, but ~nobody uses them.

* String type, `s := "Hello 世界"` - all Go source is UTF-8 encoded, double quotes or backticks for multiline strings; strings are *immutable*

Any UTF-8 character may be used as a variable name
([play](https://play.golang.org/p/eU4VDR0-jpE)):

```go
package main

import (
    "fmt"
)

func main() {
    世界  := "hello world"
    fmt.Println(世界)
}
```

There are [arrays](https://golang.org/ref/spec#Array_types) and
[slices](https://golang.org/ref/spec#Slice_types) in Go. Arrays carry their
length as part of the type, slices are dynamically sized arrays. You will
mostly encounter slices.

Example (on [play](https://play.golang.org/p/bsgboAZ82jH)):

```golang
package main

import (
    "fmt"
)

func main() {
    data := []int{1, 2, 3, 4, 5}
    fmt.Printf("%v has length %d", data, len(data))
}
```

* Struct types group zero or more fields (name, type) into a type.

Example ([play](https://play.golang.org/p/RYOrTC-CqPu)):

```golang
package main

import "fmt"

type Location struct {
    Lat  float64
    Long float64
}

type Peer struct {
    Name     string
    Location Location
}

func main() {
    peer := Peer{
        Name: "martin",
        Location: Location{
            Lat:  51.33962,
            Long: 12.37129,
        }}
    fmt.Println(peer)
    // {martin {51.33962 12.37129}}
}
```

Struct types support embedding (allows for reuse and composition) and tags (a weakly typed, but lightweight approach to annotations).

Example ([play](https://play.golang.org/p/KgKB3jqCapV)):

```golang
package main

import (
    "encoding/json"
    "fmt"
)

type Location struct {
    Lat  float64 `json:"lat"`
    Long float64 `json:"long"`
}

type Peer struct {
    Name     string   `json:"name"`
    Location Location `json:"loc"`
}

func main() {
    peer := Peer{
        Name: "martin",
        Location: Location{
            Lat:  51.33962,
            Long: 12.37129,
        }}
    b, _ := json.Marshal(peer)
    fmt.Println(string(b))
    // {"name":"martin","loc":{"lat":51.33962,"long":12.37129}}
}
```
