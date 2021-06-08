# Primitives

Go supports classic and CSP style concurrency.

## When to use what?

![](https://raw.githubusercontent.com/miku/cignotes/master/images/fig21.png)

From: [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) (2017)

Sidenote:

Example of bare bones currency control, implementing Solution of a "Problem in
Concurrent Programming Control" (Dijkstra, 1965) in Go:

* https://gist.github.com/miku/18d13ab0c7de09120a26f8ebe153ad27#file-minimal-go-L17-L47

## Goroutines

* lightweight threads (without id and less external control)
* M:N scheduler: the runtime distributes M goroutines onto N OS threads (see `GOMAXPROCS`
  environment variable, defaults to the value of `runtime.NumCPU`)

> By default, Go programs run with GOMAXPROCS set to the number of cores
> available; in prior releases it defaulted to 1. -- [Go
> 1.5](https://golang.org/doc/go1.5), [Design Doc](https://docs.google.com/document/d/1At2Ls5_fhJQ59kDK2DFVhFu3g5mATSXqqV5QrxinasI/edit)

* every program runs at least one goroutine

[embedmd]:# (../x/goroutine/main.go)
```go
package main

import (
	"log"
	"runtime"
	"time"
)

func hello() {
	log.Println("hello")
}

func main() {
	log.Printf("cpus=%d", runtime.NumCPU())
	log.Printf("goroutines=%d", runtime.NumGoroutine())

	go hello()

	log.Printf("goroutines=%d", runtime.NumGoroutine())
	time.Sleep(10 * time.Millisecond)
}
```

They are lightweight, so you can start many of them.

[embedmd]:# (../x/manygoroutines/main.go)
```go
package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

var (
	N     = flag.Int("n", 10, "number of goroutines to start")
	sleep = flag.Duration("s", 100*time.Millisecond, "how long each goroutine sleeps")
)

func main() {
	flag.Parse()
	log.Printf("starting %d goroutines", *N)
	started := time.Now()
	for i := 0; i < *N; i++ {
		go f()
	}
	log.Printf("%d/%d goroutines started/running in %v", *N, runtime.NumGoroutine(), time.Since(started))
}

func f() {
	time.Sleep(*sleep)
}
```

## Channels

* typed conduits
* unidirectional, bidirectional

## Select
