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

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	name  string
	price float64
}

/*
func fill(arr []Product, newProduct Product) []Product {
	var newArr []Product
	var positionToFill = 0
	for i := 0; i < len(arr); i++ {
		if  {
			positionToFill = i
			break
		}
		newArr[i] = arr[i]

	}

	if positionToFill == 0 {
		positionToFill = len(arr)
	}
	newArr[positionToFill] = newProduct

	return newArr
}
*/

func checkProduct(product Product) bool {
	return (product.name == "") && (product.price == 0)
}

func analyze(arr [4]Product) {

	var lastProduct Product
	var numOfItems int
	var totalPrice float64

	for i := 0; i < len(arr); i++ {
		if checkProduct(arr[i]) {
			break
		}
		lastProduct = arr[i]
		numOfItems++
		totalPrice += arr[i].price
	}
	fmt.Println("The last item is", lastProduct)
	fmt.Println("Total number of items are", numOfItems)
	fmt.Println("The total price is", totalPrice)
}

func generate(productName string) Product {
	return Product{name: productName, price: rand.Float64()}
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	var list [4]Product
	names := [3]string{"Eomma", "Aanika", "Anne"}
	for i := 0; i < 3; i++ {
		list[i] = generate(names[i])
	}
	analyze(list)
	list[3] = generate("Dan")
	analyze(list)

}
