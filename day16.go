package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func day16() {
	input := []byte(getLines("input/16.txt")[0])

	data := input
	diskSize := 272
	data = checksum(data, diskSize)
	fmt.Println("Day 16 Part 1 Result: ", string(data))

	data = input
	diskSize = 35651584
	data = checksum(data, diskSize)
	fmt.Println("Day 16 Part 2 Result: ", string(data))
}

func checksum(data []byte, diskSize int) []byte {
	for len(data) < diskSize {
		data = expandInput(data)
	}

	data = data[:diskSize]

	for len(data)%2 == 0 {
		data = hash(data)
	}
	return data
}

func hash(data []byte) []byte {
	var new []byte
	for i := 0; i < len(data)-1; i += 2 {
		a := data[i]
		b := data[i+1]
		if a == b {
			new = append(new, '1')
		} else {
			new = append(new, '0')
		}
	}
	return new
}

func expandInput(input []byte) []byte {
	var a, b []byte = make([]byte, len(input)), make([]byte, len(input))
	copy(a, input)
	copy(b, input)
	reverse(b)
	toggle(b)
	a = append(a, '0')
	a = append(a, b...)
	return a
}

func toggle(arr []byte) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == '0' {
			arr[i] = '1'
		} else {
			arr[i] = '0'
		}
	}
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
