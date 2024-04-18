package main

import (
	"fmt"
)

func main() {

	fmt.Println(isPalindrome(""))

}

// Example 1: Given a string s, return true if it is a palindrome, false otherwise.

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
