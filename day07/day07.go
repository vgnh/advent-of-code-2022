package day07

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const filename = "./day07/input.txt"

var terminalOutput = utils.ReadLines(filename)

var dirs = func() map[string][]string {
	dirs := make(map[string][]string)
	parent := ""
	for _, cmd := range terminalOutput {
		if strings.HasPrefix(cmd, "$ cd ") {
			d := cmd[5:]
			if d == ".." {
				if parent == "/" {
					continue
				}
				li := strings.LastIndex(parent, "/")
				if li == 0 {
					parent = "/"
				} else {
					parent = parent[:li]
				}
			} else {
				if d == "/" {
					parent = "/"
				} else {
					if parent == "/" {
						parent = parent + d
					} else {
						parent = parent + "/" + d
					}
				}
			}
		} else if cmd == "$ ls" {
			continue
		} else {
			dirs[parent] = append(dirs[parent], cmd)
		}
	}
	return dirs
}()

var folderSizes = make(map[string]int)

func folderSize(parent string) int {
	if size, ok := folderSizes[parent]; ok {
		return int(size)
	}

	sum := 0
	for _, v := range dirs[parent] {
		if strings.HasPrefix(v, "dir ") {
			subfolder := strings.Split(v, " ")
			if parent == "/" {
				subfolder[1] = "/" + subfolder[1]
			} else {
				subfolder[1] = parent + "/" + subfolder[1]
			}
			sum += folderSize(subfolder[1])
		} else {
			file := strings.Split(v, " ")
			num, _ := strconv.Atoi(file[0])
			sum += num
		}
	}
	folderSizes[parent] = sum
	return sum
}

func part01() int {
	sum := 0
	for _, v := range folderSizes {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func part02() int {
	space := 70000000 - folderSizes["/"]
	need := 30000000 - space
	deletable := 30000000
	for _, v := range folderSizes {
		if v >= need {
			if v < deletable {
				deletable = v
			}
		}
	}
	return deletable
}

func Main() {
	folderSize("/")
	fmt.Println("Advent of Code 2022, Day 07")
	fmt.Println(part01())
	fmt.Println(part02())
}
