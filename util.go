package main

import ( 
	"os" 
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) dist() int {
	return dist(*p)
}

func dist(p Point) int {
	return abs(p.x) + abs(p.y)
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}