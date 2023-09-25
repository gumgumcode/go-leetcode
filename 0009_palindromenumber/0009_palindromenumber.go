package main

import "fmt"

func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Reverse the integer and compare it with the original
	reversed, original := 0, x
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return reversed == original
}

func main() {
	// Example usage:
	num := 121
	result := isPalindrome(num)
	fmt.Println(result) // Output: true

	num = 122
	result = isPalindrome(num)
	fmt.Println(result) // Output: false
}
