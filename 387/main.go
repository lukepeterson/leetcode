package main

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

func firstUniqChar(s string) int {
	// Count the number of occurances of each character
	hashmap := make(map[rune]int)
	for _, v := range s {
		hashmap[v]++
	}

	// Go and find the first character that we only saw once
	for i, v := range s {
		if hashmap[v] == 1 {
			return i
		}
	}

	return -1
}
