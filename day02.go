package main

import "fmt"
import "strconv"

func day02() {
	var instructions = getLines("input/02.txt")
	var num = 5
	var num2 = 5
	var result string
	var result2 string

	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		for j := 0; j < len(inst); j++ {
			dir := inst[j]
			switch dir {
			case 'U':
				if num > 3 {
					num -= 3
				}
				if num2 != 5 && num2 != 2 && num2 != 1 && num2 != 4 && num2 != 9 {
					if num2%2 == 0 || num2 == 7 || num2 == 11 {
						num2 -= 4
					} else {
						num2 -= 2
					}
				}
			case 'D':
				if num < 7 {
					num += 3
				}
				if num2 != 5 && num2 != 10 && num2 != 13 && num2 != 12 && num2 != 9 {
					if num2%2 == 0 || num2 == 7 || num2 == 3 {
						num2 += 4
					} else {
						num2 += 2
					}
				}
			case 'L':
				if num%3 != 1 {
					num -= 1
				}
				if num2 != 5 && num2 != 2 && num2 != 1 && num2 != 10 && num2 != 13 {
					num2 -= 1
				}
			case 'R':
				if num%3 != 0 {
					num += 1
				}
				if num2 != 9 && num2 != 4 && num2 != 1 && num2 != 12 && num2 != 13 {
					num2 += 1
				}
			}
		}
		result += strconv.Itoa(num)
		result2 += strconv.FormatInt(int64(num2), 16)
	}

	fmt.Println("Day 02 Part 1 Result: ", result)
	fmt.Println("Day 02 Part 2 Result: ", result2)
}
