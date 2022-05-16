//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

//* Create a structure to store items and their security tag state
type Item struct {
	Name string
	Tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func changeToActivate(tag *SecurityTag) {
	*tag = Active
}
func changeToDeactivate(tag *SecurityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("Checking out items...")
	for i := range items {
		changeToDeactivate(&items[i].Tag)
	}
}

func main() {

	//* Create at least 4 items, all with active security tags
	item1 := Item{"Item 1", true}
	item2 := Item{"Item 2", true}
	item3 := Item{"Item 3", true}
	item4 := Item{"Item 4", true}
	//* Store them in a slice or array
	items := []Item{item1, item2, item3, item4}
	fmt.Println(items)
	//* Deactivate any one security tag in the array/slice
	changeToDeactivate(&items[0].Tag)
	fmt.Println(items)
	//* Call the checkout() function to deactivate all tags
	checkout(items)
	//* Print out the array/slice after each change
	fmt.Println(items)
}
