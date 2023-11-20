package main

import (
	"fmt"
	"strconv"

	"github.com/kylehoehns/aoc-2023-go/utils/files"
)

// Example using 2020 Day 1 https://adventofcode.com/2020/day/1
func main() {
	lines := files.ReadLines("input.txt")
	var expenses []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, i)
	}
	fmt.Println("Part 1: ", part1(expenses))
	fmt.Println("Part 2: ", part2(expenses))
}

func part1(expenses []int) int {
	for i, first := range expenses {
		for j, second := range expenses {
			if i < j && first+second == 2020 {
				return first * second
			}
		}
	}
	return -1
}

func part2(expenses []int) int {
	for _, first := range expenses {
		for _, second := range expenses {
			for _, third := range expenses {
				if first+second+third == 2020 {
					return first * second * third
				}
			}
		}
	}
	return -1
}
