# IO

Package [io](https://golang.org/pkg/io/) provides basic interfaces to I/O
primitives. Its primary job is to wrap existing implementations of such
primitives, such as those in package os, into shared public interfaces that
abstract the functionality, plus some other related primitives.

Tour sections 21-23 contain a bit about [readers](https://tour.golang.org/methods/21).

## Interfaces

The most important interfaces are [io.Reader](https://golang.org/pkg/io/#Reader)
and [io.Writer](https://golang.org/pkg/io/#Writer). There are many implementations.

A few implementation have been collected in this [repo](https://github.com/miku/exploreio).

## Utilities

The package contains utilities, of which
[io.Copy](https://golang.org/pkg/io/#Copy) is popular. It allows to transfer
data from a reader to a writer.

## Exercise

* [A simple cat in Go](template1/main.go) ([Answer](exercise1/main.go))