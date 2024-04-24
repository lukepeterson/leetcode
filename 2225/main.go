package main

import "sort"

// You are given an integer array matches where matches[i] = [winneri, loseri] indicates that the player winneri defeated player loseri in a match.

// Return a list answer of size 2 where:
// - answer[0] is a list of all players that have not lost any matches.
// - answer[1] is a list of all players that have lost exactly one match.
// - The values in the two lists should be returned in increasing order.

// Note:
// - You should only consider the players that have played at least one match.
// - The testcases will be generated such that no two matches will have the same outcome.

const winner = 0
const loser = 1

func findWinners(matches [][]int) [][]int {
	losses := make(map[int]int)

	for match := range matches {
		if losses[matches[match][winner]] == 0 {
			losses[matches[match][winner]] = 0
		}
		losses[matches[match][loser]]++
	}

	notLostAny := []int{}
	lostOne := []int{}
	for player := range losses {
		if losses[player] == 0 {
			notLostAny = append(notLostAny, player)
			continue
		}
		if losses[player] == 1 {
			lostOne = append(lostOne, player)
		}
	}
	sort.Ints(notLostAny)
	sort.Ints(lostOne)

	return [][]int{notLostAny, lostOne}
}
