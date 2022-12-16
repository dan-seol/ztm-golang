package main

import "fmt"

func main() {

	var myName = "Dan"
	fmt.Println("My name is", myName, myName)

	lastName := "Aanika"
	fmt.Println("name =", lastName)

	var userName string = "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	first, second := 1, 5
	fmt.Println("First is", first, "Second is", second)

	second, third := 2, 3
	fmt.Println("Second is", second, "Third is", third)

	sum = first + second
	fmt.Println("sum =", sum)

	var (
		lessonTopic = "Variables"
		lessonType  = "Demo"
	)

	fmt.Println("lessonTopic=", lessonTopic)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "aanika", "!"
	fmt.Println(word1, word2)

}
