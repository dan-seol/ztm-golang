//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	add Operation = iota
	sub
	mul
	div
)

func (operation Operation) calculate(lhs, rhs int) int {
	var result int
	switch operation {
	case add:
		result = lhs + rhs
	case sub:
		result = lhs - rhs
	case mul:
		result = lhs * rhs
	case div:
		result = lhs / rhs
	}
	/*
		approach 2
		return map[Operation]func(int, int) int{
			add: func(i1, i2 int) int { return i1 + i2 },
			sub: func(i1, i2 int) int { return i1 - i2 },
			mul: func(i1, i2 int) int { return i1 * i2 },
			div: func(i1, i2 int) int { return i1 / i2 }}[operation](lhs, rhs)
	*/
	return result
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
