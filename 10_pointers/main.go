package main

import "fmt"

func main() {
	a := 5
	b := &a // assign b to a pointer to a (will display a memory address where a is stored)

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // we get *int, which representes a pointer.

	// Use * to read value from address (like dereference in Rust)
	fmt.Println(*b)
	fmt.Println(*&a) // same thing

	// Change value with pointer
	*b = 10 // b is a pointer to the value of a, then we change the value with * (which dereferences it)
	fmt.Println(a)

	// Everything in Go is passed by value and not by reference! To pass by reference you have to do it explicitly.
}
