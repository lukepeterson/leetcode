package main

// Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

// If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.

// 01234567

//  abcdefg
//l ^
//r ^

func reverseStr(s string, k int) string {
	result := []rune(s)
	right := 0

	for i := 0; i < len(s); i += k * 2 {
		right = i + k
		if right > len(s) {
			right = len(s)
		}

		for j := 0; j < (right-i)/2; j++ {
			result[i+j], result[right-j-1] = result[right-j-1], result[i+j]
		}
	}

	return string(result)
}
