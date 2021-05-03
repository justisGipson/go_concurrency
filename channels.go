// Channels for goroutines. When sending messages to channels, only
// one goroutine will receive it so it is safe to access data from there
// and no explicit synchronization is needed since it is handled by
// Go under the hood
//
// Following that approach, we will have a goroutine that keeps the state
// of the counter. On the other side, other goroutines will send the messages
// that the first one will receive for interacting with the state.
// We have 2 types of messages:
// 		`incrementOP`: operation that requests incrementing the counter
// 		`valueOP`: operation that requests the value of the counter
//
// NOTE: all the operations have a `res` channel with the following
// purposes:
// 		- Receiving the operation's response
// 		- Synchronizing goroutines
// This is not based on sharing memory like atomic.go or mutex.go, but
// instead relies on sending operations through channels

package main

import (
	"log"
	"sync"
)

type op struct {
	res chan int
}

type incrementOP struct {
	op
}

type getValueOP struct {
	op
}

func newIncrementOP() incrementOP {
	return incrementOP{
		op: op{
			res: make(chan int),
		},
	}
}

func newGetValueOP() getValueOP {
	return getValueOP{
		op: op{
			res: make(chan int),
		},
	}
}

func increment(ops chan<- incrementOP, wg *sync.WaitGroup) {
	defer wg.Done()

	op := newIncrementOP()
	ops <- op
	<-op.res
}

func main() {
	incrementOps := make(chan incrementOP)
	getValueOps := make(chan getValueOP)

	go func() {
		counter := 0
		for {
			select {
			case op := <-incrementOps:
				counter++
				op.res <- counter
			case op := <-getValueOps:
				op.res <- counter
			}
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(incrementOps, &wg)
		go increment(incrementOps, &wg)
	}

	wg.Wait()

	getValueOP := newGetValueOP()
	getValueOps <- getValueOP

	log.Printf("Counter: %d", <-getValueOP.res)
}
