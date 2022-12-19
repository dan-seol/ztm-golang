package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

func init() {
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regular expression")
	}
	EmailExpr = compiled
	fmt.Println("regular expression compiled successfully")
}

func IsValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println(IsValidEmail("invalid"))
	fmt.Println(IsValidEmail("valid@example.com"))
	fmt.Println(IsValidEmail("invalid@example"))
}
