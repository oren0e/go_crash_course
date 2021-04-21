package main

import (
	"fmt"
	"sync"
)

/*
Instead of using expensive OS threads, in go, we will create an abstraction of a thread and we're going to
call it a go routine. Inside the go runtime we're going to have a scheduler that is going to map these go-routines
onto the OS threads for periods of time. And this scheduler will then take turns with every CPU thread that's available
and assign the different go routines a certain amount of processing time on those threads. We don't have to interact
with those low-level threads directly, because we have the go-routines.
The advantage of this abstraction is that go-routines can start with very small stack spaces because they can be realocated
very quickly, and thus they are very cheap to create and to destroy.
This abstraction lets us to use threads without the fear of creating or destroying them.
*/

var wg = sync.WaitGroup{}

// waitgroup is designed to synchronize multiple go-routines together.
// Here we have 2 go-routines: the one that executes the main function and the go-routine
// we're spawning off with the msg. We want to synchronize the main function to the go-routine.

func main() {
	// go sayHello() // "go" turns a function into a go routine
	// time.Sleep(100 * time.Millisecond) // if we add this we will see the message printed. If this wasn't here,
	// the message was "printed" inside the go-routine, once the main function ends
	// (when it called the go routine) the program terminates and we would not see the message
	// printed. Now, because of the artificial delay, we delay the main function a little bit
	// which gives us enough time for the go-routine to spwan off from the main function and that's
	// what is responsible for printing the message.
	// we can also do it like this with a closure.

	// It is generally not advised to access a varaible that's defined in the outer scope of a function
	// inside the function, especially with go-routines which can lead to race conditions like so:
	var msg = "Hello"
	wg.Add(1) // starts off synchronizing to zero, so we add 1 telling it that we want to
	// synchronize to this go-routine right here below.
	go func(msg string) {
		fmt.Println(msg)
		// when the go-routine is done it can tell the waitgroup that it's actually completed
		// its execution, by the .Done() method. (will decrement the number of groups the waiting group
		// is waiting on, since we added 1, it will be back at 0, and then the .Wait() method will know that
		// it's time to finish our application run)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
	// time.Sleep(100 * time.Millisecond)
	// Most of the time we will see "Goodbye" printed here, because the go scheduler will not interupt
	// the main thread until it hits the sleep call. Even though it launches another go-routine at the start
	// it didn't give it any "love" yet, it still is executing the main function. So it reassigns the msg variable
	// before the go-routine had the chance to print it out. This actually creates a "race condition".
}

func sayHello() {
	fmt.Println("Hello")
}
