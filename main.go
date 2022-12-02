package main

import (
	"advent-of-code-2022/day01"
	"advent-of-code-2022/day02"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}
