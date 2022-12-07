package day06

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const filename = "./day06/input.txt"

var datastreamBuffer = utils.ReadString(filename)

func markerPosition(markerLength int) int {
	for i := 0; i <= len(datastreamBuffer)-markerLength; i++ {
		s := make(utils.Set[rune])
		s.AddAll([]rune(datastreamBuffer[i : i+markerLength]))
		if len(s) == markerLength {
			return i + markerLength
		}
	}
	return -1
}

func part01() int {
	return markerPosition(4)
}

func part02() int {
	return markerPosition(14)
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 06")
	fmt.Println(part01())
	fmt.Println(part02())
}
