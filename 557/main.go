package main

func reverseWords(s string) string {
	result := []rune(s)
	result = append(result, ' ')

	left := 0
	right := 0
	lastSpace := -1

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			left = lastSpace + 1
			right = i - 1
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			lastSpace = i
		}

	}

	return string(result[:len(result)-1])
}
