# The Go Tool

The Go tool allows to manage many phases of the development cycle, especially:

* go get
* go fmt
* go vet
* go run
* go test
* go build

```
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.
...
```

## Download and install packages

Downloading and installing an executable.

```
$ go get -v github.com/miku/binpic/cmd/binpic
```

Download package and subpackges:

```
$ go get -v github.com/miku/binpic/...
```

## Formatting code

> Without an explicit path, it processes the standard input. Given a file, it
operates on that file; given a directory, it operates on all .go files in that
directory, recursively.

```
$ go fmt -w ./...
```

Today, the [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) is
used a lot.

You can install it via:

```
$ go get golang.org/x/tools/cmd/goimports
```

Editors plugins might download these tools for you.

## Checks with go vet

Run checks on current package and subpackages.

```
$ go vet ./...
```

* [cmd/vet](https://golang.org/cmd/vet/)

See also: [goreportcard](https://goreportcard.com/)

## Run code

Run compiles and runs the named main Go package.

```
$ go run main.go
```


## Testing code

Run all tests verbosly.

```
$ go test -v ./...
```

Runs `go vet` subset, too.

> As part of building a test binary, go test runs go vet on the package and its
test source files to identify significant problems. If go vet finds any
problems, go test reports those and does not run the test binary. Only a
high-confidence subset of the default go vet checks are used.

Can run in two modes.

> The first, called **local directory mode**, occurs when go test is
invoked with no package arguments (for example, 'go test' or 'go
test -v'). In this mode, go test compiles the package sources and
tests found in the **current directory** and then runs the resulting
test binary. In this mode, caching (discussed below) is disabled.
After the package test finishes, go test prints a summary line
showing the test status ('ok' or 'FAIL'), package name, and elapsed
time.

> The second, called **package list mode**, occurs when go test is invoked
with explicit package arguments (for example 'go test math', 'go
test ./...', and even 'go test .'). In this mode, go test compiles
and tests each of the packages listed on the command line. If a
package test passes, go test prints only the final 'ok' summary
line. If a package test fails, go test prints the full test output.
If invoked with the -bench or -v flag, go test prints the full
output even for passing package tests, in order to display the
requested benchmark results or verbose logging. After the package
tests for all of the listed packages finish, and their output is
printed, go test prints a final 'FAIL' status if any package test
has failed.

## Building binaries

```
$ go build -o hello main.go
```

Allows to pass flags to compiler and linker. Cross-compilation is supported.
Selective build is available via build tags.

