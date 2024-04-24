package main

// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

// Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3

func numJewelsInStones(jewels string, stones string) int {

	jewelCount := 0
	jewelsMap := make(map[byte]bool)
	for i := range jewels {
		jewelsMap[jewels[i]] = true
	}

	for i := range stones {
		if jewelsMap[stones[i]] {
			jewelCount++
		}
	}

	return jewelCount
}
