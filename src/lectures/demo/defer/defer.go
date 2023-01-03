package main

import "fmt"

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

func three() {
	fmt.Println("3")
}

func sample() {
	fmt.Println("Begin")

	defer one()
	defer two()
	defer three()

	fmt.Println("End")
}

func main() {
	sample()
}
