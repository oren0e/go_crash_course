package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { // can also do getSum(num1, num2 int) int {}
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Owen"))
	fmt.Println(getSum(3, 4))
}
