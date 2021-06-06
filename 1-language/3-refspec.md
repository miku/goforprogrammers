# Selected notes on ref/spec

A few picks from [ref/spec](https://golang.org/ref/spec).

## Semicolons

Even if you do not see them typically, Go uses
  [semicolons](https://golang.org/ref/spec#Semicolons) - they just get
  automatically inserted into the token stream.

## Identifiers

You can name a variable `αβ` - *first character in an identifier must be a letter*.

> Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo as Unicode letters [...]

* [Unicode 4.5 General Category](https://www.unicode.org/versions/Unicode8.0.0/ch04.pdf#page=17)


## Rune literals

A rune is an alias for int32, representing a Unicode code point.

It is written with single quotes.

> A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'.

[embedmd]:# (../x/runevalue/main.go)
```go
package main

import "fmt"

func main() {
	a := 'a'
	fmt.Printf("%T %c %d %x %v\n", a, a, a, a, a == 97)

	v := '⧉'
	fmt.Printf("%T %c %d %x %v\n", v, v, v, v, v == 10697)

	// int32 a 97 61 true
	// int32 ⧉ 10697 29c9 true
}
```

More on that in [https://blog.golang.org/strings](https://blog.golang.org/strings).
