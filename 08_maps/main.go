package main

import "fmt"

func main() {
	// Define a map
	// emails := make(map[string]string) // make(map[type_of_key]type_of_value)

	// Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
