//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Greetings", name, "!")
}

func message() string {
	return "Hello from functions.go"
}

func addThree(lhs, mhs, rhs int) int {
	return lhs + mhs + rhs
}

func six() int {
	return 6
}

func twoNumbers() (int, int) {
	return 2, 2
}

func main() {
	greeting("Bob")
	fmt.Println(message())

	a, b := twoNumbers()
	answer := addThree(a, b, six())

	fmt.Println("The answer is", answer)
}
