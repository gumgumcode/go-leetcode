package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create a map to store the elements and their indices
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num

		// Check if the complement exists in the map
		if idx, found := numMap[complement]; found {
			// Return the indices of the two numbers
			return []int{idx, i}
		}

		// Add the current number to the map
		numMap[num] = i
	}

	// If no solution is found, return an empty slice
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0 1]
}
