package main

import "math"

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2

// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

func minSubArrayLen(target int, nums []int) int {

	left := 0
	minimumLength := math.MaxInt
	subArraySum := 0

	for right := 0; right < len(nums); right++ {

		subArraySum += nums[right]

		// if the sum of the subarray is greater than the target
		for subArraySum >= target {

			// if the minimum length is less than the old minimum length
			if right-left+1 < minimumLength {

				// set new minimum length
				minimumLength = right - left + 1
			}

			subArraySum -= nums[left]
			left++
		}
	}

	if minimumLength == math.MaxInt {
		return 0
	}

	return minimumLength
}
