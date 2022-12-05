package day05

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

const filename = "./day05/input.txt"

var input = strings.Split(utils.ReadString(filename), "\n\n")

var instructions = strings.Split(input[1], "\n")

var stacks = func() []utils.Deque[rune] {
	input[0] = strings.ReplaceAll(input[0], "    ", "[-] ")
	input[0] = strings.ReplaceAll(input[0], "  ", " ")
	input[0] = strings.ReplaceAll(input[0], "][", "] [")

	stackInput := strings.Split(input[0], "\n")
	n := strings.Count(stackInput[0], "]")
	stacks := make([]utils.Deque[rune], n)
	for i := range stacks {
		stacks[i] = utils.NewDeque[rune]()
	}
	for _, s := range stackInput[:len(stackInput)-1] {
		s = strings.Trim(strings.TrimSpace(s), "[]")
		r := strings.Split(s, "] [")
		for i, v := range r {
			if v != "-" {
				stacks[i].Enqueue([]rune(v)[0])
			}
		}
	}
	return stacks
}()

func part01() string {
	cpy := make([]utils.Deque[rune], len(stacks))
	copy(cpy, stacks)
	var howMany, from, to int
	for _, s := range instructions {
		fmt.Sscanf(s, "move %d from %d to %d", &howMany, &from, &to)
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
		fmt.Sscanf(s, "move %d from %d to %d", &howMany, &from, &to)
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
