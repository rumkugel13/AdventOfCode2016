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

	for i := 0; len(result) < 8 || strings.Contains(string(result2), "z"); i++ {
		hash := md5.Sum([]byte(data + strconv.Itoa(i)))
		str := hex.EncodeToString(hash[:])
		if str[0:5] == "00000" {
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
