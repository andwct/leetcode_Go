package main

import (
	"fmt"
	"unicode"
)

// isPalindrome checks if the given string is a valid palindrome,
// considering only alphanumeric characters and ignoring cases.
func isPalindrome(s string) bool {
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	left, right := 0, len(filtered)-1
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
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"",
		"0P",
		"Madam",
	}

	for _, testCase := range testCases {
		result := isPalindrome(testCase)
		fmt.Printf("Input: %-30q => Is Palindrome? %v\n", testCase, result)
	}
}
