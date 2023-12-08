package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

type Node struct {
	value string
	left  string
	right string
}

type PuzzleInput struct {
	instructions []string
	nodes        map[string]Node
}

func part1(name string) int {
	paragraphs := files.ReadParagraphs(name)

	input := parsePuzzleInput(paragraphs)

	steps := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		node := input.nodes[currentNode]
		instructionIndex := steps % len(input.instructions)
		direction := input.instructions[instructionIndex]
		if direction == "L" {
			currentNode = node.left
		} else {
			currentNode = node.right
		}

		steps++
	}
	return steps
}

func parsePuzzleInput(paragraphs [][]string) PuzzleInput {
	nodes := make(map[string]Node)
	for _, line := range paragraphs[1] {
		node := Node{
			value: line[:3],
			left:  line[7:10],
			right: line[12:15],
		}
		nodes[node.value] = node
	}

	instructions := strings.Split(paragraphs[0][0], "")
	return PuzzleInput{
		instructions: instructions,
		nodes:        nodes,
	}
}

func part2(name string) int {
	paragraphs := files.ReadParagraphs(name)

	input := parsePuzzleInput(paragraphs)

	nodesThatEndWithA := make([]string, 0)
	for _, node := range input.nodes {
		if strings.HasSuffix(node.value, "A") {
			nodesThatEndWithA = append(nodesThatEndWithA, node.value)
		}
	}

	allNodeSteps := 1
	for _, currentNode := range nodesThatEndWithA {

		nodeSteps := 0
		for !strings.HasSuffix(currentNode, "Z") {
			node := input.nodes[currentNode]
			instructionIndex := nodeSteps % len(input.instructions)
			direction := input.instructions[instructionIndex]
			if direction == "L" {
				currentNode = node.left
			} else {
				currentNode = node.right
			}
			nodeSteps++
		}
		allNodeSteps = leastCommonMultiple(allNodeSteps, nodeSteps)

	}
	return allNodeSteps
}

func leastCommonMultiple(x int, y int) int {
	return x * y / greatestCommonDenominator(x, y)
}

func greatestCommonDenominator(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
