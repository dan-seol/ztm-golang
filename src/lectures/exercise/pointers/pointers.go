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

type Product struct {
	name        string
	securityTag bool
}

func activateTag(product *Product) {
	product.securityTag = true
}

func deactivateTag(product *Product) {
	product.securityTag = false
}

func checkout(productSlice []Product) {
	for i := range productSlice {
		deactivateTag(&productSlice[i])
	}
}

func main() {
	products := []Product{
		{name: "Shovel", securityTag: true},
		{name: "Frying Pan", securityTag: true},
		{name: "Table", securityTag: true},
		{name: "Hammer", securityTag: true}}
	fmt.Println(products)
	deactivateTag(&products[1])
	fmt.Println(products)
	checkout(products)
	fmt.Println(products)
}
