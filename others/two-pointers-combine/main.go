package main

func main() {

}

// Example 3: Given two sorted integer arrays arr1 and arr2, return a new array that combines both of them and is also sorted.

// args: args{arr1: []int{1, 5, 7}, arr2: []int{2, 4, 9}},
// want: []int{1, 2, 4, 5, 7, 9},

// 1 5 7
// ^
// 2 4 9
// ^

// 2 3
// ^
// 1 4
//   ^

func combine(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}
