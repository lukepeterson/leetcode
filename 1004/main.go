package main

import (
	"fmt"
	"math"
)

func main() {}

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0
	maxLength := 0
	zeros := 0

	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}

		fmt.Println(left, right, zeros, maxLength)

		// whilst the window is invalid
		for zeros > k {
			// move the left pointer
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

// k=2 (max zeros to flip)

//        1,1,1,0,0,0,1,1,1,1,0
//
// left:          ^
// right:           ^
// max:
