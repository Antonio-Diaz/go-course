package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < FirstClass:
		fmt.Println("Cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive item")
	}

	switch ticket := Economy; {
	case ticket == Economy:
		fmt.Println("Economy ticket")
	case ticket == Business:
		fmt.Println("Business ticket")
	case ticket == FirstClass:
		fmt.Println("First class ticket")
	default:
		fmt.Println("Unknown ticket")
	}
}
