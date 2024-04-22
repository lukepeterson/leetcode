package main

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// func moveZeroes(nums []int) []int {
// 	right := 0
// 	zeroCount := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			zeroCount++
// 		}
// 	}
// 	for left := 0; left < len(nums)-zeroCount; left++ {
// 		if nums[left] == 0 {
// 			for nums[right] == 0 {
// 				right++
// 			}
// 			nums[left], nums[right] = nums[right], nums[left]
// 		}
// 		right = left + 1
// 	}
// 	return nums
// }

func moveZeroes(nums []int) []int {
	lastNonZeroFoundAt := 0
	for cur := range nums {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
	return nums
}
