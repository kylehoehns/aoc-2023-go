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

func part1(name string) int {

	lines := files.ReadLines(name)
	grid := make([][]string, 0)
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	startRow, startCol := findStartingPosition(grid)

	steps := 1

	row, col := determineFirstStep(grid, startRow, startCol)

	priorRow := startRow
	priorCol := startCol
	// from the 2nd position, move along the pipe and keep track of the number of steps we take until we get back to S
	for !(col == startCol && row == startRow) {
		steps++
		nextRow := row
		nextCol := col
		if grid[row][col] == "|" {
			if row > priorRow {
				// if we are moving south, move down
				nextRow++
			} else {
				// if we are moving north, move up
				nextRow--
			}
		} else if grid[row][col] == "-" {
			if col > priorCol {
				// from the left, go right
				nextCol++
			} else {
				// from the right, go left
				nextCol--
			}
		} else if grid[row][col] == "L" {
			// 90 degree connecting north and east
			if priorRow < row {
				// we're coming into this from the top, so go right
				nextCol++
			} else {
				// we're coming into this from the right, so go up
				nextRow--
			}
		} else if grid[row][col] == "J" {
			// 90 degree connecting north and west
			if priorCol < col {
				// we're coming into this from the left, so go up and right
				nextRow--
			} else {
				// we're coming into this from the top, so go down and left
				nextCol--
			}
		} else if grid[row][col] == "F" {
			// 90 degree connecting south and east
			if priorRow > row {
				// from the bottom, go right
				nextCol++
			} else {
				// from the right, go down
				nextRow++
			}
		} else if grid[row][col] == "7" {
			// 90 degree connecting south and west
			if priorRow > row {
				// from the bottom, go up and left
				nextCol--
			} else {
				// from the left, go down and right
				nextRow++
			}
		}

		priorRow = row
		priorCol = col
		row = nextRow
		col = nextCol
	}

	return steps / 2
}

func determineFirstStep(grid [][]string, startRow int, startCol int) (int, int) {
	row := startRow
	col := startCol
	// look at cells around the starting position and see if any of them connect in a way that we can follow them
	if grid[startRow-1][startCol] == "|" {
		// move north
		row--
	} else if grid[startRow+1][startCol] == "|" {
		// move south
		row++
	} else if grid[startRow][startCol-1] == "-" {
		// move west
		col--
	} else if grid[startRow][startCol+1] == "-" {
		// move east
		col++
	} else if grid[startRow+1][startCol] == "L" {
		// 90 degree connecting north and east
		row++
	} else if grid[startRow+1][startCol] == "J" {
		// 90 degree connecting north and west
		row++
	} else if grid[startRow-1][startCol] == "F" {
		// 90 degree connecting south and east
		row--
	} else if grid[startRow-1][startCol] == "7" {
		// 90 degree connecting south and west
		row--
	}
	return row, col
}

func findStartingPosition(grid [][]string) (int, int) {
	// find the location of the first "S"
	var startRow, startCol int
	for row, line := range grid {
		for col, char := range line {
			if char == "S" {
				startRow = row
				startCol = col
			}
		}
	}
	return startRow, startCol
}

func part2(name string) int {
	return 0
}
