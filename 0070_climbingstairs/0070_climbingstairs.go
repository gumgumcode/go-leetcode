package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev1, prev2 = prev2, current
	}

	return prev2
}

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs2(n-1) + climbStairs2(n-2)
}

func main() {
	n := 5 // Change this to the desired value of n
	ways := climbStairs2(n)
	fmt.Printf("Number of distinct ways to climb %d stairs: %d\n", n, ways)
}

// Explanation of the recursive method:

//   1 2 3 5 8   WAYS
// _ _ _ _ _ _   STAIRS
// 0 1 2 3 4 5   STAIR NUMBER

// for f(4), we get ans from f(4-1) + f(4-2)
// i.e. f(3) + f(2) = 3+2 = 5

// for f(5), we get ans from f(5-1) + f(5-2)
// i.e. f(4) + f(3) = 5+3 = 8

// {
// 	1: 1
// 	2: {
// 		1+1,
// 		2
// 	},
// 	3: {
// 		1+1+1,
// 		1+2
// 		2+1
// 	},
// 	4: {
// 		1+1+1+1,
// 		2+2,
// 		2+1+1,
// 		1+2+1,
// 		1+1+2
// 	}
// }
