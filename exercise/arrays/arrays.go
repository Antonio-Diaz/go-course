//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func printStats(products [4]Product) {
	var totalItems int
	var totalCost float64

	for i := 0; i < len(products); i++ {
		item := products[i]
		totalCost += item.Price

		if item.Name != "" {
			totalItems++
		}
	}

	fmt.Println("Last item:", products[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", totalCost)
}

func main() {
	shoppingList := [4]Product{
		{Name: "Milk", Price: 1.99},
		{Name: "Bread", Price: 2.99},
		{Name: "Eggs", Price: 3.99},
	}
	printStats(shoppingList)
	shoppingList[3] = Product{Name: "Cheese", Price: 4.99}
	printStats(shoppingList)
}
