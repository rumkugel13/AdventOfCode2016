package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type maybe struct {
	char      byte
	used      bool
	iteration int
}

func day14() {
	var input = getLines("input/14.txt")[0]
	var data = []byte(input)

	for part := 0; part < 2; part++ {
		var maybes []maybe
		var keyIndex []int
		for i := 0; len(keyIndex) < 64; i++ {
			hash := md5.Sum(append(data, []byte(strconv.Itoa(i))...))
			str := hex.EncodeToString(hash[:])
			if part == 1 {
				str = stretch(str)
			}

			for maybeIdx := 0; maybeIdx < len(maybes); maybeIdx++ {
				maybe := maybes[maybeIdx]
				if (maybe.iteration+1000 > i) && !maybe.used {
					found5 := findFive(str, maybe.char)
					if found5 {
						keyIndex = append(keyIndex, maybe.iteration)
						maybes[maybeIdx].used = true
					}
				}
			}

			char, found3 := findThree(str)
			if found3 {
				maybes = append(maybes, maybe{char, false, i})
			}
		}
		fmt.Println("Day 14 Part", part+1, "Result:", keyIndex[63])
	}
}

func stretch(str string) string {
	for i := 0; i < 2016; i++ {
		hash := md5.Sum([]byte(str))
		str = hex.EncodeToString(hash[:])
	}
	return str
}

func findThree(hash string) (byte, bool) {
	count := 1
	char := hash[0]
	for i := 1; i < len(hash); i++ {
		if hash[i] == char {
			count++
		} else {
			char = hash[i]
			count = 1
		}

		if count == 3 {
			return char, true
		}
	}

	return 0, false
}

func findFive(hash string, char byte) bool {
	count := 0
	for i := 0; i < len(hash); i++ {
		if hash[i] == char {
			count++
		} else {
			count = 0
		}

		if count == 5 {
			return true
		}
	}

	return false
}
