package main

import (
	"advent-of-code-2022/day01"
	"advent-of-code-2022/day02"
	"advent-of-code-2022/day03"
	"advent-of-code-2022/day04"
	"advent-of-code-2022/day05"
	"advent-of-code-2022/day06"
	"advent-of-code-2022/day07"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()
	day03.Main()
	day04.Main()
	day05.Main()
	day06.Main()
	day07.Main()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}
