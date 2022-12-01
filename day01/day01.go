package day01

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strings"
)

const filename = "./day01/input.txt"

var calories = func() []int {
	elves := strings.Split(utils.ReadString(filename), "\n\n")
	calories := make([]int, len(elves))
	for i, elf := range elves {
		calories[i] = utils.Sum(utils.MapToInt(strings.Split(elf, "\n")))
	}
	sort.Ints(calories)
	return calories
}()

func part01() int {
	return calories[len(calories)-1]
}

func part02() int {
	return utils.Sum(calories[len(calories)-3:])
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 01")
	fmt.Println(part01())
	fmt.Println(part02())
}
