package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", performMirageMaintenance("input.txt", false))
	fmt.Println("Part 2: ", performMirageMaintenance("input.txt", true))
}

func performMirageMaintenance(name string, reverse bool) int {
	lines := files.ReadLines(name)
	report := createEnvironmentalReport(lines, reverse)

	sumOfPredicatedValues := 0
	for _, history := range report.histories {
		sumOfPredicatedValues += history.predictNextValue()
	}

	return sumOfPredicatedValues
}

type EnvironmentalReport struct {
	histories []History
}

func createEnvironmentalReport(lines []string, reverse bool) EnvironmentalReport {
	histories := make([]History, 0)
	for _, line := range lines {
		values := ints.FromStringSlice(strings.Fields(line))
		if reverse {
			slices.Reverse(values)
		}
		histories = append(histories, History{values})
	}
	return EnvironmentalReport{histories}
}

type History struct {
	values []int
}

func (h History) predictNextValue() int {
	allHistoryValues := h.extrapolateAndBuildAllHistories()
	updatedHistories := make([][]int, len(allHistoryValues))

	for i := len(allHistoryValues) - 1; i >= 0; i-- {
		historyCopy := make([]int, len(allHistoryValues[i]))
		copy(historyCopy, allHistoryValues[i])

		if ints.AllSame(historyCopy) {
			historyCopy = append(historyCopy, historyCopy[0])
		} else {
			lastValue := historyCopy[len(historyCopy)-1]
			lastValueInRowBelow := updatedHistories[i+1][len(updatedHistories[i+1])-1]
			historyCopy = append(historyCopy, lastValue+lastValueInRowBelow)
		}
		updatedHistories[i] = historyCopy
	}
	return updatedHistories[0][len(updatedHistories[0])-1]
}

func (h History) extrapolateAndBuildAllHistories() [][]int {
	var allHistoryValues [][]int
	allHistoryValues = append(allHistoryValues, h.values)
	currentHistory := h

	for !ints.AllSame(currentHistory.values) {
		newValues := make([]int, len(currentHistory.values)-1)
		for i := 1; i < len(currentHistory.values); i++ {
			newValues[i-1] = currentHistory.values[i] - currentHistory.values[i-1]
		}
		currentHistory = History{newValues}
		allHistoryValues = append(allHistoryValues, currentHistory.values)
	}

	return allHistoryValues
}
