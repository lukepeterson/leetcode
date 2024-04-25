package main

// Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	letterMap := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if prev, ok := letterMap[s[right]]; ok {
			left = max(left, prev)
			letterMap[s[right]]--
		}

		maxLength = max(maxLength, right-left+1)
		letterMap[s[right]] = right + 1
	}
	return maxLength
}
