package main

import (
	"fmt"
)

func day18() {
	input := []byte(getLines("input/18.txt")[0])

	rows := make([][]byte, 0, 400000)
	rows = append(rows, input)

	for len(rows) < 400000 {
		newRow := make([]byte, len(input))
		for i := 0; i < len(newRow); i++ {
			var left, center, right byte
			if i-1 < 0 {
				left = '.'
			} else {
				left = rows[len(rows)-1][i-1]
			}
			center = rows[len(rows)-1][i]
			if i+1 >= len(newRow) {
				right = '.'
			} else {
				right = rows[len(rows)-1][i+1]
			}

			if (left == '^' && center == '^' && right == '.') ||
				(left == '.' && center == '^' && right == '^') ||
				(left == '^' && center == '.' && right == '.') ||
				(left == '.' && center == '.' && right == '^') {
				newRow[i] = '^'
			} else {
				newRow[i] = '.'
			}
		}
		rows = append(rows, newRow)
	}

	fmt.Println("Day 18 Part 1 Result: ", countSafeTiles(rows[:40]))
	fmt.Println("Day 18 Part 2 Result: ", countSafeTiles(rows))
}

func countSafeTiles(tiles [][]byte) int {
	var count int
	for _, row := range tiles {
		for _, col := range row {
			if col == '.' {
				count++
			}
		}
	}
	return count
}
