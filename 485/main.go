package main

import (
	"math"
)

func main() {

}

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
//
// Example 1:
// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s.
// The maximum number of consecutive 1s is 3.

// Example 2:
// Input: nums = [1,0,1,1,0,1]
// Output: 2

//i 0 1 2 3
// [1,1,0,1]
//l ^
//r     ^

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
		}

		maxCount = int(math.Max(float64(maxCount), float64(count)))
	}

	return maxCount
}
