package main

import (
	"fmt"
	"strconv"
)

func day19() {
	input, _ := strconv.Atoi(getLines("input/19.txt")[0])

	elves := make([]int, input)
	for i := range elves {
		elves[i] = i + 1
	}

	var result int
outer:
	for round := 0; round < input; round++ {
		for i := 0; i < len(elves); i++ {
			if elves[i] == 0 {
				continue
			}
			next := findNext(elves, i)
			if next == i {
				result = i + 1
				break outer
			}

			elves[next] = 0
		}
	}

	fmt.Println("Day 19 Part 1 Result: ", result)
}

func findNext(presents []int, start int) int {
	for i := start + 1; i < start+len(presents); i++ {
		if presents[(i)%len(presents)] != 0 {
			return i % len(presents)
		}
	}
	return start
}
