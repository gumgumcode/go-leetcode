package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIdx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIdx] {
			uniqueIdx++
			nums[uniqueIdx] = nums[i]
		}
	}

	return uniqueIdx + 1
}

func main() {
	// Sample array with duplicates
	nums := []int{1, 1, 2, 2, 2, 3, 4, 4, 5, 5, 5}

	fmt.Println("Original Array:", nums)

	// Call the function to remove duplicates in-place
	length := removeDuplicates(nums)

	fmt.Println("Modified Array:", nums[:length])
	fmt.Println("New Length:", length)
}
