package main

// Write a function that takes the binary representation of a positive integer and returns the number of set bits it has (also known as the Hamming weight).
//
// Example 1:
// Input: n = 11
// Output: 3
// Explanation: The input binary string 1011 has a total of three set bits.
//
// Example 2:
// Input: n = 128
// Output: 1
// Explanation: The input binary string 10000000 has a total of one set bit.

func hammingWeight(n int) int {

	result := 0
	for n > 0 {
		if n&1 == 1 {
			result++
		}
		n >>= 1
	}
	return result
}
