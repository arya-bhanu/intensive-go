package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func ReverseString(str string) string {
	var runes []rune
	for i := len(str) - 1; i >= 0; i-- {
		runes = append(runes, rune(str[i]))
	}
	return string(runes)
}
