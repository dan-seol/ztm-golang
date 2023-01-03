//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	countNonBlankLines := 0
	countCommands := 0

	for {
		input, inputErr := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "q" {
			break
		}

		splitInput := strings.Split(input, " ")
		for _, command := range splitInput {
			switch strings.ToLower(command) {
			case "hello", "hi":
				fmt.Println("Hello, user")
				countCommands++
			case "bye":
				fmt.Println("Bye, user")
				countCommands++
			}

		}
		countNonBlankLines++
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Println("Count of non-blank lines:", countNonBlankLines)
	fmt.Println("Count of commands:", countCommands)
}
