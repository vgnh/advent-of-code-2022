package main

import (
	"advent-of-code-2022/day01"
	"advent-of-code-2022/day02"
	"advent-of-code-2022/day03"
	"advent-of-code-2022/day04"
	//"advent-of-code-2022/day05"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()
	day03.Main()
	day04.Main()
	//day05.Main()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}
