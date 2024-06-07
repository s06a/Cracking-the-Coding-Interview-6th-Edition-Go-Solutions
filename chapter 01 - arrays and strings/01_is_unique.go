package main

import (
	"fmt"
)

func isUnique(input string) bool {
	chars := make(map[rune]bool)
	for _, char := range input {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}

func main() {
	fmt.Println("input a word")
	var input string
	fmt.Scanf("%s", &input)

	// check if unqiue
	if isUnique(input) {
		fmt.Println(input, "is unique")
	} else {
		fmt.Println(input, "is not unique")
	}
}