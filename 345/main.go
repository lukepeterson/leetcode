package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
}

// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
//
// Example 1:
// Input: s = "hello" .  eo
// Output: "holle" .     oe

// Example 2:
// Input: s = "leetcode" . eeoe
// Output: "leotcede" .    eoee

func reverseVowels(s string) string {

	var vowelIndexes []int
	for i, letter := range s {
		if isVowel(rune(letter)) {
			vowelIndexes = append(vowelIndexes, i)
		}
	}

	result := []rune(s)
	for i, j := range vowelIndexes {
		result[j] = rune(s[vowelIndexes[len(vowelIndexes)-i-1]])
	}
	return string(result)

}

func isVowel(letter rune) bool {

	return (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' ||
		letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U')

}
