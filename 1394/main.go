package main

// Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

// Return the largest lucky integer in the array. If there is no lucky integer return -1.

func findLucky(arr []int) int {
	hashmap := make(map[int]int)
	result := -1

	for _, v := range arr {
		hashmap[v]++
	}

	for i := range hashmap {
		if hashmap[i] == i {
			result = max(result, i)
		}
	}

	return result
}
