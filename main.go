package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
}

func IsPalindrome(s string) bool {
	lowered := strings.ToLower(s)
	cleanedString := ""
	for _, r := range lowered {
		if validString(r) {
			cleanedString += string(r)
		}
	}
	maxloop := len(cleanedString) / 2
	for i := range maxloop {
		j := len(cleanedString) - 1 - i
		if cleanedString[i] != cleanedString[j] {
			return false
		}
	}
	return true
}

func validString(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
