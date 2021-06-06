# Intro

> Go is an [open source](https://golang.org/LICENSE) programming language that makes it easy to build simple,
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
```go
package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
```

## Execution Model

* Go is a compiled, strongly typed, garbage-collected language. An executable
  lives in a package main and contains a function main as entry point.

## Where does Go code live?

All Go code is organized in packages.

> Go code is organized into packages. Within a package, code can refer to any
> identifier (name) defined within, while clients of the package may only
> reference the package's exported types, functions, constants, and variables

## What is a package?

> A package in Go is simply a directory/folder with one or more .go files inside
> of it. -- [https://rakyll.org/style-packages/](https://rakyll.org/style-packages/)


## Import resolution

At import time, Go first look for packages in `$GOROOT/src`, which contains the
standard library (see Extra: Stdlib packages).

Depending on the project setup, e.g. GOPATH (early) or module-aware (current)
code will be searched in GOPATH or the module cache.

## Getting Go packages

There is no extra package manager - the go tool is used to install packages:

```
$ go get golang.org/x/tools/cmd/goimports
```

> [Go] get resolves its command-line arguments to packages at specific module
versions, updates go.mod to require those versions, downloads source code into
the module cache, then builds and installs the named packages.

You see `go get` a lot in READMEs. Caveat: as of 1.17 [Using go get for
installing executables
deprecated](https://golang.org/doc/go-get-install-deprecation).

These frictions come from mixed use cases with `go get` to fetch and build code
- and with the introduction of go modules, to also change the `go.mod` file
describing dependencies.

Go modules are the preferred way to start applications and libraries today.

To install executables on your system, run `go install` instead:

* without version suffix, if works on module level
* with version suffix (e.g. "@latest") is ignores modules

```
$ go install golang.org/x/tools/cmd/goimports@latest
```

## Extra: Stdlib packages

```
$ tree -d -I 'vendor|testdata|internal' /usr/local/go/src/
/usr/local/go/src/
├── archive
│   ├── tar
│   └── zip
├── bufio
├── builtin
├── bytes
├── cmd
│   ├── addr2line
│   ├── api
│   ├── asm
│   ├── buildid
│   ├── cgo
│   ├── compile
│   ├── cover
│   ├── dist
│   ├── doc
│   ├── fix
│   ├── go
│   ├── gofmt
│   ├── link
│   ├── nm
│   ├── objdump
│   ├── pack
│   ├── pprof
│   ├── test2json
│   ├── trace
│   └── vet
├── compress
│   ├── bzip2
│   ├── flate
│   ├── gzip
│   ├── lzw
│   └── zlib
├── container
│   ├── heap
│   ├── list
│   └── ring
├── context
├── crypto
│   ├── aes
│   ├── cipher
│   ├── des
│   ├── dsa
│   ├── ecdsa
│   ├── ed25519
│   ├── elliptic
│   ├── hmac
│   ├── md5
│   ├── rand
│   ├── rc4
│   ├── rsa
│   ├── sha1
│   ├── sha256
│   ├── sha512
│   ├── subtle
│   ├── tls
│   └── x509
│       └── pkix
├── database
│   └── sql
│       └── driver
├── debug
│   ├── dwarf
│   ├── elf
│   ├── gosym
│   ├── macho
│   ├── pe
│   └── plan9obj
├── embed
├── encoding
│   ├── ascii85
│   ├── asn1
│   ├── base32
│   ├── base64
│   ├── binary
│   ├── csv
│   ├── gob
│   ├── hex
│   ├── json
│   ├── pem
│   └── xml
├── errors
├── expvar
├── flag
├── fmt
├── go
│   ├── ast
│   ├── build
│   │   └── constraint
│   ├── constant
│   ├── doc
│   ├── format
│   ├── importer
│   ├── parser
│   ├── printer
│   ├── scanner
│   ├── token
│   └── types
│       └── fixedbugs
├── hash
│   ├── adler32
│   ├── crc32
│   ├── crc64
│   ├── fnv
│   └── maphash
├── html
│   └── template
├── image
│   ├── color
│   │   └── palette
│   ├── draw
│   ├── gif
│   ├── jpeg
│   └── png
├── index
│   └── suffixarray
├── io
│   ├── fs
│   └── ioutil
├── log
│   └── syslog
├── math
│   ├── big
│   ├── bits
│   ├── cmplx
│   └── rand
├── mime
│   ├── multipart
│   └── quotedprintable
├── net
│   ├── http
│   │   ├── cgi
│   │   ├── cookiejar
│   │   ├── fcgi
│   │   ├── httptest
│   │   ├── httptrace
│   │   ├── httputil
│   │   └── pprof
│   ├── mail
│   ├── rpc
│   │   └── jsonrpc
│   ├── smtp
│   ├── textproto
│   └── url
├── os
│   ├── exec
│   ├── signal
│   └── user
├── path
│   └── filepath
├── plugin
├── reflect
├── regexp
│   └── syntax
├── runtime
│   ├── cgo
│   ├── debug
│   ├── metrics
│   ├── msan
│   ├── pprof
│   ├── race
│   └── trace
├── sort
├── strconv
├── strings
├── sync
│   └── atomic
├── syscall
│   └── js
├── testing
│   ├── fstest
│   ├── iotest
│   └── quick
├── text
│   ├── scanner
│   ├── tabwriter
│   └── template
│       └── parse
├── time
│   └── tzdata
├── unicode
│   ├── utf16
│   └── utf8
└── unsafe

184 directories

```
