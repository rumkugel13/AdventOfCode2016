package main

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