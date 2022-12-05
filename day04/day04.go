package day04

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const filename = "./day04/input.txt"

type elf struct {
	one []int
	two []int
}

var elfPairs = func() []elf {
	assignmentPairs := utils.ReadLines(filename)
	elfPairs := make([]elf, len(assignmentPairs))
	for i, ap := range assignmentPairs {
		pair := strings.Split(ap, ",")
		elf1Str := strings.Split(pair[0], "-")
		e1Low, _ := strconv.Atoi(elf1Str[0])
		e1High, _ := strconv.Atoi(elf1Str[1])
		elf2Str := strings.Split(pair[1], "-")
		e2Low, _ := strconv.Atoi(elf2Str[0])
		e2High, _ := strconv.Atoi(elf2Str[1])
		elfPairs[i] = elf{
			one: []int{e1Low, e1High},
			two: []int{e2Low, e2High},
		}
	}
	return elfPairs
}()

func part01() int {
	count := 0
	for _, elf := range elfPairs {
		if (elf.one[0] <= elf.two[0] && elf.one[1] >= elf.two[1]) || (elf.one[0] >= elf.two[0] && elf.one[1] <= elf.two[1]) {
			count++
		}
	}
	return count
}

func part02() int {
	count := 0
	for _, elf := range elfPairs {
		if !((elf.one[1] < elf.two[0]) || (elf.two[1] < elf.one[0])) {
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
