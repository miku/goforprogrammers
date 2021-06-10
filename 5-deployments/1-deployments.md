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