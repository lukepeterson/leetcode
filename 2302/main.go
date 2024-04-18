package main

func main() {

}

// The score of an array is defined as the product of its sum and its length.
// For example, the score of [1, 2, 3, 4, 5] is (1 + 2 + 3 + 4 + 5) * 5 = 75.
// Given a positive integer array nums and an integer k,
// return the number of non-empty subarrays of nums whose score is
// strictly less than k.
// A subarray is a contiguous sequence of elements within an array.

// Input: nums = [2,1,4,3,5], k = 10
// Output: 6
// Explanation:
// The 6 subarrays having scores less than 10 are:
// - [2] with score 2 * 1 = 2.
// - [1] with score 1 * 1 = 1.
// - [4] with score 4 * 1 = 4.
// - [3] with score 3 * 1 = 3.
// - [5] with score 5 * 1 = 5.
// - [2,1] with score (2 + 1) * 2 = 6.
// Note that subarrays such as [1,4] and [4,3,5] are not considered because their scores are 10 and 36 respectively, while we need scores strictly less than 10.

// {2, 1, 4, 3, 5}, k: 10
//
//	0  1  2  3  4
func countSubarrays(nums []int, k int64) int64 {
	left := 0
	right := 0
	runningSum := 0
	score := 0
	result := 0

	for right < len(nums) {

		runningSum += nums[right]

		score = runningSum * (right - left + 1)

		for score >= int(k) {
			runningSum -= nums[left]
			left++
			score = runningSum * (right - left + 1)
		}

		result += (right - left + 1)
		right++
	}

	return int64(result)
}

// 0  1  2  3  4
// 2, 1, 4, 3, 5

// 2

//    1
// 2  1

//       4
//    1  4
// 2  1  4

// 8 9 3 7 8 7 7
//         ^
//   ^
//
// l = 1
// r = 5

//     7
//    87
//   787
//  3787
// 92787

//    8
//   78
//  378
// 9378
