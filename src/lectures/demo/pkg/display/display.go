package display

import "fmt"

// the function Display will be exported as its name starts with an uppercase letter
func Display(message string) {
	fmt.Println(message)
}

//the function hello will not be exported as its name starts with a lowercase letter

func hello(message string) {
	fmt.Println("hello,", message)
}
