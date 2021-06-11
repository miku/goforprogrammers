# Context

Added to the standard library in Go 1.7.

> Package context defines the Context type, which carries deadlines,
> cancellation signals, and other request-scoped values across API boundaries
> and between processes.

By convention, if context is used, it should be the first parameter.

> Do not store Contexts inside a struct type; instead, pass a Context explicitly
> to each function that needs it. The Context should be the first parameter,
> typically named ctx.


```go
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

There are four ways to create a context:

* WithValue
* WithDeadline
* WithTimeout
* WithCancel

All function return a pair of `ctx, cancelFunc`.

A [`CancelFunc`](https://golang.org/pkg/context/#CancelFunc) tells an operation to abandon its work.

The CancelFunc should stay where it was created.

## Top level context

Provided via `context.Background`, can be used to derive other contexts.

```go
ctx, cancel := context.Background()
```

## Placeholder

Use `context.TODO` as a placeholder.

```go
ctx, cancel := context.TODO()
```

## Remnant of previous approach

The
[`NewRequestWithContext`](https://tip.golang.org/pkg/net/http/#NewRequestWithContext)
has been added in Go 1.13.

```
    // Cancel is an optional channel whose closure indicates that the client
    // request should be regarded as canceled. Not all implementations of
    // RoundTripper may support Cancel.
    //
    // For server requests, this field is not applicable.
    //
    // Deprecated: Set the Request's context with NewRequestWithContext
    // instead. If a Request's Cancel field and context are both
    // set, it is undefined whether Cancel is respected.
    Cancel <-chan struct{} // Go 1.5
```

## Code Review

* [Example 1](example1/main.go)
* [Example 2](example2/main.go)
* [Example 3](example3/main.go)
* [Example 4](example4/main.go)
* [Example 5](example5/main.go)
