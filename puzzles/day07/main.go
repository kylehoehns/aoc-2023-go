package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
	"sort"
)

const (
	HandTypeFiveOfAKind  = 7
	HandTypeFourOfAKind  = 6
	HandTypeFullHouse    = 5
	HandTypeThreeOfAKind = 4
	HandTypeTwoPair      = 3
	HandTypeOnePair      = 2
	HandTypeHighCard     = 1
)

var LetterMapping = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var LetterMappingPart2 = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := files.ReadLines(name)
	rounds := parseInput(lines, false)

	sort.Slice(rounds, sortRoundsWorstFirst(rounds, false))

	winnings := 0
	for i, round := range rounds {
		winnings += (i + 1) * round.bid
	}
	return winnings
}

func part2(name string) int {
	lines := files.ReadLines(name)
	rounds := parseInput(lines, true)

	sort.Slice(rounds, sortRoundsWorstFirst(rounds, true))

	winnings := 0
	for i, round := range rounds {
		winnings += (i + 1) * round.bid
	}
	return winnings
}

type LetterAndCount struct {
	letter string
	count  int
}

func isFiveOfAKind(counts []LetterAndCount) bool {
	return len(counts) == 1
}

func isFourOfAKind(counts []LetterAndCount) bool {
	return len(counts) == 2 && (counts[0].count == 4 || counts[1].count == 4)
}

func isFullHouse(counts []LetterAndCount) bool {
	return len(counts) == 2 && (counts[0].count == 3 || counts[1].count == 3)
}

func isThreeOfAKind(counts []LetterAndCount) bool {
	return len(counts) == 3 && (counts[0].count == 3 || counts[1].count == 3 || counts[2].count == 3)
}

func isTwoPair(counts []LetterAndCount) bool {
	return len(counts) == 3 &&
		(counts[0].count == 2 && counts[1].count == 2) ||
		(counts[0].count == 2 && counts[2].count == 2) ||
		(counts[1].count == 2 && counts[2].count == 2)
}

func isOnePair(counts []LetterAndCount) bool {
	return len(counts) == 4
}

func isHighCard(counts []LetterAndCount) bool {
	return len(counts) == 5
}

type CamelCardRound struct {
	hand          string
	bid           int
	handTypeScore int
}

func sortRoundsWorstFirst(rounds []CamelCardRound, part2 bool) func(i int, j int) bool {
	return func(i, j int) bool {
		if rounds[i].handTypeScore == rounds[j].handTypeScore {
			for k := 0; k < 5; k++ {
				if rounds[i].hand[k] != rounds[j].hand[k] {
					if part2 {
						return LetterMappingPart2[string(rounds[i].hand[k])] < LetterMappingPart2[string(rounds[j].hand[k])]
					} else {
						return LetterMapping[string(rounds[i].hand[k])] < LetterMapping[string(rounds[j].hand[k])]
					}
				}
			}
			return false
		} else {
			return rounds[i].handTypeScore < rounds[j].handTypeScore
		}
	}
}

func parseInput(lines []string, part2 bool) []CamelCardRound {
	rounds := make([]CamelCardRound, 0)
	for _, line := range lines {
		letters := line[:5]
		bid := ints.FromString(line[6:])

		handTypeScore := determineHandTypeScore(letters, part2)

		rounds = append(rounds, CamelCardRound{letters, bid, handTypeScore})
	}
	return rounds
}

func determineHandTypeScore(letters string, part2 bool) int {

	counts := make(map[string]int)
	for _, c := range letters {
		counts[string(c)]++
	}
	lettersAndCounts := make([]LetterAndCount, 0)
	for k, v := range counts {
		lettersAndCounts = append(lettersAndCounts, LetterAndCount{k, v})
	}

	if part2 {
		lettersAndCounts = alterCountsForJokers(lettersAndCounts)
	}

	if isFiveOfAKind(lettersAndCounts) {
		return HandTypeFiveOfAKind
	} else if isFourOfAKind(lettersAndCounts) {
		return HandTypeFourOfAKind
	} else if isFullHouse(lettersAndCounts) {
		return HandTypeFullHouse
	} else if isThreeOfAKind(lettersAndCounts) {
		return HandTypeThreeOfAKind
	} else if isTwoPair(lettersAndCounts) {
		return HandTypeTwoPair
	} else if isOnePair(lettersAndCounts) {
		return HandTypeOnePair
	} else if isHighCard(lettersAndCounts) {
		return HandTypeHighCard
	} else {
		panic("Unknown hand type " + letters)
	}
}

func alterCountsForJokers(lettersAndCounts []LetterAndCount) []LetterAndCount {
	shouldRemoveJ := false
	letterToUpdate := ""
	amountToUpdate := 0
	for _, lc := range lettersAndCounts {
		if lc.letter == "J" {
			// if we have a Joker, determine the largest letter and add the count to it
			amountToUpdate = lc.count
			if largestLetterCount, shouldReplace := getLargestLetterCount(lettersAndCounts); shouldReplace {
				shouldRemoveJ = true
				letterToUpdate = largestLetterCount.letter
			}
			break
		}
	}
	if shouldRemoveJ {
		// update the count for the biggest letter
		for i, letterCount := range lettersAndCounts {
			if letterCount.letter == letterToUpdate {
				lettersAndCounts[i].count += amountToUpdate
			}
		}
		// remove jokers
		for i, letterCount := range lettersAndCounts {
			if letterCount.letter == "J" {
				lettersAndCounts = append(lettersAndCounts[:i], lettersAndCounts[i+1:]...)
			}
		}
	}
	return lettersAndCounts
}

func getLargestLetterCount(lettersAndCounts []LetterAndCount) (LetterAndCount, bool) {
	largestLetterCount := LetterAndCount{}
	for _, lc := range lettersAndCounts {
		if lc.letter == "J" {
			continue
		}
		if largestLetterCount.letter == "" {
			largestLetterCount = lc
		} else {
			if lc.count > largestLetterCount.count {
				largestLetterCount = lc
			}
		}
	}
	return largestLetterCount, largestLetterCount.letter != ""
}
