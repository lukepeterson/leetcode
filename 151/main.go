package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseWords("the sky is blue  "))
}

// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.
//
// Example 1:
// Input: s = "the sky is blue"
// Output: "blue is sky the"

// Example 2:
// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

// Example 3:
// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

// 2ND ATTEMPT (using built-ins)
func reverseWords(s string) string {
	split := strings.Fields(s)
	var result = []string{}

	for i := range split {
		result = append(result, split[len(split)-i-1])
	}

	return strings.Join(result, " ")
}

// 1ST ATTEMPT - PANIC ON INDEX[-1]
// func reverseWords(s string) string {

// 	wordStart := 0
// 	wordEnd := len(s) - 1

// 	var split string

// 	for wordEnd > 0 {
// 		for s[wordEnd] == ' ' {
// 			wordEnd--
// 		}

// 		wordStart = wordEnd
// 		for s[wordStart] != ' ' && wordStart > 0 {
// 			wordStart--
// 		}
// 		wordStart++

// 		split += s[wordStart : wordEnd+1]

// 		for s[wordEnd] != ' ' && wordEnd > 0 {
// 			wordEnd--
// 		}
// 		if wordEnd == 0 {
// 			break
// 		}
// 	}

// 	returnsplit
// }
