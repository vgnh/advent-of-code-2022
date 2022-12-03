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
		r1 := r[:len(r)/2]
		r2 := r[len(r)/2:]
		itemsInR1 := make(map[rune]struct{})
		for _, c := range r1 {
			itemsInR1[c] = struct{}{}
		}
		for _, c := range r2 {
			if _, ok := itemsInR1[c]; ok {
				if int(c)-64 > 26 {
					sum += int(c) - 64 - 32
				} else {
					sum += int(c) - 64 + 26
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
		itemsInR1 := make(map[rune]struct{})
		for _, c := range rucksacks[i] {
			itemsInR1[c] = struct{}{}
		}
		itemsInR1R2 := make(map[rune]struct{})
		for _, c := range rucksacks[i+1] {
			if _, ok := itemsInR1[c]; ok {
				itemsInR1R2[c] = struct{}{}
			}
		}
		for _, c := range rucksacks[i+2] {
			if _, ok := itemsInR1R2[c]; ok {
				if int(c)-64 > 26 {
					sum += int(c) - 64 - 32
				} else {
					sum += int(c) - 64 + 26
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
