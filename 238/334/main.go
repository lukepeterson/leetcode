package main

import (
	"fmt"
	"math"
)

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
// If no such indices exists, return false.
//
// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
//
// Example 2:
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
//
// Example 3:
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

func main() {
	// fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	// fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	// fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))

}

func increasingTriplet(nums []int) bool {

	i := math.MaxInt32
	j := math.MaxInt32

	if len(nums) < 3 {
		return false
	}

	for _, n := range nums {
		if n <= i {
			i = n
		} else if n <= j {
			j = n
		} else {
			return true
		}
	}

	// 1st attempt - doesn't take into account nums[j] could be more than 1 away from nums[k]
	// for i := 0; i <= len(nums)-3; i++ {
	// 	if nums[i] < nums[i+1] && nums[i+1] < nums[i+2] {
	// 		return true
	// 	}
	// }

	return false

}
