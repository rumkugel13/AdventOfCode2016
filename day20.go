package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day20() {
	lines := getLines("input/20.txt")
	ranges := parseRanges(lines)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	lowest := 0
	for _, r := range ranges {
		if lowest < r.start-1 {
			break
		}
		if r.end > lowest {
			lowest = r.end
		}
	}

	fmt.Println("Day 20 Part 1 Result: ", lowest+1)

	lastEnd := 0
	count := 0

	for _, r := range ranges {
		if lastEnd < r.start-1 {
			count += (r.start - 1 - lastEnd)
			lastEnd = r.end
			continue
		}
		if r.end > lastEnd {
			lastEnd = r.end
		}
	}

	fmt.Println("Day 20 Part 2 Result: ", count)

}

func parseRanges(lines []string) []Range {
	ranges := make([]Range, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}

type Range struct {
	start, end int
}
