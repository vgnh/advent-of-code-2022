package day06

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const filename = "./day06/input.txt"

var buffer = []rune(utils.ReadString(filename))

func position(length int) int {
	m := make(map[rune]int)
	for i := 0; i < len(buffer); i++ {
		if _, ok := m[buffer[i]]; !ok {
			m[buffer[i]] = i
			if len(m) == length {
				return i + 1
			}
		} else {
			i = m[buffer[i]]
			m = make(map[rune]int)
		}
	}
	return -1
}

func part01() int {
	return position(4)
}

func part02() int {
	return position(14)
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 06")
	fmt.Println(part01())
	fmt.Println(part02())
}
