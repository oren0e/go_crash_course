package main

import "fmt"

func foo() {
        defer fmt.Println("Done!") // this will be evaluated immediately but only run when the rest of foo() is done.
        defer fmt.Println("Are we done?")  // this will be printed before the above defer
        fmt.Println("Doing some stuff")
        // Generally, we can defer some sort of function that will handle an error once a program panics,
        // the program does not have to stop running in order to handle this.
        // The "defer" statements are going to run even if the function is not successfull, like when 
        // panic occures.
}

func main() {
        foo()
}
