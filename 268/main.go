package main

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

func missingNumber(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		if !set[v] {
			set[v] = true
		}
	}

	for i := 0; i <= len(set); i++ {
		if !set[i] {
			return i
		}
	}

	return -1
}
