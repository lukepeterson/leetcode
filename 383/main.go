package main

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make(map[rune]int)

	if len(magazine) < len(ransomNote) {
		return false
	}

	for _, letter := range magazine {
		hashmap[letter]++
	}

	for _, letter := range ransomNote {
		if hashmap[letter] > 0 {
			hashmap[letter]--
		} else {
			return false
		}
	}

	return true
}
