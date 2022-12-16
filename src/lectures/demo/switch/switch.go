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
	case p < 2:
		fmt.Println("cheap")
	case p < 10:
		fmt.Println("moderate")
	default:
		fmt.Println("expensive")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First class seating")
	default:
		fmt.Println("Other seating")
	}
}
