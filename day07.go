package main

import (
	"fmt"
)

func day07() {
	lines := getLines("input/07.txt")
	var result = 0
	var result2 = 0

	for _, line := range lines {
		if supportsTLS(line) {
			result++
		}
		if supportsSSL(line) {
			result2++
		}
	}

	fmt.Println("Day 07 Part 1 Result: ", result)
	fmt.Println("Day 07 Part 2 Result: ", result2)
}

func supportsTLS(line string) bool {
	inside := false
	valid := false
	invalid := false

	for i := 3; i < len(line); i++ {
		if line[i] == '[' {
			inside = true
		} else if line[i] == ']' {
			inside = false
		}

		a, b, c, d := line[i-3], line[i-2], line[i-1], line[i]

		if !inside {
			if a == b {
				continue
			}
			if a == d && b == c {
				valid = true
			}
		} else {
			if a == b {
				continue
			}
			if a == d && b == c {
				invalid = true
			}
		}
	}

	return valid && !invalid
}

func supportsSSL(line string) bool {
	inside := false
	possible := map[string]bool{}
	inverted := map[string]bool{}

	for i := 2; i < len(line); i++ {
		if line[i] == '[' {
			inside = true
			continue
		} else if line[i] == ']' {
			inside = false
			continue
		}

		a, b, c := line[i-2], line[i-1], line[i]

		if !inside {
			if a == b {
				continue
			}
			if a == c {
				if _, found := inverted[string(b)+string(a)+string(b)]; found {
					return true
				}
				possible[line[i-2:i+1]] = true
			}
		} else {
			if a == b {
				continue
			}
			if a == c {
				if _, found := possible[string(b)+string(a)+string(b)]; found {
					return true
				}
				inverted[line[i-2:i+1]] = true
			}
		}
	}

	return false
}
