package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := files.ReadLines(name)
	sumOfPartNumbers := 0
	numberLocations, symbolLocations := buildNumberAndSymbolLocationsForOnlyStartingNumbers(lines)

	for k, v := range numberLocations {
		keyParts := strings.Split(k, ":")
		row := ints.FromString(keyParts[0])
		column := ints.FromString(keyParts[1])
		numberLength := len(v)

		// check if left or right of the number is a symbol
		if symbolLocations[fmt.Sprintf("%d:%d", row, column-1)] != "" ||
			symbolLocations[fmt.Sprintf("%d:%d", row, column+numberLength)] != "" {
			sumOfPartNumbers += ints.FromString(v)
			continue
		}

		// check all indices +/- 1 on the row above to see if it's a symbol
		for i := column - 1; i < column+numberLength+1; i++ {
			if symbolLocations[fmt.Sprintf("%d:%d", row-1, i)] != "" {
				sumOfPartNumbers += ints.FromString(v)
				continue
			}
		}

		// check all indices +/- 1 on the row below to see if it's a symbol
		for i := column - 1; i < column+numberLength+1; i++ {
			if symbolLocations[fmt.Sprintf("%d:%d", row+1, i)] != "" {
				sumOfPartNumbers += ints.FromString(v)
				continue
			}
		}

	}

	return sumOfPartNumbers
}

func part2(name string) int {
	lines := files.ReadLines(name)
	sumOfGearRatios := 0
	numberLocations, symbolLocations := buildNumberAndSymbolLocationsForAllNumbers(lines)

	for k, v := range symbolLocations {
		if v == "*" {
			keyParts := strings.Split(k, ":")
			row := ints.FromString(keyParts[0])
			column := ints.FromString(keyParts[1])
			matches := make([]string, 0)
			// check 3 above (1 left/right of the current spot)
			above := numberLocations[fmt.Sprintf("%d:%d", row-1, column)]
			if above != "" {
				matches = append(matches, above)
			} else {
				aboveLeft := numberLocations[fmt.Sprintf("%d:%d", row-1, column-1)]
				if aboveLeft != "" {
					matches = append(matches, aboveLeft)
				}
				aboveRight := numberLocations[fmt.Sprintf("%d:%d", row-1, column+1)]
				if aboveRight != "" {
					matches = append(matches, aboveRight)
				}
			}

			// check 3 below (1 left/right of the current spot)
			below := numberLocations[fmt.Sprintf("%d:%d", row+1, column)]
			if below != "" {
				matches = append(matches, below)
			} else {
				belowLeft := numberLocations[fmt.Sprintf("%d:%d", row+1, column-1)]
				if belowLeft != "" {
					matches = append(matches, belowLeft)
				}
				belowRight := numberLocations[fmt.Sprintf("%d:%d", row+1, column+1)]
				if belowRight != "" {
					matches = append(matches, belowRight)
				}
			}

			left := numberLocations[fmt.Sprintf("%d:%d", row, column-1)]
			if left != "" {
				matches = append(matches, left)
			}

			right := numberLocations[fmt.Sprintf("%d:%d", row, column+1)]
			if right != "" {
				matches = append(matches, right)
			}

			if len(matches) == 2 {
				sumOfGearRatios += ints.FromString(matches[0]) * ints.FromString(matches[1])
			}
		}
	}

	return sumOfGearRatios
}

func buildNumberAndSymbolLocationsForAllNumbers(lines []string) (map[string]string, map[string]string) {

	digitsRegex, _ := regexp.Compile(`\d+`)
	symbolsRegex, _ := regexp.Compile(`[^\d.]+`)

	numberLocations := make(map[string]string)
	symbolLocations := make(map[string]string)
	for i, line := range lines {
		for _, digitLocation := range digitsRegex.FindAllStringIndex(line, -1) {
			partNumber := line[digitLocation[0]:digitLocation[1]]
			for j := 0; j < len(partNumber); j++ {
				numberLocations[fmt.Sprintf("%d:%d", i, digitLocation[0]+j)] = partNumber
			}
		}
		for _, symbolLocation := range symbolsRegex.FindAllStringIndex(line, -1) {
			symbolLocations[fmt.Sprintf("%d:%d", i, symbolLocation[0])] = line[symbolLocation[0]:symbolLocation[1]]
		}
	}

	return numberLocations, symbolLocations
}

func buildNumberAndSymbolLocationsForOnlyStartingNumbers(lines []string) (map[string]string, map[string]string) {

	digitsRegex, _ := regexp.Compile(`\d+`)
	symbolsRegex, _ := regexp.Compile(`[^\d.]+`)

	numberLocations := make(map[string]string)
	symbolLocations := make(map[string]string)
	for i, line := range lines {
		for _, digitLocation := range digitsRegex.FindAllStringIndex(line, -1) {
			partNumber := line[digitLocation[0]:digitLocation[1]]
			numberLocations[fmt.Sprintf("%d:%d", i, digitLocation[0])] = partNumber
		}
		for _, symbolLocation := range symbolsRegex.FindAllStringIndex(line, -1) {
			symbolLocations[fmt.Sprintf("%d:%d", i, symbolLocation[0])] = line[symbolLocation[0]:symbolLocation[1]]
		}
	}

	return numberLocations, symbolLocations
}
