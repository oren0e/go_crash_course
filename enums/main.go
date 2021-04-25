package main

import "fmt"

type UserError int

const (
	IncorrectName UserError = iota
	IncorrectPassword
	InvalidEmail
)

func (err UserError) String() string {
	return []string{"Bad username, try again",
					"Wrong password, did you forget it?",
					"Invalid email address, should be user@domain.com",
					}[err]
}

func main() {
	var myErr UserError = InvalidEmail
	fmt.Println(myErr)
}