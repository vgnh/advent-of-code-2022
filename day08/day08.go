package day08

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
)

const filename = "./day08/input.txt"

var grid = func() [][]int {
	input := utils.ReadLines(filename)
	grid := make([][]int, len(input))
	for i := range input {
		grid[i] = make([]int, len(input[i]))
		for j, r := range []rune(input[i]) {
			grid[i][j], _ = strconv.Atoi(string(r))
		}
	}
	return grid
}()

func visible(grid [][]int, r, c int) bool {
	// top
	top := true
	for i := 0; i < r; i++ {
		if grid[i][c] >= grid[r][c] {
			top = false
			break
		}
	}

	// bottom
	bottom := true
	for i := r + 1; i < len(grid); i++ {
		if grid[i][c] >= grid[r][c] {
			bottom = false
			break
		}
	}

	// left
	left := true
	for j := 0; j < c; j++ {
		if grid[r][j] >= grid[r][c] {
			left = false
			break
		}
	}

	// right
	right := true
	for j := c + 1; j < len(grid[0]); j++ {
		if grid[r][j] >= grid[r][c] {
			right = false
			break
		}
	}

	return top || bottom || left || right
}

func part01() int {
	count := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if visible(grid, i, j) {
				count++
			}
		}
	}
	return count + len(grid[0])*2 + len(grid)*2 - 4
}

func scenicScore(grid [][]int, r, c int) int {
	// top
	top := 0
	for i := r - 1; i >= 0; i-- {
		top++
		if grid[i][c] >= grid[r][c] {
			break
		}
	}

	// bottom
	bottom := 0
	for i := r + 1; i < len(grid); i++ {
		bottom++
		if grid[i][c] >= grid[r][c] {
			break
		}
	}

	// left
	left := 0
	for j := c - 1; j >= 0; j-- {
		left++
		if grid[r][j] >= grid[r][c] {
			break
		}
	}

	// right
	right := 0
	for j := c + 1; j < len(grid[0]); j++ {
		right++
		if grid[r][j] >= grid[r][c] {
			break
		}
	}

	return top * bottom * left * right
}

func part02() int {
	highest := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			s := scenicScore(grid, i, j)
			if s > highest {
				highest = s
			}
		}
	}
	return highest
}

func Main() {
	fmt.Println("Advent of Code 2022, Day 08")
	fmt.Println(part01())
	fmt.Println(part02())
}
