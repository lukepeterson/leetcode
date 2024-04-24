package main

// Given an integer array nums, return the largest integer that only occurs once. If no integer occurs once, return -1.

func largestUniqueNumber(nums []int) int {

	hashmap := make(map[int]int)
	for i := range nums { // O(n)
		hashmap[nums[i]]++
	}

	largest := -1
	for i, v := range hashmap { // O(n)
		if v == 1 && i > largest {
			largest = i
		}
	}

	return largest
}
