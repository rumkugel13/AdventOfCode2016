package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day15() {
	lines := getLines("input/15.txt")
	discs := parseDiscs(lines)

	time := runDiscs(discs)
	fmt.Println("Day 15 Part 1 Result: ", time)

	discs = append(discs, disc{11, 0})
	time = runDiscs(discs)
	fmt.Println("Day 15 Part 2 Result: ", time)
}

func runDiscs(discs []disc) int {
outer:
	for time := 0; ; time++ {
		for num, disc := range discs {
			if (disc.start+time+1+num)%disc.positions != 0 {
				continue outer
			}
		}
		return time
	}
}

type disc struct {
	positions, start int
}

func parseDiscs(lines []string) []disc {
	var discs []disc
	for _, line := range lines {
		parts := strings.Fields(line)
		positions, _ := strconv.Atoi(parts[3])
		start, _ := strconv.Atoi(parts[11][:len(parts[11])-1])
		discs = append(discs, disc{positions, start})
	}
	return discs
}
