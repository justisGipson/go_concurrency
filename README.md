# Concurrency is hard.


> “Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.” - *Rob Pike*

Concurrency has a lot to do with the design of your program.

When design comes to the table, it is a good idea to reuse patterns that worked for previous problems to avoid reinventing the wheel.

There are 3 different examples of concurrency patterns in Go.

# tl/dr:
  - [`sync.atomic`](https://golang.org/pkg/sync/atomic): Useful when making operations with integers.
  - [`sync.Mutex`](https://golang.org/pkg/sync/#Mutex) and [`sync.RWMutex`](https://golang.org/pkg/sync/#RWMutex): For synchronizing the access to more complex data structures. It’s the classical approach and allows for custom locking.
  - channels: When mutexes are not an option or it is complicated to operate with them.
