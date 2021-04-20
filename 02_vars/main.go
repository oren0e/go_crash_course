package main

import "fmt"

func main() {
	var name = "Owen"
	var age = 37 // if you create a variable you have to use it
	// var age int32 = 37 // cast to another type
	const isCool = true
	only_in_function := "Owen" // shorthand for creating vars

	size := 1.3
	//email := "owen@example.com"

	name, email := "Owen", "owen@example.com"

	fmt.Println(name, age, isCool, only_in_function, email)
	fmt.Printf("%T\n", name) // print the type of "name"
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

}
