package main

import "math"

func main() {}

// You are given a binary string s (a string containing only "0" and "1").
// You may choose up to one "0" and flip it to a "1".
// What is the length of the longest substring achievable that contains only "1"s?

// For example, given s = "1101100111", the answer is 5.
// If you perform the flip at index 2, the string becomes 1111100111.

//  1101100111
//l ^
//r  ^

func longestSubstring(input string) int {

	left := 0
	right := 0

	inputLength := len(input)
	// currentLength := 0
	maximumLength := 0
	zeros := 0

	for right < inputLength {

		if input[right] == '0' {
			zeros++
		}

		// constraint metric = "how many 0s are in the substring"
		// numeric restriction = "<= one 0s"
		for zeros > 1 {
			if input[left] == '0' {
				zeros--
			}
			left++
		}

		maximumLength = int(math.Max(float64(maximumLength), float64(right-left+1)))

		right++
	}

	return maximumLength

}
