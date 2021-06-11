# Cross compilation

Go supports cross compilation out of the box. There are two relevant environment variables:

* GOOS
* GOARCH

Example creating an ARM 32 bit Linux executable on an x86_64 Linux.

```
$ uname -sp
Linux x86_64

$ env GOOS=linux GOARCH=arm go build main.go
$ file main
main:    ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, not stripped
```

To list all supported platforms (44 currently), run:

```shell
$ go tool dist list
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/386
darwin/amd64
darwin/arm
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
illumos/amd64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/s390x
nacl/386
nacl/amd64p32
nacl/arm
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
windows/386
windows/amd64
windows/arm
```

Can you spot the outlier? Yes, [Web Assembly](https://webassembly.org/) is supported as well.
