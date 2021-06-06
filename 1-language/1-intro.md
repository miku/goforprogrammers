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

At import time, Go first look for packages in `$GOROOT/src`, which contains the standard library:

```
$ tree -d -I vendor /usr/local/go/src/ | pbcopy
/usr/local/go/src/
├── archive
│   ├── tar
│   │   └── testdata
│   └── zip
│       └── testdata
├── bufio
├── builtin
├── bytes
├── cmd
│   ├── addr2line
│   ├── api
│   │   └── testdata
│   │       └── src
│   │           ├── issue21181
│   │           │   ├── dep
│   │           │   ├── indirect
│   │           │   └── p
│   │           ├── issue29837
│   │           │   └── p
│   │           └── pkg
│   │               ├── p1
│   │               ├── p2
│   │               └── p3
│   ├── asm
│   │   └── internal
│   │       ├── arch
│   │       ├── asm
│   │       │   └── testdata
│   │       │       └── avx512enc
│   │       ├── flags
│   │       └── lex
│   ├── buildid
│   ├── cgo
│   ├── compile
│   │   └── internal
│   │       ├── amd64
│   │       ├── arm
│   │       ├── arm64
│   │       ├── gc
│   │       │   ├── builtin
│   │       │   └── testdata
│   │       │       ├── gen
│   │       │       └── reproducible
│   │       ├── logopt
│   │       ├── mips
│   │       ├── mips64
│   │       ├── ppc64
│   │       ├── riscv64
│   │       ├── s390x
│   │       ├── ssa
│   │       │   ├── gen
│   │       │   └── testdata
│   │       ├── syntax
│   │       │   └── testdata
│   │       ├── test
│   │       ├── types
│   │       ├── wasm
│   │       └── x86
│   ├── cover
│   │   └── testdata
│   │       └── html
│   ├── dist
│   ├── doc
│   │   └── testdata
│   │       ├── merge
│   │       └── nested
│   │           ├── empty
│   │           └── nested
│   ├── fix
│   ├── go
│   │   ├── internal
│   │   │   ├── auth
│   │   │   ├── base
│   │   │   ├── bug
│   │   │   ├── cache
│   │   │   ├── cfg
│   │   │   ├── clean
│   │   │   ├── cmdflag
│   │   │   ├── doc
│   │   │   ├── envcmd
│   │   │   ├── fix
│   │   │   ├── fmtcmd
│   │   │   ├── fsys
│   │   │   ├── generate
│   │   │   ├── get
│   │   │   ├── help
│   │   │   ├── imports
│   │   │   │   └── testdata
│   │   │   │       ├── android
│   │   │   │       ├── illumos
│   │   │   │       └── star
│   │   │   ├── list
│   │   │   ├── load
│   │   │   ├── lockedfile
│   │   │   │   └── internal
│   │   │   │       └── filelock
│   │   │   ├── modcmd
│   │   │   ├── modconv
│   │   │   │   └── testdata
│   │   │   ├── modfetch
│   │   │   │   ├── codehost
│   │   │   │   └── zip_sum_test
│   │   │   │       └── testdata
│   │   │   ├── modget
│   │   │   ├── modinfo
│   │   │   ├── modload
│   │   │   ├── mvs
│   │   │   ├── par
│   │   │   ├── renameio
│   │   │   ├── robustio
│   │   │   ├── run
│   │   │   ├── search
│   │   │   ├── str
│   │   │   ├── test
│   │   │   ├── tool
│   │   │   ├── trace
│   │   │   ├── txtar
│   │   │   ├── vcs
│   │   │   ├── version
│   │   │   ├── vet
│   │   │   ├── web
│   │   │   └── work
│   │   └── testdata
│   │       ├── failssh
│   │       ├── mod
│   │       ├── modlegacy
│   │       │   └── src
│   │       │       ├── new
│   │       │       │   ├── p1
│   │       │       │   ├── p2
│   │       │       │   └── sub
│   │       │       │       ├── inner
│   │       │       │       │   └── x
│   │       │       │       └── x
│   │       │       │           └── v1
│   │       │       │               └── y
│   │       │       └── old
│   │       │           ├── p1
│   │       │           └── p2
│   │       ├── script
│   │       └── testterminal18153
│   ├── gofmt
│   │   └── testdata
│   ├── internal
│   │   ├── archive
│   │   │   └── testdata
│   │   │       └── mycgo
│   │   ├── bio
│   │   ├── browser
│   │   ├── buildid
│   │   │   └── testdata
│   │   ├── codesign
│   │   ├── diff
│   │   ├── dwarf
│   │   ├── edit
│   │   ├── gcprog
│   │   ├── goobj
│   │   ├── moddeps
│   │   ├── obj
│   │   │   ├── arm
│   │   │   ├── arm64
│   │   │   ├── mips
│   │   │   ├── ppc64
│   │   │   ├── riscv
│   │   │   │   └── testdata
│   │   │   │       └── testbranch
│   │   │   ├── s390x
│   │   │   ├── wasm
│   │   │   └── x86
│   │   ├── objabi
│   │   ├── objfile
│   │   ├── pkgpath
│   │   ├── src
│   │   ├── sys
│   │   ├── test2json
│   │   │   └── testdata
│   │   └── traceviewer
│   ├── link
│   │   ├── internal
│   │   │   ├── amd64
│   │   │   ├── arm
│   │   │   ├── arm64
│   │   │   ├── benchmark
│   │   │   ├── ld
│   │   │   │   └── testdata
│   │   │   │       ├── deadcode
│   │   │   │       ├── httptest
│   │   │   │       │   └── main
│   │   │   │       ├── issue10978
│   │   │   │       ├── issue25459
│   │   │   │       │   ├── a
│   │   │   │       │   └── main
│   │   │   │       ├── issue26237
│   │   │   │       │   ├── b.dir
│   │   │   │       │   └── main
│   │   │   │       ├── issue32233
│   │   │   │       │   ├── lib
│   │   │   │       │   └── main
│   │   │   │       ├── issue38192
│   │   │   │       ├── issue39256
│   │   │   │       └── issue39757
│   │   │   ├── loadelf
│   │   │   ├── loader
│   │   │   ├── loadmacho
│   │   │   ├── loadpe
│   │   │   ├── loadxcoff
│   │   │   ├── mips
│   │   │   ├── mips64
│   │   │   ├── ppc64
│   │   │   ├── riscv64
│   │   │   ├── s390x
│   │   │   ├── sym
│   │   │   ├── wasm
│   │   │   └── x86
│   │   └── testdata
│   │       ├── testBuildFortvOS
│   │       ├── testHashedSyms
│   │       ├── testIndexMismatch
│   │       ├── testPErsrc
│   │       ├── testPErsrc-complex
│   │       └── testRO
│   ├── nm
│   ├── objdump
│   │   └── testdata
│   │       └── testfilenum
│   ├── pack
│   ├── pprof
│   ├── test2json
│   ├── trace
│   └── vet
│       └── testdata
│           ├── asm
│           ├── assign
│           ├── atomic
│           ├── bool
│           ├── buildtag
│           ├── cgo
│           ├── composite
│           ├── copylock
│           ├── deadcode
│           ├── httpresponse
│           ├── lostcancel
│           ├── method
│           ├── nilfunc
│           ├── print
│           ├── rangeloop
│           ├── shift
│           ├── structtag
│           ├── tagtest
│           ├── testingpkg
│           ├── unmarshal
│           ├── unsafeptr
│           └── unused
├── compress
│   ├── bzip2
│   │   └── testdata
│   ├── flate
│   │   └── testdata
│   ├── gzip
│   │   └── testdata
│   ├── lzw
│   ├── testdata
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
│   │   └── testdata
│   ├── ed25519
│   │   ├── internal
│   │   │   └── edwards25519
│   │   └── testdata
│   ├── elliptic
│   ├── hmac
│   ├── internal
│   │   ├── randutil
│   │   └── subtle
│   ├── md5
│   ├── rand
│   ├── rc4
│   ├── rsa
│   │   └── testdata
│   ├── sha1
│   ├── sha256
│   ├── sha512
│   ├── subtle
│   ├── tls
│   │   └── testdata
│   └── x509
│       ├── internal
│       │   └── macos
│       ├── pkix
│       └── testdata
├── database
│   └── sql
│       └── driver
├── debug
│   ├── dwarf
│   │   └── testdata
│   ├── elf
│   │   └── testdata
│   ├── gosym
│   │   └── testdata
│   ├── macho
│   │   └── testdata
│   ├── pe
│   │   └── testdata
│   └── plan9obj
│       └── testdata
├── embed
│   └── internal
│       └── embedtest
│           └── testdata
│               ├── _hidden
│               └── i
│                   └── j
│                       └── k
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
│   │   └── testdata
│   ├── pem
│   └── xml
├── errors
├── expvar
├── flag
├── fmt
├── go
│   ├── ast
│   ├── build
│   │   ├── constraint
│   │   └── testdata
│   │       ├── cgo_disabled
│   │       ├── doc
│   │       ├── empty
│   │       ├── multi
│   │       ├── other
│   │       │   └── file
│   │       └── withvendor
│   │           └── src
│   │               └── a
│   │                   └── b
│   ├── constant
│   ├── doc
│   │   └── testdata
│   ├── format
│   ├── importer
│   ├── internal
│   │   ├── gccgoimporter
│   │   │   └── testdata
│   │   ├── gcimporter
│   │   │   └── testdata
│   │   │       └── versions
│   │   └── srcimporter
│   │       └── testdata
│   │           ├── issue20855
│   │           ├── issue23092
│   │           └── issue24392
│   ├── parser
│   │   └── testdata
│   │       └── issue42951
│   │           └── not_a_file.go
│   ├── printer
│   │   └── testdata
│   ├── scanner
│   ├── token
│   └── types
│       ├── fixedbugs
│       └── testdata
├── hash
│   ├── adler32
│   ├── crc32
│   ├── crc64
│   ├── fnv
│   └── maphash
├── html
│   └── template
│       └── testdata
├── image
│   ├── color
│   │   └── palette
│   ├── draw
│   ├── gif
│   ├── internal
│   │   └── imageutil
│   ├── jpeg
│   ├── png
│   │   └── testdata
│   │       └── pngsuite
│   └── testdata
├── index
│   └── suffixarray
├── internal
│   ├── bytealg
│   ├── cfg
│   ├── cpu
│   ├── execabs
│   ├── fmtsort
│   ├── goroot
│   ├── goversion
│   ├── lazyregexp
│   ├── lazytemplate
│   ├── nettrace
│   ├── obscuretestdata
│   ├── oserror
│   ├── poll
│   ├── profile
│   ├── race
│   ├── reflectlite
│   ├── singleflight
│   ├── syscall
│   │   ├── execenv
│   │   ├── unix
│   │   └── windows
│   │       ├── registry
│   │       └── sysdll
│   ├── sysinfo
│   ├── testenv
│   ├── testlog
│   ├── trace
│   │   └── testdata
│   ├── unsafeheader
│   └── xcoff
│       └── testdata
├── io
│   ├── fs
│   └── ioutil
│       └── testdata
├── log
│   └── syslog
├── math
│   ├── big
│   ├── bits
│   ├── cmplx
│   └── rand
├── mime
│   ├── multipart
│   │   └── testdata
│   ├── quotedprintable
│   └── testdata
├── net
│   ├── http
│   │   ├── cgi
│   │   │   └── testdata
│   │   ├── cookiejar
│   │   ├── fcgi
│   │   ├── httptest
│   │   ├── httptrace
│   │   ├── httputil
│   │   ├── internal
│   │   ├── pprof
│   │   └── testdata
│   ├── internal
│   │   └── socktest
│   ├── mail
│   ├── rpc
│   │   └── jsonrpc
│   ├── smtp
│   ├── testdata
│   ├── textproto
│   └── url
├── os
│   ├── exec
│   ├── signal
│   │   └── internal
│   │       └── pty
│   ├── testdata
│   │   ├── dirfs
│   │   │   └── dir
│   │   └── issue37161
│   └── user
├── path
│   └── filepath
├── plugin
├── reflect
├── regexp
│   ├── syntax
│   └── testdata
├── runtime
│   ├── cgo
│   ├── debug
│   ├── internal
│   │   ├── atomic
│   │   ├── math
│   │   └── sys
│   ├── metrics
│   ├── msan
│   ├── pprof
│   │   └── testdata
│   │       └── mappingtest
│   ├── race
│   │   └── testdata
│   ├── testdata
│   │   ├── testfaketime
│   │   ├── testprog
│   │   ├── testprogcgo
│   │   │   └── windows
│   │   ├── testprognet
│   │   ├── testwinlib
│   │   ├── testwinlibsignal
│   │   └── testwinsignal
│   └── trace
├── sort
├── strconv
│   └── testdata
├── strings
├── sync
│   └── atomic
├── syscall
│   └── js
├── testdata
├── testing
│   ├── fstest
│   ├── internal
│   │   └── testdeps
│   ├── iotest
│   └── quick
├── text
│   ├── scanner
│   ├── tabwriter
│   └── template
│       ├── parse
│       └── testdata
├── time
│   └── tzdata
├── unicode
│   ├── utf16
│   └── utf8
└── unsafe

543 directories
```
