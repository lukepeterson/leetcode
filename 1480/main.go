package main

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

// Return the running sum of nums.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

// Example 3:

// Input: nums = [3,1,2,10,1]
// Output: [3,4,6,16,17]

func runningSum(nums []int) []int {
	prefixSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefixSum = append(prefixSum, nums[i]+prefixSum[i-1])
	}
	return prefixSum
}

// solve time = 6 minutes
