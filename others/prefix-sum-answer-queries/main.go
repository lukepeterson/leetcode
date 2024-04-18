package main

func main() {

	// fmt.Println(answerQueries([]int{1}, [][]int{{3, 3}}, 3))

}

// Example 1: Given an integer array nums, an array queries where queries[i] = [x, y] and an integer limit,
// return a boolean array that represents the answer to each query.
// A query is true if the sum of the subarray from x to y is less than limit, or false otherwise.

// For example, given nums = [1, 6, 3, 2, 7, 2], queries = [[0, 3], [2, 5], [2, 4]], and limit = 13,
// the answer is [true, false, true]. For each query, the subarray sums are [12, 14, 12].

func answerQueries(input []int, queries [][]int, limit int) []bool {

	answer := []bool{}

	prefix := []int{input[0]}
	for i := 1; i < len(input); i++ {
		prefix = append(prefix, input[i]+prefix[i-1])
	}

	// fmt.Println(prefix)

	for _, query := range queries {

		x := query[0]
		y := query[1]
		//index   0  1   2   3   4   5
		//input  [1, 6,  3,  2,  7,  2]
		//prefix [1, 7, 10, 12, 19, 21]

		//		[2 5] 10 21

		// fmt.Println(query, prefix[y]-prefix[x]+input[x])
		answer = append(answer, prefix[y]-prefix[x]+input[x] < limit)
	}

	// queries[0] := [][]int{

	// pairs := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// fmt.Println(pairs[0])
	// fmt.Println(pairs[1][1])

	return answer
}
