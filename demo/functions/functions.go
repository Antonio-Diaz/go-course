package main

import "fmt"

func double(x int) int {
	return x * x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Greets from functions.go")
}

func main() {
	greet()

	dozen := double(12)
	fmt.Println("12 squared is", dozen)

	sum := add(dozen, 1)
	fmt.Println("12 squared + 1 is", sum)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("6 squared + 1 is", anotherBakersDozen)
}
