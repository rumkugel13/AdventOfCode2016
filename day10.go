package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Inst struct {
	low, high       int
	lowout, highout bool
}

func day10() {
	lines := getLines("input10.txt")
	bots := map[int][]int{}
	outputs := map[int][]int{}
	insts := map[int]Inst{}
	hasTwo := []int{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line[0] == 'v' {
			split := strings.Split(line, " ")
			value, _ := strconv.Atoi(split[1])
			bot, _ := strconv.Atoi(strings.TrimSpace(split[5]))
			bots[bot] = append(bots[bot], value)
			if len(bots[bot]) == 2 {
				hasTwo = append(hasTwo, bot)
			}
		} else {
			split := strings.Split(line, " ")
			bot, _ := strconv.Atoi(split[1])
			low, _ := strconv.Atoi(split[6])
			lowout := split[5] == "output"
			high, _ := strconv.Atoi(strings.TrimSpace(split[11]))
			highout := split[10] == "output"
			insts[bot] = Inst{low, high, lowout, highout}
		}
	}
	var result = 0

	for len(hasTwo) > 0 {
		current := hasTwo[len(hasTwo)-1]
		hasTwo = hasTwo[:len(hasTwo)-1]

		inst := insts[current]
		low := min(bots[current][0], bots[current][1])
		high := max(bots[current][0], bots[current][1])

		if inst.lowout {
			outputs[inst.low] = append(outputs[inst.low], low)
		} else {
			bots[inst.low] = append(bots[inst.low], low)
			if len(bots[inst.low]) == 2 {
				hasTwo = append(hasTwo, inst.low)
			}
		}
		if inst.highout {
			outputs[inst.high] = append(outputs[inst.high], high)
		} else {
			bots[inst.high] = append(bots[inst.high], high)
			if len(bots[inst.high]) == 2 {
				hasTwo = append(hasTwo, inst.high)
			}
		}
		bots[current] = []int{}
		if (low == 61 && high == 17) || (low == 17 && high == 61) {
			result = current
		}
	}

	fmt.Println("Day 10 Part 1 Result: ", result)
	var result2 = outputs[0][0] * outputs[1][0] * outputs[2][0]
	fmt.Println("Day 10 Part 2 Result: ", result2)
}
