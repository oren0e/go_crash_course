package main

import (
	"fmt"
	"sync"
)

// The following race condition can be fixed with a mutex, which is basically a lock
// that the application is going to honor.

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Read-Write mutex
// A simple mutex is simply locked or unlocked, so if the mutex is locked and something is trying to manipulate
// that value, it has to wait until the mutex is unlocked and it can obtain the mutex lock itself. We can protect parts
// of our code so that only one entity can be manipulating our code at a time. Typically this is done with data to protect it
// and make sure that only one thing can access it at a single time.
// With RWMutex, we say - many things can read this data but only one can write to it at a time. If anything is reading,
// then we can't write to it at all. So we can have an infinite number of readers but only one writer.

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

// try to protect the counter from reading and writing at the same time
func sayHello() {
	//m.RLock() // locks for reading.
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	//m.Lock() // locks for writing.
	counter++
	m.Unlock()
	wg.Done()
}
