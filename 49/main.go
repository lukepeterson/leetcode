package main

import (
	"sort"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)
	result := [][]string{}

	for _, anagram := range strs {
		anagramKey := sortString(anagram)
		hashmap[anagramKey] = append(hashmap[anagramKey], anagram)
	}

	for _, anagram := range hashmap {
		result = append(result, anagram)
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
