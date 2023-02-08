package main

import (
	"fmt"
	"math/rand"
	"time"
)

//generator giving random numbers

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randInt := generateRandInt(1, 100)

	fmt.Println("generateRandInt infinite")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("generateRandIntn using finite count")
	randIntnRange := generateRandIntn(2, 1, 10)
	for i := range randIntnRange {
		fmt.Println(i)
	}

	fmt.Println(<-randIntnRange)
	fmt.Println("generateRandIntn using infinite loop")

	randIntn := generateRandIntn(3, 1, 10)
	for {
		n, open := <-randIntn
		if !open {
			break
		}
		fmt.Println(n)
	}

}
