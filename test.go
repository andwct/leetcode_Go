package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var filtered []rune
	//  check if
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered = append(filtered, unicode.ToLower(char))
		}
	}
	left := 0
	right := len(filtered) - 1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	inputStr := "A man, a plan, a canal: Panama"
	result := isPalindrome(inputStr)
	fmt.Println(result)

}
