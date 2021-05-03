// increases counter concurrently from different
// goroutines.
//
// we call `AddInt32`, passing a pointer to the counter
// and a delta as parameters. Then to safely access the
// counter value we use `LoadInt32
// NOTE: goroutines are still running when `log.Printf` is
// executed, so we wait with `sync.WaitGroup()`
// Basically, this is split into 100 chunks and dispatched
// in batches of two that increment the counter by one
// asynchronously
// this approach is based on sharing memory – more specifcally
// a reference to an integer – and sychronizing the access to it

package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func increment(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt32(counter, 1)
}

func main() {
	counter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&counter, &wg)
		go increment(&counter, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", atomic.LoadInt32(&counter))
}
