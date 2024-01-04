package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day09() {
	input := getLines("09.txt")[0]

	var result = decompress(input, false)
	fmt.Println("Day 09 Part 1 Result: ", result)
	var result2 = decompress(input, true)
	fmt.Println("Day 09 Part 2 Result: ", result2)
}

func decompress(input string, recursive bool) int {
	var length = 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			if idx := strings.IndexByte(input[i:], ')'); idx >= 0 {
				marker := strings.Split(input[i+1:i+idx], "x")
				start := i + idx + 1
				charCount, _ := strconv.Atoi(marker[0])
				repeat, _ := strconv.Atoi(marker[1])

				str := input[start : start+charCount]
				l := len(str)

				if recursive && strings.Contains(str, "(") {
					l = decompress(str, true)
				}

				for r := 0; r < repeat; r++ {
					length += l
				}

				i += idx + charCount
			}
		} else {
			length++
		}
	}
	return length
}
