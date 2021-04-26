package main

import (
	"fmt"
	"strings"
)

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

type UserErrorString string

const (
	IncorrectNameString UserErrorString = "Bad username, try again"
	IncorrectPasswordString UserErrorString = "Wrong password, did you forget it?"
	InvalidEmailString UserErrorString = "Invalid email address, should be user@domain.com"
)

func (err UserErrorString) String() string {
	errors := []string{"Bad username, try again",
	"Wrong password, did you forget it?",
	"Invalid email address, should be user@domain.com",
	}

	x := string(err)
	for _, v := range errors {
		if v == x {
			return x
		}
	}
	return ""
}

func main() {
	var myErr UserError
	myErrString := "Invalid email address, should be user@domain.com"
	if strings.Contains(string(myErrString), "address") {
		fmt.Println("It is an email error!")
	}
	fmt.Println(myErr) // prints "Invalid email address, should be user@domain.com"
}