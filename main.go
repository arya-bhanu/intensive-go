package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello World")
}

func CountWordFrequency(text string) map[string]int {
	IsDelimiter := func(c rune) bool {
		return c == ' ' || c == '\n' || c == '\t' || c == '-'
	}

	maps := make(map[string]int)
	splitted := strings.FieldsFunc(text, IsDelimiter)

	if len(splitted) == 1 && splitted[0] == "" {
		return maps
	}

	for _, word := range splitted {
		lowered := FilterLetter(strings.ToLower(word))

		currCount, ok := maps[lowered]
		if !ok {
			maps[lowered] = 1
			continue
		}
		currCount++
		maps[lowered] = currCount
	}
	return maps
}

func FilterLetter(word string) string {
	var cleanWord []rune
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleanWord = append(cleanWord, r)
		}
	}
	return string(cleanWord)
}
