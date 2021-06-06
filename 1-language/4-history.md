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