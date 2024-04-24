package main

// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

// You can use each character in text at most once. Return the maximum number of instances that can be formed.

func maxNumberOfBalloons(text string) int {
	h := make(map[byte]int)
	for i := range text {
		if isBalloonChar(text[i]) {
			h[text[i]]++
		}
	}

	h['l'] = h['l'] / 2
	h['o'] = h['o'] / 2

	return min(h['b'], min(h['a'], min(h['l'], min(h['o'], min(h['n'])))))
}

func isBalloonChar(ch byte) bool {
	if ch == 'b' || ch == 'a' || ch == 'l' || ch == 'o' || ch == 'n' {
		return true
	}

	return false
}
