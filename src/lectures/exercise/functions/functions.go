//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func getSampleMessage() string {
	return "Here is a random message!"
}

func addThree(one, two, three int) int {
	return one + two + three
}

func getSampleNumber() int {
	return 319
}

func getSampleNumbers() (int, int) {
	return 115, 520
}

func main() {
	greet("Aanika")
	fmt.Println(getSampleMessage())
	one := getSampleNumber()
	two, three := getSampleNumbers()
	result := addThree(one+one, two, three)
	fmt.Println("result is", result)
}
