package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 3
	shoppingList["bread"] += 2

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk was deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")
	
	if cereal, ok := shoppingList["cereal"]; ok {
		fmt.Println("You need", cereal, "cereal")
	} else {
		fmt.Println("You don't need cereal")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
}
