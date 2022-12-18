package main

import "fmt"

func main() {
	/*
		myMap := make(map[string]string)
		myMap["favourite"] = "Aanika"
		aanika := myMap["favourite"]
		missing := myMap["missing"] // will give default value
		delete(myMap, "favourite")
		price, found := myMap["priceLabel"]
		if !found {
			fmt.Println("price not found")
			return
		}
	*/
	/*
		myMap := map[string]int{
			"item1": 1,
			"item2": 2,
			"item3": 3
		}
		for key, value := range myMap {

		}
	*/
	shoppingList := make(map[string]int)
	shoppingList["bacon"] = 11
	shoppingList["ginger ale"] = 3
	shoppingList["rice"] += 20
	shoppingList["bacon"] += 1
	fmt.Println(shoppingList)
	delete(shoppingList, "rice")
	fmt.Println("Milk deleted, new list:", shoppingList)
	cereal, found := shoppingList["cereal"]
	fmt.Println("Got cereal?")
	if !found {
		fmt.Println("No, we need to")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}
	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("The total number of items is", totalItems)
}
