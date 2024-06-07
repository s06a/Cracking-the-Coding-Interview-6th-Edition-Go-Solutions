package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceSpaces(str string) string {
	tempString := strings.Split(str, "")
	finalString := make([]string, 0)
	for _, char := range tempString {
		if char == " " {
			finalString = append(finalString, "%20")
		} else {
			finalString = append(finalString, char)
		}
	}
	return strings.Join(finalString, "")
}

func main() {
	// scanf does not work for multiple words
	input := bufio.NewReader(os.Stdin)
	fmt.Println("input a sentence")
	str, _ := input.ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println(replaceSpaces(str))
}