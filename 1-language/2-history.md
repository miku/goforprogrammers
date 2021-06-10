# History

* First presentation in 10 Nov, 2009 ([Google Tech Talk](https://www.youtube.com/watch?v=rKnDgT73v8s))
* [Go 1](https://blog.golang.org/go1),  28 March 2012
* [Go 1.5](https://golang.org/doc/go1.5#c), no more C

Release history:
[https://golang.org/doc/devel/release](https://golang.org/doc/devel/release),
latest as of 06/2021: [go1.16.5](https://golang.org/doc/devel/release#go1.16).


## Repo History

```
$ git summary | head -30

 project  : go
 repo age : 49 years
 active   : 4479 days
 commits  : 48214
 files    : 10118
 authors  :
  6852  Russ Cox                           14.2%
  3348  Robert Griesemer                   6.9%
  2977  Rob Pike                           6.2%
  2341  Brad Fitzpatrick                   4.9%
  2125  Ian Lance Taylor                   4.4%
  1482  Austin Clements                    3.1%
  1403  Josh Bleecher Snyder               2.9%
  1190  Andrew Gerrand                     2.5%
  1170  Keith Randall                      2.4%
  1112  Matthew Dempsky                    2.3%
  1024  Cherry Zhang                       2.1%
   922  Shenghou Ma                        1.9%
   771  Alex Brainman                      1.6%
   752  Mikio Hara                         1.6%
   710  Bryan C. Mills                     1.5%
   690  Dmitriy Vyukov                     1.4%
   523  Adam Langley                       1.1%
   508  Ken Thompson                       1.1%
   476  Dave Cheney                        1.0%
   455  Nigel Tao                          0.9%
   392  David Crawshaw                     0.8%
   381  Joel Sing                          0.8%
   365  David Chase                        0.8%
```

Tree at first commit:

```
$ tree
.
└── src
    └── pkg
        └── debug
            └── macho
                └── testdata
                    └── hello.b

5 directories, 1 file

$ cat src/pkg/debug/macho/testdata/hello.b
main( ) {
        extrn a, b, c;
        putchar(a); putchar(b); putchar(c); putchar('!*n');
}
a 'hell';
b 'o, w';
c 'orld';
```

# Go: A timeline of events

* [The Go programming language](https://www.youtube.com/watch?v=rKnDgT73v8s), November 2009
* [Go version 1 is released](https://blog.golang.org/go-version-1-is-released), 28 March 2012

The Go 1 release marked an important milestone, which is codified in [Go 1 and
the Future of Go Programs](https://golang.org/doc/go1compat):

> Go 1 defines two things: first, the specification of the language; and
> second, the specification of a set of core APIs, the "standard packages" of
> the Go library. The Go 1 release includes their implementation in the form
> of two compiler suites (gc and gccgo), and the core libraries themselves.

> It is intended that programs written to the Go 1 specification will continue
> to compile and run correctly, unchanged, over the lifetime of that
> specification.

In Spring 2016, Brad Fitzpatrick gave a talk: [Go 1.6: Asymptotically
approaching boring](https://www.youtube.com/watch?v=4Dr8FXs9aJM)
([slides](https://docs.google.com/presentation/d/1JsCKdK_AvDdn8EkummMNvpo7ntqteWQfynq9hFTCkhQ/view?slide=id.p)).

Go has releases approximately every six months (seems attractive to
[others](https://www.infoq.com/news/2017/09/Java6Month), too).

Rationale: Stable foundation - to build stuff on top.

* Language ([ref/spec](https://golang.org/ref/spec))
* Standard Library ([pkg](https://golang.org/pkg/))
* Runtime ([GC](https://www.youtube.com/watch?v=aiv1JOfMjm0), scheduler, and
  other pieces under active development)
* Tools (go, godoc, go vet, gofmt, goimports, ... under active development)
* Ecosystem (external packages, conferences, user groups, and much more)

The current release is [Go 1.16](https://blog.golang.org/go1.16) from 16 February
2021.

# Origins and influences

A multitude of ideas and influences:

* C++ might be slow to compile and bloated
* all programming languages seem to add and add features
* we entered the multicore era ([Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm))
* design for a networked world
* bridge the gap between dynamic and static programming languages (be safe, yet
  ease to write)
* focus on long-term maintenance
* designed with tools in mind (gofmt being the prototypical tool)
* a different approach to concurrency
* a stripped version of object orientation

The original designers are Robert Griesemer, Rob Pike, Ken Thomson.
