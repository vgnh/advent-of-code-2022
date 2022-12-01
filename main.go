package main

import (
	"advent-of-code-2022/day01"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}
