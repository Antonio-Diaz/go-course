package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook Chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop Salad")
	fmt.Println("add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Preparing dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("Preparing dish %v\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chincken Wings"), Salad("Caesar Salad")}
	PrepareDishes(dishes)
}
