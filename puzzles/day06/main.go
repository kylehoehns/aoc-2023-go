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

type Race struct {
	time           int
	recordDistance int
}

func (r Race) determineWaysToBeatRecord() int {
	waysToBeat := 0
	for speed := 0; speed <= r.time; speed++ {
		timeToRace := r.time - speed
		distance := speed * timeToRace
		if distance > r.recordDistance {
			waysToBeat++
		}
	}
	return waysToBeat
}

func part1(name string) int {
	lines := files.ReadLines(name)

	races := parseInput(lines)

	waysToBeatRecord := 0
	for _, race := range races {
		raceWaysToBeatRecord := race.determineWaysToBeatRecord()
		if waysToBeatRecord == 0 {
			waysToBeatRecord = raceWaysToBeatRecord
		} else {
			waysToBeatRecord *= raceWaysToBeatRecord
		}
	}

	return waysToBeatRecord
}

func part2(name string) int {
	lines := files.ReadLines(name)

	race := parseInputPart2(lines)

	return race.determineWaysToBeatRecord()
}

func parseInput(lines []string) []Race {
	timesRow, distanceRow := getTimesAndDistances(lines)

	races := make([]Race, 0)
	for i := 0; i < len(timesRow); i++ {
		races = append(races, Race{
			time:           ints.FromString(timesRow[i]),
			recordDistance: ints.FromString(distanceRow[i]),
		})
	}
	return races
}

func parseInputPart2(lines []string) Race {
	timesRow, distanceRow := getTimesAndDistances(lines)

	totalTime := ""
	totalDistance := ""
	for i := 0; i < len(timesRow); i++ {
		totalTime += timesRow[i]
		totalDistance += distanceRow[i]
	}
	return Race{
		time:           ints.FromString(totalTime),
		recordDistance: ints.FromString(totalDistance),
	}
}

func getTimesAndDistances(lines []string) ([]string, []string) {
	timesRow := strings.Fields(lines[0])[1:]
	distanceRow := strings.Fields(lines[1])[1:]
	return timesRow, distanceRow
}
