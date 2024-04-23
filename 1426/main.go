package main

// Given an integer array arr, count how many elements x there are, such that x + 1 is also in arr. If there are duplicates in arr, count them separately.

// Input: arr = [1,2,3]
// Output: 2
// Explanation: 1 and 2 are counted because 2 and 3 are in arr.

// Input: arr = [1,1,3,3,5,5,7,7]
// Output: 0
// Explanation: No numbers are counted, because there is no 2, 4, 6, or 8 in arr.

func countElements(arr []int) int {
	count := 0
	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}

	for _, v := range arr {
		if set[v+1] {
			count++
		}
	}

	return count
}
