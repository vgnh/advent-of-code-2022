package day03

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const filename = "./day03/input.txt"

var rucksacks = utils.ReadLines(filename)

func part01() int {
	sum := 0
	for _, r := range rucksacks {
		c1 := make(utils.Set[rune])
		c1.AddAll([]rune(r[:len(r)/2]))

		for _, char := range r[len(r)/2:] {
			if c1.Contains(char) {
				if int(char)-64 > 26 {
					sum += int(char) - 64 - 32
				} else {
					sum += int(char) - 64 + 26
				}
				break
			}
		}
	}
	return sum
}

func part02() int {
	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		r1 := make(utils.Set[rune])
		r1.AddAll([]rune(rucksacks[i]))

		r1r2 := r1.RetainAll([]rune(rucksacks[i+1]))

		for _, char := range rucksacks[i+2] {
			if r1r2.Contains(char) {
				if int(char)-64 > 26 {
					sum += int(char) - 64 - 32
				} else {
					sum += int(char) - 64 + 26
				}
				break
			}
		}
	}
	return sum
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 03")
	fmt.Println(part01())
	fmt.Println(part02())
}
