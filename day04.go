package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day04() {
	lines := getLines("input/04.txt")
	var sectorIds int
	var storage int

	for i := 0; i < len(lines); i++ {
		split := strings.Split(lines[i], "-")
		count := map[byte]int{}
		for j := 0; j < len(split)-1; j++ {
			for k := 0; k < len(split[j]); k++ {
				count[split[j][k]]++
			}
		}

		split2 := strings.Split(split[len(split)-1], "[")
		id, _ := strconv.Atoi(split2[0])
		checksum := strings.Split(split2[1], "]")[0]

		if check(count, checksum) {
			sectorIds += id

			var sector string
			for j := 0; j < len(split)-1; j++ {
				for k := 0; k < len(split[j]); k++ {
					sector += string(rot(split[j][k], id))
				}
				if sector == "northpole" {
					storage = id
				}
			}
		}
	}

	var result = sectorIds
	var result2 = storage
	fmt.Println("Day 04 Part 1 Result: ", result)
	fmt.Println("Day 04 Part 2 Result: ", result2)
}

func check(count map[byte]int, checksum string) bool {
	pairs := make([][2]interface{}, 0, len(count))
	for k, v := range count {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1].(int) == pairs[j][1].(int) {
			return pairs[i][0].(byte) < pairs[j][0].(byte)
		}

		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	for i := 0; i < len(checksum); i++ {
		char := checksum[i]
		if pairs[i][0] != char {
			return false
		}
	}

	return true
}

func rot(c byte, amount int) byte {
	return byte(mod(int(c)-'a'+amount, 26) + 'a')
}
