package main

import "math"

func main() {

}

// We need to find the length of the longest subarray that has a
// sum less than or equal to k.

// For this example, let nums = [3, 2, 1, 3, 1, 1] and k = 5.

//  3, 2, 1, 3, 1, 1
//l ^
//r ^

func longestSubArray(nums []int, k int) int {

	// sliding window pointers
	left := 0
	right := 0

	maximumSum := 0
	currentSum := 0

	for right < len(nums) {
		currentSum += nums[right]

		// constraint failed.  sum is now more than k, so we move
		// the left side of the window until our constraint is remet
		for currentSum > k {
			currentSum -= nums[left] // subtract the left from our sum
			left++
		}

		// after all that window moving, check whether we have a new maximum
		maximumSum = int(math.Max(float64(maximumSum), float64(right-left+1)))

		right++ // expand the window
	}

	return maximumSum
}
