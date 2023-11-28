package main

import "fmt"
import "strings"
import "strconv"

func day01() {
	var instructions = strings.Split(input, ", ")
	var p = Point{0, 0}
	var zero = p
	var dir = 0 // 0 is up
	visit := map[Point]bool{{0, 0}: true}
	var twice Point

	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		amount, _ := strconv.Atoi(inst[1:])
		if inst[0] == 'R' {
			dir = mod(dir+1, 4)
		} else {
			dir = mod(dir-1, 4)
		}

		for i := 0; i < amount; i++ {
			switch dir {
			case 0:
				p.y -= 1
			case 1:
				p.x += 1
			case 2:
				p.y += 1
			case 3:
				p.x -= 1
			}

			before := len(visit)
			visit[p] = true
			if before == len(visit) && twice == zero {
				twice = p
			}
		}
	}

	var result = p.dist()
	var result2 = twice.dist()
	fmt.Println("Day 01 Part 1 Result: ", result)
	fmt.Println("Day 01 Part 2 Result: ", result2)
}

var input = "L5, R1, R3, L4, R3, R1, L3, L2, R3, L5, L1, L2, R5, L1, R5, R1, L4, R1, R3, L4, L1, R2, R5, R3, R1, R1, L1, R1, L1, L2, L1, R2, L5, L188, L4, R1, R4, L3, R47, R1, L1, R77, R5, L2, R1, L2, R4, L5, L1, R3, R187, L4, L3, L3, R2, L3, L5, L4, L4, R1, R5, L4, L3, L3, L3, L2, L5, R1, L2, R5, L3, L4, R4, L5, R3, R4, L2, L1, L4, R1, L3, R1, R3, L2, R1, R4, R5, L3, R5, R3, L3, R4, L2, L5, L1, L1, R3, R1, L4, R3, R3, L2, R5, R4, R1, R3, L4, R3, R3, L2, L4, L5, R1, L4, L5, R4, L2, L1, L3, L3, L5, R3, L4, L3, R5, R4, R2, L4, R2, R3, L3, R4, L1, L3, R2, R1, R5, L4, L5, L5, R4, L5, L2, L4, R4, R4, R1, L3, L2, L4, R3"
