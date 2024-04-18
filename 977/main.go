package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))

}

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
//
// Example 1:
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

func sortedSquares(nums []int) []int {

	i := 0
	j := len(nums)

	for i < j {
		nums[i] *= nums[i]

		// if i > 0 {
		// 	nums[i], nums[i-1] = nums[i-1], nums[i]
		// 	// fmt.Println("swap ", nums[i], " and ", nums[i-1])
		// }

		// if nums[i] > nums[i+1] {
		// 	nums[i], nums[i+1] = nums[i+1], nums[i]
		// }

		i++
	}

	// [16 1 9 100 0]
	// [1 16 9 100 0]

	// i = 0
	// for i < j-1 {

	// 	if nums[i] > nums[i+1] {
	// 		nums[i], nums[i+1] = nums[i+1], nums[i]
	// 	}
	// 	i++
	// }

	sort.Ints(nums)

	return nums
}
