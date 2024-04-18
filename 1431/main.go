package main

import (
	"fmt"
	"slices"
)

func main() {

	var candies = []int{2, 3, 5, 1, 3}
	var extraCandies int = 3
	fmt.Println(kidsWithCandies(candies, extraCandies))

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxCandies := slices.Max(candies)

	result := []bool{}
	for i := range candies {
		result = append(result, candies[i]+extraCandies >= maxCandies)
	}
	return result

	// 1ST ATTTEMPT
	//
	// for i := range candies {
	// 	extraCandyTest := candies[i] + extraCandies

	// 	mostCandies := false
	// 	for j := range candies {
	// 		mostCandies = false
	// 		if extraCandyTest >= candies[j] {
	// 			mostCandies = true
	// 		} else {
	// 			break
	// 		}
	// 	}

	// 	result = append(result  mostCandies)
	// }

	// return result
}
