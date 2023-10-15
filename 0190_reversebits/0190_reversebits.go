package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}
	return result
}

// num & 1: This part of the expression performs a bitwise AND operation between
// num and 1. Since 1 in binary is represented as 00000001, this operation
// effectively extracts the least significant bit of num. For example, if num is
// 11010100, then num & 1 would result in 00000000 because all the bits in num
// are ANDed with 0 except for the last bit.

// result |= num & 1: Here, result is a variable used to build the reversed
// number. The |= operator is a bitwise OR assignment operator, which means it
// performs a bitwise OR between result and the value on the right side (num &
// 1) and assigns the result back to result. This operation effectively appends
// the least significant bit extracted in the previous step to the result.

func main() {
	num := uint32(43261596) // You can replace this with the input number you want to reverse.
	reversed := reverseBits(num)
	fmt.Printf("Reversed bits of %d: %d\n", num, reversed)
	fmt.Printf("Original number %d in binary: %b\n", num, num)
	fmt.Printf("Reversed bits of %d in binary: %b\n", num, reversed)
}
