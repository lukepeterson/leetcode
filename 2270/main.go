package main

// You are given a 0-indexed integer array nums of length n.

// nums contains a valid split at index i if the following are true:

// The sum of the first i + 1 elements is greater than or equal to the sum of the last n - i - 1 elements.
// There is at least one element to the right of i. That is, 0 <= i < n - 1.

// Return the number of valid splits in nums.

// func waysToSplitArray(nums []int) int {
// 	splitCount := 0
// 	prefixSum := []int{nums[0]}
// 	for i := 1; i < len(nums); i++ {
// 		prefixSum = append(prefixSum, nums[i]+prefixSum[i-1])
// 	}

// 	for i := 0; i < len(prefixSum)-1; i++ {
// 		leftPrefixSum := prefixSum[i]
// 		rightPrefixSum := prefixSum[len(nums)-1] - prefixSum[i]
// 		if leftPrefixSum >= rightPrefixSum {
// 			splitCount++
// 		}
// 	}

// 	return splitCount
// }

func waysToSplitArray(nums []int) int {

	splitCount := 0
	total := 0
	for i := range nums {
		total += nums[i]
	}

	left := 0
	for i := range len(nums) - 1 {
		left += nums[i]
		right := total - left
		if left >= right {
			splitCount++
		}
	}

	return splitCount
}
