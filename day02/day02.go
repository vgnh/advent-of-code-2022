package day02

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

const filename = "./day02/input.txt"

var strategy = utils.ReadLines(filename)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
)

var points = map[string]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

func draw(opponent string, player string) bool {
	return opponent == player
}

func win(opponent string, player string) bool {
	if (opponent == ROCK && player == PAPER) || (opponent == PAPER && player == SCISSORS) || (opponent == SCISSORS && player == ROCK) {
		return true
	}
	return false
}

func getPoints(opponent string, player string) int {
	if draw(opponent, player) {
		return points[player] + 3
	}
	if win(opponent, player) {
		return points[player] + 6
	}
	return points[player]
}

var alias = map[string]string{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

func part01() int {
	total := 0
	for _, s := range strategy {
		p := strings.Split(s, " ")
		total += getPoints(alias[p[0]], alias[p[1]])
	}
	return total
}

func part02() int {
	opposite := map[string]string{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}

	total := 0
	for _, s := range strategy {
		p := strings.Split(s, " ")
		switch p[1] {
		case "X":
			total += points[opposite[alias[p[0]]]]
		case "Y":
			total += points[alias[p[0]]] + 3
		case "Z":
			total += points[opposite[opposite[alias[p[0]]]]] + 6
		}
	}
	return total
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 02")
	fmt.Println(part01())
	fmt.Println(part02())
}
