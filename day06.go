package main

import (
	"fmt"
	"sort"
)

func day06() {
	lines := getLines("day06.test")
	counts := []map[byte]int{}
	for i := 0; i < len(lines[0]); i++ {
		counts = append(counts, map[byte]int{})
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			counts[j][line[j]]++
		}
	}

	var result = decode(counts, false)
	var result2 = decode(counts, true)
	fmt.Println("Day 06 Part 1 Result: ", result)
	fmt.Println("Day 06 Part 2 Result: ", result2)
}

func decode(counts []map[byte]int, part2 bool) string {
	var result []byte
	for i := 0; i < len(counts); i++ {

		pairs := make([][2]interface{}, 0, len(counts[i]))
		for k, v := range counts[i] {
			pairs = append(pairs, [2]interface{}{k, v})
		}

		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][1].(int) > pairs[j][1].(int)
		})

		if !part2 {
			result = append(result, pairs[0][0].(byte))
		} else {
			result = append(result, pairs[len(pairs)-1][0].(byte))
		}
	}
	return string(result)
}
