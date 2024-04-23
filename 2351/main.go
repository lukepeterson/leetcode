package main

// Given a string s consisting of lowercase English letters, return the first letter to appear twice.

// Note:
// - A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
// - s will contain at least one letter that appears twice.

// With hashmap (keys and values):
// func repeatedCharacter(s string) byte {
// 	hashmap := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		_, ok := hashmap[s[i]]
// 		if ok {
// 			return byte(s[i])
// 		} else {
// 			hashmap[s[i]]++
// 		}
// 	}

// 	return 0
// }

// With hashmap as a set
func repeatedCharacter(s string) byte {
	set := map[byte]bool{}
	for _, v := range s {
		if set[byte(v)] {
			return byte(v) // we've been here before
		} else {
			set[byte(v)] = true // otherwise, set for next time
		}
	}

	return 0
}
