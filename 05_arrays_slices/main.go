package main

import "fmt"

func main() {
	// Arrays (fixed length)
	// var fruitArray [2]string

	// Assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)
	// fmt.Println(fruitArray[1]) // 0-based index

	// Slices (like arrays but with unknown length)
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2]) // inclusive:non-inclusive
}
