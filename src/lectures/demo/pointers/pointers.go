package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
	fmt.Println("Counter with hits=", counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	aanika := "aanika"
	fmt.Println(hello, aanika)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, aanika)
	phrase := []string{hello, aanika}
	fmt.Println(phrase)
	replace(&phrase[1], "Go", &counter)
	fmt.Println(phrase)
}
