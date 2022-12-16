//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rolls, dice, sides := 3, 2, 6
	var sum int
	for r := 0; r < rolls; r++ {
		for d := 0; d < dice; d++ {
			eye := rand.Intn(sides) + 1
			fmt.Println("roll:", r, "dice:", d, "eye:", eye)
			sum += eye
		}
	}

	fmt.Println("sum =", sum)

	if (dice == 2) && (sum == 2) {
		fmt.Println("Snake Eyes")
	} else if sum == 7 {
		fmt.Println("Lucky 7")
	}
	parity := sum % 2
	switch parity {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}
