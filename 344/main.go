package main

func main() {
	// test := []byte{'a', 'c', 'e'}
	// test[0] = test[2]

	// // test[1], test[2] = test[2], test[1]

	// fmt.Println(test)

}

// Write a function that reverses a string. The input string is given as an array of characters s.
//
// You must do this by modifying the input array in-place with O(1) extra memory.
//
// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
//
// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

//   hello
// i ^     0
// j     ^ 4

//   oellh
// i  ^    1
// j    ^  3

//   olleh
// i   ^  2
// j   ^  2

func reverseString(s []byte) []byte {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
