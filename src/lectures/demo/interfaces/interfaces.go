package main

import "fmt"

type DishPreparer interface {
	// interface usually gets -er suffix
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("roasting serving in the oven")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop vegetables")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []DishPreparer) {
	fmt.Println("Prepare dishes:")

	for _, dish := range dishes {
		fmt.Printf("--Dish: %v --\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []DishPreparer{Chicken("Chicken wings"), Salad("Caesar salad")}
	prepareDishes(dishes)
}
