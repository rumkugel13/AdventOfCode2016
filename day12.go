package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day12() {
	instructions := getLines("12.txt")

	var result = run(instructions, false)
	fmt.Println("Day 12 Part 1 Result: ", result)
	var result2 = run(instructions, true)
	fmt.Println("Day 12 Part 2 Result: ", result2)
}

func run(insts []string, part2 bool) int {
	var regs = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	if part2 {
		regs["c"] = 1
	}

	for i := 0; i < len(insts); i++ {
		inst := insts[i]
		split := strings.Fields(inst)
		switch inst[0] {
		case 'c':
			dest := split[2]
			src, num := strconv.Atoi(split[1])
			if num == nil {
				regs[dest] = src
			} else {
				regs[dest] = regs[split[1]]
			}
		case 'i':
			reg := split[1]
			regs[reg]++
		case 'd':
			reg := split[1]
			regs[reg]--
		case 'j':
			reg, num := strconv.Atoi(split[1])
			amount, _ := strconv.Atoi(split[2])
			if (num != nil && regs[split[1]] != 0) || (num == nil && reg != 0) {
				i += amount - 1
			}
		}
	}

	return regs["a"]
}
