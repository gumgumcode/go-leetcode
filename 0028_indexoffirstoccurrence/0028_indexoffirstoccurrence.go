package main

import "fmt"

func findInString(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"

	index := findInString(haystack, needle)

	if index != -1 {
		fmt.Printf("The needle '%s' is found at index %d in the haystack '%s'\n", needle, index, haystack)
	} else {
		fmt.Printf("The needle '%s' is not found in the haystack '%s'\n", needle, haystack)
	}
}
