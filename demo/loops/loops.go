package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum of integers 1 to 10:")
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is ", i)
	}
}
