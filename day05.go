package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func day05() {
	var data = getLines("input/05.txt")[0]
	var result string
	var result2 = []byte("zzzzzzzz")

	var doorID = []byte(data)
	for i := 0; len(result) < 8 || strings.Contains(string(result2), "z"); i++ {
		hash := md5.Sum(append(doorID, []byte(strconv.Itoa(i))...))
		if hash[0] == 0 && hash[1] == 0 && hash[2] <= 0x0F {
			str := hex.EncodeToString(hash[:])
			if len(result) < 8 {
				result += string(str[5])
			}
			index := str[5]
			if index >= '0' && index <= '7' && result2[index-'0'] == 'z' {
				result2[index-'0'] = str[6]
			}
		}
	}

	fmt.Println("Day 05 Part 1 Result: ", result)
	fmt.Println("Day 05 Part 2 Result: ", string(result2))
}
