package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day03() {
	lines := getLines("day03.test")
	var possible int

	for i := 0; i < len(lines); i++ {
		parts := strings.Fields(lines[i])
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		if a+b > c && b+c > a && c+a > b {
			possible++
		}
	}

	var possible2 int
	for i := 0; i < len(lines); i += 3 {
		parts := strings.Fields(lines[i])
		a1, _ := strconv.Atoi(parts[0])
		b1, _ := strconv.Atoi(parts[1])
		c1, _ := strconv.Atoi(parts[2])

		parts2 := strings.Fields(lines[i+1])
		a2, _ := strconv.Atoi(parts2[0])
		b2, _ := strconv.Atoi(parts2[1])
		c2, _ := strconv.Atoi(parts2[2])

		parts3 := strings.Fields(lines[i+2])
		a3, _ := strconv.Atoi(parts3[0])
		b3, _ := strconv.Atoi(parts3[1])
		c3, _ := strconv.Atoi(parts3[2])

		if a1+a2 > a3 && a2+a3 > a1 && a3+a1 > a2 {
			possible2++
		}
		if b1+b2 > b3 && b2+b3 > b1 && b3+b1 > b2 {
			possible2++
		}
		if c1+c2 > c3 && c2+c3 > c1 && c3+c1 > c2 {
			possible2++
		}
	}

	var result = possible
	var result2 = possible2
	fmt.Println("Day 03 Part 1 Result: ", result)
	fmt.Println("Day 03 Part 2 Result: ", result2)
}
