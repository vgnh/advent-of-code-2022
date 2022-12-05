package day05

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const filename = "./day05/input.txt"

var instructions = utils.ReadLines(filename)

var stacks = func() []utils.Deque[rune] {
	initialStates := []string{
		"WBDNCFJ",
		"PZVQLST",
		"PZBGJT",
		"DTLJZBHC",
		"GVBJS",
		"PSQ",
		"BVDFLMPN",
		"PSMFBDLR",
		"VDTR",
	}
	stacks := make([]utils.Deque[rune], len(initialStates))
	for i, s := range initialStates {
		stacks[i] = utils.NewDeque[rune]()
		for _, r := range s {
			stacks[i].Push(r)
		}
	}
	return stacks
}()

func part01() string {
	cpy := make([]utils.Deque[rune], len(stacks))
	copy(cpy, stacks)
	var howMany, from, to int
	for _, s := range instructions {
		_, err := fmt.Sscanf(s, "move %d from %d to %d", &howMany, &from, &to)
		if err != nil {
			continue
		}
		from--
		to--
		for i := 0; i < howMany; i++ {
			popped, _ := cpy[from].Pop()
			cpy[to].Push(popped)
		}
	}
	s := ""
	for _, stack := range cpy {
		s += string(stack.Peek())
	}
	return s
}

func part02() string {
	var howMany, from, to int
	for _, s := range instructions {
		_, err := fmt.Sscanf(s, "move %d from %d to %d", &howMany, &from, &to)
		if err != nil {
			continue
		}
		from--
		to--
		helperStack := utils.NewDeque[rune]()
		for i := 0; i < howMany; i++ {
			popped, _ := stacks[from].Pop()
			helperStack.Push(popped)
		}
		for i := 0; i < howMany; i++ {
			popped, _ := helperStack.Pop()
			stacks[to].Push(popped)
		}
	}
	s := ""
	for _, stack := range stacks {
		s += string(stack.Peek())
	}
	return s
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 05")
	fmt.Println(part01())
	fmt.Println(part02())
}
