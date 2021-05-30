# Intro

> Go is an open source programming language that makes it easy to build simple,
> reliable, and efficient software. --
> [https://golang.org/](https://golang.org/)

## Spec

* [ref/spec](https://golang.org/ref/spec)

## Other Resources

![](https://raw.githubusercontent.com/miku/goforprogrammers/master/static/go_21_s.jpg)
![](https://raw.githubusercontent.com/miku/goforprogrammers/master/static/go_prog_lang_s.jpg)

* [more Go books](https://www.google.com/search?channel=fs&client=ubuntu&q=golang+books)

![](https://raw.githubusercontent.com/miku/goforprogrammers/master/static/goog_books.png)

## Hello, World

[embedmd]:# (../x/helloworld/main.go)

## Execution Model

* Go is a compiled, strongly typed, garbage-collected language. An executable
  lives in a package main and contains a function main as entry point.

## Variables

* There are four notations to
  [declare](https://golang.org/ref/spec#Variable_declarations) and initialize
variables.

You'll mostly the following two (type inference):

[embedmd]:# (../x/vars/main.go /^.*a := 3.14/ /var b = 1.0.*/)

You can name the type explicitly. The type follows the identifier (as
opposed to C, Java, ...):

[embedmd]:# (../x/vars/main.go /^.*c int/ /var d int8 =.*/)

* The default integer type will by [int](https://play.golang.org/p/z1emUwkp1rL).
* Every type has a [zero value](https://golang.org/ref/spec#The_zero_value)

> **When storage is allocated for a variable**, either through a declaration or
> a call of new, or when a new value is created, either through a composite
> literal or a call of make, **and no explicit initialization is provided, the
> variable or value is given a default value**. Each element of such a variable
> or value is set to the **zero value for its type**: false for booleans, 0 for
> numeric types, "" for strings, and nil for pointers, functions, interfaces,
> slices, channels, and maps. This initialization is done recursively, so for
> instance each element of an array of structs will have its fields zeroed if
> no value is specified.


## Types

### Primitive Types

* boolean, called `bool`, predeclared constants `true` and `false`
* [https://golang.org/ref/spec#Numeric_types](https://golang.org/ref/spec#Numeric_types)


[embedmd]:# (https://golang.org/ref/spec go /^uint8.*/ /^rune.*/)


There are implementation specific predeclared types for `uint` (32 or 64 bits),
`int` (same as `uint`), `uintptr` large enough to store pointer value.

* [proposal: spec: remove complex numbers #19921](https://github.com/golang/go/issues/19921) (Apr 11, 2017)

> Go supports complex numbers, but ~nobody uses them.

* String type, `s := "Hello 世界"` - all Go source is UTF-8 encoded, double quotes or backticks for multiline strings; strings are *immutable*.

Any UTF-8 character may be used as a variable name
([play](https://play.golang.org/p/eU4VDR0-jpE)):

[embedmd]:# (../x/hellojp/main.go)

### Arrays and Slices

There are [arrays](https://golang.org/ref/spec#Array_types) and
[slices](https://golang.org/ref/spec#Slice_types) in Go. Arrays carry their
length as part of the type, slices are dynamically sized arrays. You will
mostly encounter slices.

Example (on [play](https://play.golang.org/p/bsgboAZ82jH)):

[embedmd]:# (../x/slicehello/main.go)

#### Creating a slice

[embedmd]:# (../x/sliceinit/main.go)

#### The append builtin function

* [Go Tour: Appending to slice](https://tour.golang.org/moretypes/15), it is a
  [built-in](https://golang.org/pkg/builtin/#append) function

[embedmd]:# (../x/sliceappend/main.go)

### Structs

* Struct types group zero or more fields (name, type) into a type.

Example ([play](https://play.golang.org/p/RYOrTC-CqPu)):

[embedmd]:# (../x/structinit/main.go)

Struct types support embedding (allows for reuse and composition) and tags (a
weakly typed, but lightweight approach to annotations).

Example ([play](https://play.golang.org/p/KgKB3jqCapV)):

[embedmd]:# (../x/structembed/main.go)

### Maps

* Map types are the way to implement hash maps or associative arrays.

Example ([play](https://play.golang.org/p/kMy733hPcFT)):

[embedmd]:# (../x/mapinit/main.go)
