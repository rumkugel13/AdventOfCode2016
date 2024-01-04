package main

import "fmt"
import "strings"
import "strconv"

func day01() {
	var instructions = strings.Split(getLines("01.txt")[0], ", ")
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
