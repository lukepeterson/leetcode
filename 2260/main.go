package main

import (
	"math"
)

// You are given an integer array cards where cards[i] represents the value of the ith card. A pair of cards are matching if the cards have the same value.

// Return the minimum number of consecutive cards you have to pick up to have a pair of matching cards among the picked cards. If it is impossible to have matching cards, return -1.

// Input: cards = [3,4,2,3,4,7]
// Output: 4

func minimumCardPickup(cards []int) int {
	hashmap := make(map[int][]int)
	minimumCards := math.MaxInt

	for i, card := range cards {
		hashmap[card] = append(hashmap[card], i)
	}

	for _, v := range hashmap {
		for i := 0; i < len(v)-1; i++ {
			minimumCards = min(minimumCards, v[i+1]-v[i]+1)
		}
	}

	if minimumCards == math.MaxInt {
		return -1
	}

	return minimumCards
}
