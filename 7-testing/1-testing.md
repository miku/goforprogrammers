# Testing

## The go test subcommand

Go has testing support in the [standard
library](https://golang.org/pkg/testing/) and in the [go
tool](https://golang.org/cmd/go/#hdr-Test_packages).

```
usage: go test [build/test flags] [packages] [build/test flags & test binary flags]
```

Example: Testing in *local directory mode*

```
$ go test
```

Be verbose:

```
$ go test -v
```

Example: Package list mode, e.g. package name or patterns like `.` or `./...`.

```
$ go test math
$ go test .
$ go test ./...
```

## Naming conventions

Typical naming pattern, e.g. `..._test.go` for tests and benchmarks and `example_...` for examples.

Example from [sort](https://github.com/go4org/go4/tree/master/sort):

```
example_interface_test.go
example_keys_test.go
example_multi_test.go
example_slice_test.go
example_test.go
example_wrapper_test.go
export_test.go
genzfunc.go
search.go
search_test.go
sort.go
sort_test.go
zfuncversion.go
```

Any function in test files following a naming convention will get picked up:

[embedmd]:# (../x/testabs/main.go /import/ $)
```go
import "testing"

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
```

The `testing.T` type helps running the test, e.g. logging, skipping, temporary
directories, etc.

## Table-Driven Tests

A simple pattern:

* [https://github.com/golang/go/wiki/TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)

> Writing good tests is not trivial, but in many situations a lot of ground can
> be covered with table-driven tests: Each table entry is a complete test case
> with inputs and expected results, and sometimes with additional information
> such as a test name to make the test output easily readable.

[embedmd]:# (../x/tabledriven/main.go /^import/ $)
```go
import "testing"

var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	{"%+a", "[%+a]"},
	{"%#a", "[%#a]"},
	{"% a", "[% a]"},
	{"%0a", "[%0a]"},
	{"%1.2a", "[%1.2a]"},
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
	{"%-+1.2a", "[%+-1.2a]"},
	{"%-+1.2abc", "[%+-1.2a]bc"},
	{"%-1.2abc", "[%-1.2a]bc"},
}

func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
```
