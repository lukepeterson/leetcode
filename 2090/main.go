package main

// You are given a 0-indexed array nums of n integers, and an integer k.
//
// The k-radius average for a subarray of nums centered at some index i with the radius k is the average of all elements in nums between the indices i - k and i + k (inclusive). If there are less than k elements before or after the index i, then the k-radius average is -1.
//
// Build and return an array avgs of length n where avgs[i] is the k-radius average for the subarray centered at index i.
//
// The average of x elements is the sum of the x elements divided by x, using integer division. The integer division truncates toward zero, which means losing its fractional part.
//
// For example, the average of four elements 2, 3, 1, and 5 is (2 + 3 + 1 + 5) / 4 = 11 / 4 = 2.75, which truncates to 2.
//

func getAverages(nums []int, k int) []int {
	// start by filling all elements of the array with -1
	result := []int{}
	for i := 0; i < len(nums); i++ {
		result = append(result, -1)
	}

	// calculate the prefixSum array, which we'll use to calculate the averages for the sliding window
	prefixSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
	}

	width := k*2 + 1                   // width is the "radius", which includes k on each side, plus the middle of the window
	for i := k; i < len(nums)-k; i++ { // for each window of size `k*2+1` that fits within our nums array
		startIndex := i - k // start in the middle of the leftmost window
		endIndex := i + k   // end in the middle of the leftmost window
		// calculate the average of the window size
		result[i] = (prefixSum[endIndex] - prefixSum[startIndex] + nums[startIndex]) / width
	}

	return result
}
