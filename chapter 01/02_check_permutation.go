package main

import (
	"fmt"
	"sort"
	"strings"
)

func isPermutation(firstString, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}

	firstString = sortString(firstString)
	secondString = sortString(secondString)

	return firstString == secondString
}

func sortString(str string) string {
	temp := strings.Split(str, "")
	sort.Strings(temp)
	return strings.Join(temp, "")
}

func main() {
	var firstString, secondString string
	fmt.Println("input first string")
	fmt.Scanf("%s", &firstString)
	fmt.Println("input second string")
	fmt.Scanf("%s", &secondString)

	if isPermutation(firstString, secondString) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}