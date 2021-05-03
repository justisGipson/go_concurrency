// Mutexes allow you to synchronize the access to more complex data structures.
// Here we will be using `sync.RWMutex`, which provides two locking modes:
//
//  Lock/Unlock: Locks/unlocks the data structure in writing mode.
// Neither writers nor readers will be able to access it.
//
// - RLock/RUnlock: Locks/unlocks the data structure in reading mode.
//
// Readers will be able to access it, but not writers. Using this, you can
// achieve better performance in scenarios with a lot of readers and few writers.
// Considering this, we will create a counter data structure that consists of
// the actual value of the counter and a mutex for synchronizing its access.
// We will also implement the increment and getValue methods to safely update
// and access the value of the counter, respectively
//
// NOTE: we need to unlock the mutexes after performing the operations.
// Otherwise, the upcoming ones will wait indefinitely and your program will eventually crash.

package main

import (
	"log"
	"sync"
)

type counter struct {
	value int
	mux   sync.RWMutex
}

func (c *counter) increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *counter) getValue() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.value
}

func increment(counter *counter, wg *sync.WaitGroup) {
	defer wg.Done()

	counter.increment()
}

func main() {
	c := counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&c, &wg)
		go increment(&c, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", c.getValue())
}
