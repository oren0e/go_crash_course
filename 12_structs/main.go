package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver method = methods that don't change anything)
func (p Person) greet() string { // func (x strcut_of_the_method), x is like "self"
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer receiver = methods that actually change something)
func (p *Person) hasBirthday() { // does not return anything, just changes things
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	// person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F",
	//	age: 25}
	// Alternative
	person1 := Person{"Samantha", "Smith", "Boston", "F", 25}
	person2 := Person{"Bob", "Johnson", "New York", "M", 30}

	// fmt.Println(person1)

	// Get single field
	// fmt.Println(person1.age)
	// person1.age++
	// fmt.Println(person1.age)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
