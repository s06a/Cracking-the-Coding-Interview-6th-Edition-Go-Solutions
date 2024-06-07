package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// if count is even or there is just one character with
// odd count, then the string is palindrome
func isPermutationOfPalindrome(str string) bool {
	// remove spaces
	var tempString []string
	for _, char := range strings.Split(str, "") {
		if char != " " {
			tempString = append(tempString, char)
		}
	}

	str = strings.Join(tempString, "")

	// count occurance of characters
	counts := make(map[rune]int)
	for _, char := range str {
		counts[char]++
	}

	// check if only one character has odd count
	var hasOddCount bool
	for _, count := range counts {
		if count%2 != 0 {
			if hasOddCount {
				return false
			}
			hasOddCount = true
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input something to check")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	if isPermutationOfPalindrome(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}