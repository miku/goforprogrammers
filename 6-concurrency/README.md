# Concurrency

* [Introduction](1-primitives.md)
* Go Tour, [Concurrency](https://tour.golang.org/concurrency/1)
* [CSP](csp)
* [Goroutines](goroutines)
* [Channels](channels)
* [Select](select)

## Workout

1. Start 10 Go routines, each of which counts from 0 to 3 with a small delay,
   e.g. 10 milliseconds. How can you make sure the program waits until all
   goroutines are finished?

2. Write a small "worker function" that uses a string channel as queue. Setup 3
   workers. The put 10 elements on the queue channel.  

