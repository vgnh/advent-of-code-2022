package day04

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const filename = "./day04/input.txt"

var pair = func() [][]int {
	input := utils.ReadLines(filename)
	pair := make([][]int, len(input))
	for i, s := range input {
		pair[i] = make([]int, 4)
		fmt.Sscanf(s, "%d-%d,%d-%d", &pair[i][0], &pair[i][1], &pair[i][2], &pair[i][3])
	}
	return pair
}()

func part01() int {
	count := 0
	for _, p := range pair {
		if (p[0] <= p[2] && p[1] >= p[3]) || (p[0] >= p[2] && p[1] <= p[3]) {
			count++
		}
	}
	return count
}

func part02() int {
	count := 0
	for _, p := range pair {
		if !((p[1] < p[2]) || (p[3] < p[0])) {
			count++
		}
	}
	return count
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 04")
	fmt.Println(part01())
	fmt.Println(part02())
}
