package main

func main() {}

func prefixSum(input []int) []int {

	// Prefix sum is a technique that can be used on arrays (of numbers).
	// The idea is to create an array prefix where prefix[i] is the sum of
	// all elements up to the index i (inclusive).

	// For example, given nums = [5, 2, 1, 6, 3, 8],
	// we would have prefix = [5, 7, 8, 14, 17, 25].

	prefix := []int{input[0]}
	for i := 1; i < len(input); i++ {
		// for each prefix
		// append the current value + the one before it
		prefix = append(prefix, input[i]+prefix[i-1])
	}

	return prefix

}
