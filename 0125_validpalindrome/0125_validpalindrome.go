package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove non-alphanumeric characters
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	// Check if the resulting string is a palindrome
	return isPalindromeHelper(s)
}

func removeNonAlphanumeric(s string) string {
	var result strings.Builder

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func isPalindromeHelper(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	input := "A man, a plan, a canal: Panama"
	fmt.Printf("Is the string '%s' a valid palindrome? %v\n", input, isPalindrome(input))
}
