package main

func main() {}

// Input: nums = [10,5,2,6], k = 100
// Output: 8
//
// Explanation: The 8 subarrays that have product less than 100 are:
// [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]

func numSubarrayProductLessThanK(nums []int, k int) int {

	result := 0
	left := 0
	right := 0

	product := 1

	if k <= 1 {
		return 0
	}

	for right < len(nums) {
		product *= nums[right]
		for product >= k {
			product /= nums[left]
			left++
		}

		result += right - left + 1
		right++
	}

	return result
}
