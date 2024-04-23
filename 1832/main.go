package main

// A pangram is a sentence where every letter of the English alphabet appears at least once.

// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.

func checkIfPangram(sentence string) bool {
	seen := make(map[rune]bool)
	count := 0

	for _, v := range sentence {
		if !seen[v] {
			seen[v] = true
			count++
		}

		if count >= 26 { // edge case: exit early if we've found a pangram already
			return true
		}
	}

	return count >= 26
}
