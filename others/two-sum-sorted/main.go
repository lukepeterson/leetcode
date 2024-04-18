package main

func main() {}

// Example 2: Given a sorted array of unique integers and a target integer, return true if there exists a pair of numbers that sum to target, false otherwise.
// For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true because 4 + 9 = 13.

// func twoSum(input []int, target int) bool {
func twoSum(input []int, target int) bool {
	left := 0
	right := len(input) - 1

	for left < right {

		curr := input[left] + input[right]
		if curr == target {
			return true
		}
		if curr > target {
			right--
		} else {
			left++
		}
	}

	return false
}
