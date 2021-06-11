# Deployments

## Building

```
$ go build -o file.exe file.go
```

## Cross-Compilation

* using GOOS and GOARCH environment variables

```
$ GOOS=linux GOARCH=arm64 go build main.go
```

* list of available options: https://github.com/golang/go/blob/master/src/go/build/syslist.go


## Smaller binaries

* use `go build -ldflags="-s -w" main.go` to omit symbol tables
* use [upx](https://upx.github.io/), to compress binaries 

## Linux Images from Scratch

Building small Go docker images.

* [fromscratch](fromscratch)

