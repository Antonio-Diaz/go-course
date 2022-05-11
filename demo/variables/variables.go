package main

import "fmt"

func main() {

	// delcare a variable without a type
	var myName = "Jose"
	fmt.Println("My name is", myName, myName)

	// declare a variable with a tyope and a value
	var name string = "Jose"
	fmt.Println("My name is", name)

	// Declare and assign
	userName := "admin"
	fmt.Println("My username is", userName)

	// Declare and initialize a variable
	var sum int
	fmt.Println("the sum is", sum)

	// Declare and assign multiple variables
	part1, other := 1, 2
	fmt.Println("part1 is", part1, "and other is", other)

	// Change the value of a multiple variables
	part2, other := 2, 5
	fmt.Println("part2 is", part2, "and other is", other)

	// assign the value of part1, part2 to sum that will be used in the next line
	sum = part1 + part2
	fmt.Println("the sum is", sum)

	// Declare a block of variables
	var (
		lessonName = "Go"
		lessonType = "Programming Language"
	)

	fmt.Println("The lesson is", lessonName, "and it is a", lessonType)

	// Declare and assign multiple variables at once and ignore the last one with underscore
	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)

}
