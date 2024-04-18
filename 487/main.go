package main

import "math"

func main() {}

func findMaxConsecutiveOnes(nums []int) int {

	left := 0
	right := 0

	maxLength := 0
	zeros := 0

	for right < len(nums) {

		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))

		right++
	}

	return maxLength

}
