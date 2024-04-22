package main

// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

func maxVowels(s string, k int) int {

	vowel := map[byte]bool{}
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		vowel[v] = true
	}

	maxVowels := 0
	currVowels := 0

	// Count the number of vowels in our first window
	// It's important we establish this baseline, which is O(n) time and O(1) space
	for i := 0; i < k; i++ {
		if vowel[s[i]] {
			currVowels++
		}
	}
	maxVowels = currVowels

	// Edge case = no need to continue if we've already maxed out our vowel count
	if maxVowels == k {
		return k
	}

	// Start checking the other windows for maximum length substrings
	// We start with i:=1 to slide the left of the window one character to the right and
	// continue until the right of the window equals k
	for left := 1; left < len(s)-k+1; left++ {
		if vowel[s[left-1]] { // if the item we just slid from is a vowel
			currVowels--
		}
		if vowel[s[left+k-1]] { // if the item we just slid into is a vowel
			currVowels++
		}

		maxVowels = max(currVowels, maxVowels) // Have we reached a new record?
	}

	return maxVowels
}

// func vowel(char byte) bool {
// 	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
// }
