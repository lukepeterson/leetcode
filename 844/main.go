package main

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '#' {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			} else {
				stack = []rune{}
			}
		} else {
			stack = append(stack, rune(char))
		}
	}

	return string(stack)
}
