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

	{
		var visited = make(map[Point]int)
		visited[start] = 0

		var todo []Point
		todo = append(todo, start)

		var dist int = 0

		for len(todo) > 0 {
			current := todo[0]
			todo = todo[1:]
			currDist := visited[current]

			if current == goal {
				dist = currDist
				break
			}

			for _, p := range []Point{{+1, 0}, {-1, 0}, {0, +1}, {0, -1}} {
				next := Point{current.x + p.x, current.y + p.y}
				if next.x >= 0 && next.y >= 0 && !isWall(next.x, next.y, input) {
					_, found := visited[next]
					if !found {
						todo = append(todo, next)
						visited[next] = currDist + 1
					}
				}
			}
		}

		fmt.Println("Day 13 Part 1 Result: ", dist)
	}

	{
		var visited2 = make(map[Point]int)
		visited2[start] = 0

		var todo2 []Point
		todo2 = append(todo2, start)

		for len(todo2) > 0 {
			current := todo2[0]
			todo2 = todo2[1:]
			currDist := visited2[current]

			if currDist == 50 {
				continue
			}

			for _, p := range []Point{{+1, 0}, {-1, 0}, {0, +1}, {0, -1}} {
				next := Point{current.x + p.x, current.y + p.y}
				if next.x >= 0 && next.y >= 0 && !isWall(next.x, next.y, input) {
					_, found := visited2[next]
					if !found {
						todo2 = append(todo2, next)
						visited2[next] = currDist + 1
					}
				}
			}
		}

		fmt.Println("Day 13 Part 2 Result: ", len(visited2))
	}
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
