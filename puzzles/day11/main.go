package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

const galaxy = "#"

type coordinate struct {
	x, y int
}

func part1(name string) int {
	lines := expandUniverse(files.ReadLines(name))

	// find all coordinates where a galaxy is present
	galaxyCoordinates := make([]coordinate, 0)
	for y, line := range lines {
		for x, char := range line {
			if string(char) == galaxy {
				galaxyCoordinates = append(galaxyCoordinates, coordinate{x, y})
			}
		}
	}

	sumOfShortestPaths := 0
	// determine distance between each galaxy
	for _, coord := range galaxyCoordinates {
		for _, otherCoord := range galaxyCoordinates {
			if coord == otherCoord {
				continue
			}
			// find the shortest path between two coordinates
			sumOfShortestPaths += ints.Abs(coord.x-otherCoord.x) + ints.Abs(coord.y-otherCoord.y)
		}
	}

	// divide by 2 because we are double counting paths between galaxies, only need to count one of the paths
	return sumOfShortestPaths / 2
}

func expandUniverse(lines []string) []string {
	// loop through rows, add another row if it has no galaxies
	newLines := make([]string, 0)
	for _, line := range lines {
		newLines = append(newLines, line)
		if !strings.Contains(line, galaxy) {
			newLines = append(newLines, strings.Repeat(".", len(line)))
		}
	}

	// find all columns with no galaxies
	columnsWithoutGalaxies := make([]int, 0)
	for i := 0; i < len(lines[0]); i++ {
		columnHasGalaxy := false
		for _, line := range lines {
			if string(line[i]) == galaxy {
				columnHasGalaxy = true
				break
			}
		}
		if !columnHasGalaxy {
			columnsWithoutGalaxies = append(columnsWithoutGalaxies, i)
		}
	}

	// figure out which columns need an extra column
	columnsToAdd := make([]int, 0)
	columnsAdded := 0
	for _, column := range columnsWithoutGalaxies {
		columnsToAdd = append(columnsToAdd, column+columnsAdded)
		columnsAdded++
	}

	// add columns
	linesWithColumns := make([]string, len(newLines))
	copy(linesWithColumns, newLines)
	for i := range newLines {
		for _, column := range columnsToAdd {
			if column < len(linesWithColumns[i]) {
				linesWithColumns[i] = linesWithColumns[i][:column+1] + "." + linesWithColumns[i][column+1:]
			} else {
				linesWithColumns[i] = linesWithColumns[i] + "."
			}
		}
	}
	return linesWithColumns
}

func part2(name string) int {
	return 0
}
