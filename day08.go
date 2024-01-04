package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day08() {
	lines := getLines("08.txt")
	screen := make([][]bool, 6, 50)
	for i := 0; i < len(screen); i++ {
		screen[i] = make([]bool, 50)
	}

	var result = 0
	var result2 = "See above"

	for _, line := range lines {
		process(line, &screen)
	}

	result = count(screen)
	show(screen)

	fmt.Println("Day 08 Part 1 Result: ", result)
	fmt.Println("Day 08 Part 2 Result: ", result2)
}

func process(line string, screen *[][]bool) {
	if line[1] == 'e' {
		size := strings.Split(strings.Split(line, " ")[1], "x")
		x, _ := strconv.Atoi(size[0])
		y, _ := strconv.Atoi(size[1])
		rect(x, y, screen)
	} else if line[7] == 'c' {
		rot := strings.Split(strings.Split(line, "=")[1], " by ")
		col, _ := strconv.Atoi(rot[0])
		amount, _ := strconv.Atoi(rot[1])
		rotateCol(col, amount, screen)
	} else {
		rot := strings.Split(strings.Split(line, "=")[1], " by ")
		row, _ := strconv.Atoi(rot[0])
		amount, _ := strconv.Atoi(rot[1])
		rotateRow(row, amount, screen)
	}
}

func rect(dx, dy int, screen *[][]bool) {
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			(*screen)[y][x] = true
		}
	}
}

func rotateCol(col, amount int, screen *[][]bool) {
	duplicate := make([][]bool, len((*screen)))
	for i := range *screen {
		duplicate[i] = make([]bool, len((*screen)[i]))
		copy(duplicate[i], (*screen)[i])
	}

	for y := 0; y < len(*screen); y++ {
		(*screen)[(y+amount)%len(*screen)][col] = duplicate[y][col]
	}
}

func rotateRow(row, amount int, screen *[][]bool) {
	duplicate := make([][]bool, len((*screen)))
	for i := range *screen {
		duplicate[i] = make([]bool, len((*screen)[i]))
		copy(duplicate[i], (*screen)[i])
	}

	for x := 0; x < len((*screen)[0]); x++ {
		(*screen)[row][(x+amount)%len((*screen)[0])] = duplicate[row][x]
	}
}

func count(screen [][]bool) int {
	var c = 0
	for y := 0; y < len(screen); y++ {
		for x := 0; x < len(screen[y]); x++ {
			if screen[y][x] {
				c++
			}
		}
	}
	return c
}

func show(screen [][]bool) {
	for y := 0; y < len(screen); y++ {
		for x := 0; x < len(screen[y]); x++ {
			if screen[y][x] {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
