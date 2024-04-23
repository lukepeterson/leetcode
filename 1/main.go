package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// One pass with hash tables
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := hashmap[complement]; ok {
			return []int{j, i}
		}
		hashmap[num] = i
	}

	return nil
}

// // Brute force solution - O(n^2)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if i != j {
// 				if nums[i]+nums[j] == target {
// 					return []int{i, j}
// 				}
// 			}
// 		}
// 	}
//
// 	return []int{}

// // Two pass with hash tables
// func twoSum(nums []int, target int) []int {
// 	// Convert our array into a hashmap
// 	hashmap := make(map[int]int)
// 	for i, v := range nums {
// 		hashmap[v] = i
// 	}

// 	for i, num := range nums {
// 		// Check whether our hashmap has an element that is target - nums[i]
// 		v, ok := hashmap[target-num]
// 		if ok && i != v { // Make sure we don't count the same element twice
// 			return []int{i, v}
// 		}
// 	}
// 	return nil
// }
