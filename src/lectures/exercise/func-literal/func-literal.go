//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func classify(lines []string) {
	classifyFunc := func(r rune) (int, int, int, int) {
		letter, digit, space, punctuation := 0, 0, 0, 0
		if unicode.IsLetter(r) {
			letter = 1
		} else if unicode.IsDigit(r) {
			digit = 1
		} else if unicode.IsSpace(r) {
			space = 1
		} else if unicode.IsPunct(r) {
			punctuation = 1
		}

		return letter, digit, space, punctuation
	}
	countLetters, countDigits, countSpaces, countPunctuations := 0, 0, 0, 0
	for _, line := range lines {
		for _, char := range line {
			l, d, s, p := classifyFunc(char)
			countLetters += l
			countDigits += d
			countSpaces += s
			countPunctuations += p
		}
	}
	fmt.Println("Letters:", countLetters)
	fmt.Println("Digits:", countDigits)
	fmt.Println("Spaces:", countSpaces)
	fmt.Println("Punctuations:", countPunctuations)
}

func classifyImproved(lines []string) {
	classifyFunc := func(line string) (int, int, int, int) {
		letters, digits, spaces, punctuations := 0, 0, 0, 0
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++
			} else if unicode.IsDigit(r) {
				digits++
			} else if unicode.IsSpace(r) {
				spaces++
			} else if unicode.IsPunct(r) {
				punctuations++
			}
		}

		return letters, digits, spaces, punctuations
	}
	countLetters, countDigits, countSpaces, countPunctuations := 0, 0, 0, 0
	for _, line := range lines {
		l, d, s, p := classifyFunc(line)
		countLetters += l
		countDigits += d
		countSpaces += s
		countPunctuations += p
	}
	fmt.Println("Letters:", countLetters)
	fmt.Println("Digits:", countDigits)
	fmt.Println("Spaces:", countSpaces)
	fmt.Println("Punctuations:", countPunctuations)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	classifyImproved(lines)
}
