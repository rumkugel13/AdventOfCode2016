package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day12() {
	instructions := getLines("input/12.txt")

	var result = run(instructions, false)
	fmt.Println("Day 12 Part 1 Result: ", result)
	var result2 = run(instructions, true)
	fmt.Println("Day 12 Part 2 Result: ", result2)
}

func run(insts []string, part2 bool) int {
	var regs [4]int
	if part2 {
		regs['c'-'a'] = 1
	}

	for i := 0; i < len(insts); i++ {
		inst := insts[i]
		split := strings.Fields(inst)
		switch inst[0] {
		case 'c':
			dest := split[2][0]-'a'
			src, num := strconv.Atoi(split[1])
			if num == nil {
				regs[dest] = src
			} else {
				reg := split[1][0]-'a'
				regs[dest] = regs[reg]
			}
		case 'i':
			reg := split[1][0]-'a'
			regs[reg]++
		case 'd':
			reg := split[1][0]-'a'
			regs[reg]--
		case 'j':
			imm, convErr := strconv.Atoi(split[1])
			amount, _ := strconv.Atoi(split[2])
			reg := split[1][0]-'a'
			if (convErr != nil && regs[reg] != 0) || (convErr == nil && imm != 0) {
				i += amount - 1
			}
		}
	}

	return regs['a'-'a']
}
