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

