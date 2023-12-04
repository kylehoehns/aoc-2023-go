package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

type Card struct {
	cardNumber     string
	winningNumbers []int
	numbersYouHave []int
}

func (c *Card) points() int {
	points := 0
	for _, numberYouHave := range c.numbersYouHave {
		if slices.Contains(c.winningNumbers, numberYouHave) {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
	}
	return points
}

func (c *Card) matchingNumbers() int {
	matchingNumbers := 0
	for _, numberYouHave := range c.numbersYouHave {
		if slices.Contains(c.winningNumbers, numberYouHave) {
			matchingNumbers++
		}
	}
	return matchingNumbers
}

func (c *Card) totalScratchcards(childCards []Card, cache map[string]int) int {

	if cache[c.cardNumber] != 0 {
		return cache[c.cardNumber]
	}

	points := 1
	for i := 0; i < c.matchingNumbers(); i++ {
		if len(childCards) > i {
			points += childCards[i].totalScratchcards(childCards[i+1:], cache)
		}
	}
	cache[c.cardNumber] = points
	return points
}

func part1(name string) int {
	cards := parseInput(files.ReadLines(name))
	points := 0
	for _, card := range cards {
		points += card.points()
	}
	return points
}

func part2(name string) int {
	cards := parseInput(files.ReadLines(name))
	totalScratchcards := 0
	cache := make(map[string]int)
	for i, card := range cards {
		forCard := card.totalScratchcards(cards[i+1:], cache)
		totalScratchcards += forCard
	}
	return totalScratchcards
}

func parseInput(lines []string) []Card {
	// parse "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53" into a Card struct
	cards := make([]Card, 0)
	for _, line := range lines {
		cardNumber := line[5:strings.Index(line, ":")]
		winningNumberString := line[strings.Index(line, ":")+2 : strings.Index(line, "|")-1]
		winningNumbers := make([]int, 0)
		for _, winningNumber := range strings.Split(winningNumberString, " ") {
			if winningNumber != "" {
				winningNumbers = append(winningNumbers, ints.FromString(winningNumber))
			}
		}
		numbersYouHaveString := line[strings.Index(line, "|")+2:]
		numbersYouHave := make([]int, 0)
		for _, numberYouHave := range strings.Split(numbersYouHaveString, " ") {
			if numberYouHave != "" {
				numbersYouHave = append(numbersYouHave, ints.FromString(numberYouHave))
			}
		}
		cards = append(cards, Card{cardNumber: cardNumber, winningNumbers: winningNumbers, numbersYouHave: numbersYouHave})
	}

	return cards
}
