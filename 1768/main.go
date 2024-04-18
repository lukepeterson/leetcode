package main

import (
	"fmt"
)

func main() {

	fmt.Println(mergeAlternately("abcd", "pq"))

}

// Example 1:

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

// Example 2:

// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s

// Example 3:

// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q
// merged: a p b q c   d

func mergeAlternately(word1 string, word2 string) string {

	// 2ND ATTEMPT - ONE POINTER:
	var result string
	maxLength := len(word1) + len(word2)
	for i := 0; i < maxLength; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}
	}
	return string(result)

	// 1ST ATTEMPT - TWO POINTERS:
	// var result string
	// var word1Pointer int
	// var word2Pointer int

	// for i := 0; i < len(word1+word2); i++ {
	// 	if word1Pointer < len(word1) {
	// 		result += string(word1[word1Pointer])
	// 		word1Pointer++
	// 	}

	// 	if word2Pointer < len(word2) {
	// 		result += string(word2[word2Pointer])
	// 		word2Pointer++
	// 	}
	// }

	// return string(result)
}
