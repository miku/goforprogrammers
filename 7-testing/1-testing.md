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

[embedmd]:# (../x/testabs/main.go)
```go
package testabs

import "testing"

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
```

