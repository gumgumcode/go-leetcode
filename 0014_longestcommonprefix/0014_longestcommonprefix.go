package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Find the shortest string in the array
	shortest := strs[0]
	for _, s := range strs {
		if len(s) < len(shortest) {
			shortest = s
		}
	}

	for i := 0; i < len(shortest); i++ {
		for j := 0; j < len(strs); j++ {
			if shortest[i] != strs[j][i] {
				return shortest[:i]
			}
		}
	}

	return shortest
}

func main() {
	input := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(input)

	fmt.Printf("The longest common prefix is: %s\n", result)
}

// the overall time complexity of this approach is O(N + S), where N is the
// number of strings, and S is the total number of characters in the input
// strings
