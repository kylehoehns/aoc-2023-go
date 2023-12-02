package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", SumOfPossibleGameIds("input.txt", 12, 13, 14))
	fmt.Println("Part 2: ", SumOfMinimumPower("input.txt"))
}

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	blue  int
	green int
}

func SumOfPossibleGameIds(name string, desiredRed int, desiredGreen int, desiredBlue int) int {
	games := createGames(name)

	gameIdSum := 0
	for _, game := range games {
		validSets := 0
		for _, set := range game.sets {
			if set.red <= desiredRed && set.green <= desiredGreen && set.blue <= desiredBlue {
				validSets++
			}
		}
		if validSets == len(game.sets) {
			gameIdSum += game.id
		}
	}

	return gameIdSum
}

func SumOfMinimumPower(name string) int {
	games := createGames(name)

	minimumPower := 0
	for _, game := range games {
		minRed, minBlue, minGreen := 0, 0, 0
		for _, set := range game.sets {
			if set.red > minRed {
				minRed = set.red
			}
			if set.blue > minBlue {
				minBlue = set.blue
			}
			if set.green > minGreen {
				minGreen = set.green
			}
		}
		minimumPower += minRed * minBlue * minGreen
	}

	return minimumPower
}

func createGames(name string) []Game {
	lines := files.ReadLines(name)

	games := make([]Game, 0)
	for _, line := range lines {
		games = append(games, convertLineToGame(line))
	}
	return games
}

func convertLineToGame(line string) Game {
	// parse something like "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green" into a game struct
	colon := strings.Index(line, ":")
	firstSpace := strings.Index(line, " ")
	id := line[firstSpace+1 : colon]

	sets := make([]Set, 0)
	for _, rawSet := range strings.Split(line[colon+1:], ";") {
		colorGroups := strings.Split(strings.TrimSpace(rawSet), ", ")
		red, blue, green := determineSetValues(colorGroups)
		sets = append(sets, Set{
			red:   red,
			blue:  blue,
			green: green,
		})
	}

	return Game{
		id:   ints.FromString(id),
		sets: sets,
	}
}

func determineSetValues(parts []string) (int, int, int) {
	red, blue, green := 0, 0, 0
	for _, part := range parts {
		colorRoll := strings.Split(part, " ")
		val := ints.FromString(colorRoll[0])
		switch colorRoll[1] {
		case "red":
			red = val
		case "blue":
			blue = val
		case "green":
			green = val
		}
	}
	return red, blue, green
}
