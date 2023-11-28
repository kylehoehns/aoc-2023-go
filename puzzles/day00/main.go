package main

import (
	"fmt"
	"sort"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
)

type elf struct {
	totalCalories int
}

// Example using 2022 Day 1 https://adventofcode.com/2022/day/1
func main() {
	fmt.Println("Part 1: ", elfWithMostCalories("input.txt"))
	fmt.Println("Part 2: ", sumOfCaloriesCarriedByTopThreeElves("input.txt"))
}

func elfWithMostCalories(name string) int {
	elves := elvesFromFile(name)
	maxCalories := 0
	for _, elf := range elves {
		sum := elf.totalCalories
		if sum > maxCalories {
			maxCalories = sum
		}
	}
	return maxCalories
}

func sumOfCaloriesCarriedByTopThreeElves(name string) int {
	elves := elvesFromFile(name)
	// sort the elves by total calories
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories > elves[j].totalCalories
	})

	// sum the top three elves
	sum := 0
	for i := 0; i < 3; i++ {
		sum += elves[i].totalCalories
	}
	return sum
}

func elvesFromFile(name string) []elf {
	elves := make([]elf, 0)
	paragraphs := files.ReadParagraphs(name)
	for _, group := range paragraphs {
		elves = append(elves, elf{
			totalCalories: ints.ToStringAndSum(group),
		})
	}
	return elves
}
