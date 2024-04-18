package main

import (
	"fmt"
	"strconv"
)

// Given an array of characters chars, compress it using the following algorithm:
//
// Begin with an empty string s. For each group of consecutive repeating characters in chars:
//
// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
// Note that group lengths that are 10 or longer will be split into multiple characters in chars.
//
// After you are done modifying the input array, return the new length of the array.
//
// You must write an algorithm that uses only constant extra space.
//
// Example 1:
// Input: chars = ["a","a","b","b","c","c","c"]
// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

// Example 2:
// Input: chars = ["a"]
// Output: Return 1, and the first character of the input array should be: ["a"]
// Explanation: The only group is "a", which remains uncompressed since it's a single character.

// Example 3:
// Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
// Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
// Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".

func main() {

	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b'}))

}

func compress(chars []byte) int {
	i := 0
	res := 0
	for i < len(chars) {
		grouplength := 1
		for i+grouplength < len(chars) && chars[i+grouplength] == chars[i] {
			grouplength += 1
		}
		chars[res] = chars[i]
		res += 1
		if grouplength > 1 {
			for _, c := range []byte(strconv.Itoa(grouplength)) {
				chars[res] = c
				res += 1
			}
		}
		i += grouplength
	}
	return res
}

// func compress(chars []byte) int {

// 	result := ""
// 	char := byte(chars[0])
// 	count := 0

// 	for _, c := range chars {
// 		// for i := 0; i < len(chars); i++ {
// 		if c != char {
// 			if count == 1 {
// 				result += string(char)
// 			} else {
// 				result += string(char) + fmt.Sprint(count)
// 			}
// 			char = c
// 			count = 0
// 		}
// 		count++
// 	}

// 	if count == 1 {
// 		result += string(char)
// 	} else {
// 		result += string(char) + fmt.Sprint(count)
// 	}

// 	return len(result)
// }
