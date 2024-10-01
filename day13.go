package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func day13() {
	input, _ := strconv.Atoi(getLines("input/13.txt")[0])

	start := Point{1, 1}
	goal := Point{31, 39}

	var visited = make(map[Point]int)
	visited[start] = 0

	var queue []Point
	queue = append(queue, start)

	var dist int = 0

	for len(queue) > 0 {
		currPoint := queue[0]
		queue = queue[1:]
		currDist := visited[currPoint]

		if currPoint == goal {
			dist = currDist
			break
		}

		for _, dir := range []Point{{+1, 0}, {-1, 0}, {0, +1}, {0, -1}} {
			nextPoint := Point{currPoint.x + dir.x, currPoint.y + dir.y}
			if nextPoint.x >= 0 && nextPoint.y >= 0 && !isWall(nextPoint.x, nextPoint.y, input) {
				_, found := visited[nextPoint]
				if !found {
					queue = append(queue, nextPoint)
					visited[nextPoint] = currDist + 1
				}
			}
		}
	}

	fmt.Println("Day 13 Part 1 Result: ", dist)

	// assuming dist > 50, count how many are below
	var count int
	for _, p := range visited {
		if p <= 50 {
			count++
		}
	}
	fmt.Println("Day 13 Part 2 Result: ", count)
}

func isWall(x, y, in int) bool {
	res := x*x + 3*x + 2*x*y + y + y*y
	res += in
	count := bits.OnesCount(uint(res))
	if count%2 == 0 {
		return false
	} else {
		return true
	}
}
