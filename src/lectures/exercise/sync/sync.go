//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

func countWord(wg *sync.WaitGroup, result *int, word string) {
	letterCount := 0
	for _, r := range word {
		if unicode.IsLetter(r) {
			letterCount++
		}
	}
	defer wg.Done()
	(*result) += letterCount
}

func main() {
	var wg sync.WaitGroup
	//channel := make(chan int, 200)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	letterCount := 0

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			wg.Add(1)
			go countWord(&wg, &letterCount, word)
		}
	}
	wg.Wait()
	fmt.Println("result:", letterCount)
}

/*
type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	letters := 0
	for _, r := range word {
		if unicode.IsLetter(r) {
			letters++
		}
	}
}

func main() {
	totalLetters := Count{}
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		words := getWords(scanner.Text())
		for _, word;= range words {
			wordCopy := word
			wg.Add(1)
			go func() {
				totalLetters.Lock()
				defer totalLetters.Unlock()
				defer wg.Done()
				sum := countLetters(wordcopy)
				totalLetters.count += sum
			}()
		}
	}
	wg.Wait()
	totalLetters.Lock()
	sum := totalLetters.count
	totalLetters.Unlock()
	fmt.Println("total letters=", totalLetters)
}
*/
