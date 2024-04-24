package main

// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

func findMaxLength(nums []int) int {
	h := make(map[int]int)
	h[0] = -1

	result := 0
	sum := 0

	for i, num := range nums {
		if num == 0 {
			sum -= 1
		} else {
			sum += 1
		}

		prevIndex, exists := h[sum]
		if exists {
			length := i - prevIndex
			result = max(length, result)
		} else {
			h[sum] = i
		}
	}

	return result
}
