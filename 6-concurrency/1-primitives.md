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

* lightweight threads

## Channels

* typed conduits
* unidirectional, bidirectional

## Select
