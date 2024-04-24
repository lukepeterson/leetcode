package main

import "fmt"

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.
// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

func equalPairs(grid [][]int) int {
	rows := make(map[string]int)
	cols := make(map[string]int)
	totalPairs := 0

	for i := 0; i < len(grid); i++ {
		row := []int{}
		col := []int{}
		for j := 0; j < len(grid[0]); j++ {
			row = append(row, grid[i][j])
			col = append(col, grid[j][i])
		}

		rows[convertToString(row)]++
		cols[convertToString(col)]++
	}

	for i := range rows {
		if cols[i] != 0 {
			totalPairs += rows[i] * cols[i]
		}
	}

	return totalPairs
}

func convertToString(input []int) string {
	result := ""
	for _, digit := range input {
		result += fmt.Sprint(digit) + "|"
	}
	return result
}
