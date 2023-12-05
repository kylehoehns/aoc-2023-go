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

type Almanac struct {
	seeds                 []int
	seedToSoil            Mappings
	soilToFertilizer      Mappings
	fertilizerToWater     Mappings
	waterToLight          Mappings
	lightToTemperature    Mappings
	temperatureToHumidity Mappings
	humidityToLocation    Mappings
}

type Mappings struct {
	name     string
	mappings []Mapping
}

func (m *Mappings) determineDestination(s int) int {
	for _, mapping := range m.mappings {
		if s >= mapping.sourceStart && s < mapping.sourceStart+mapping.rangeLength {
			return mapping.destinationStart + (s - mapping.sourceStart)
		}
	}
	return s
}

type Mapping struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func part1(name string) int {
	paragraphs := files.ReadParagraphs(name)

	seeds := parseSeedsPart1(paragraphs)
	a := parseAlmanac(paragraphs, seeds)

	return determineSmallestLocationPart1(a)
}

func part2(name string) int {
	paragraphs := files.ReadParagraphs(name)

	seeds := parseSeedsPart1(paragraphs)
	a := parseAlmanac(paragraphs, seeds)

	return determineSmallestLocationPart2(a)
}

func parseAlmanac(paragraphs [][]string, seeds []int) Almanac {
	return Almanac{
		seeds:                 seeds,
		seedToSoil:            parseMappings("seedToSoil", paragraphs[1]),
		soilToFertilizer:      parseMappings("soilToFertilizer", paragraphs[2]),
		fertilizerToWater:     parseMappings("fertilizerToWater", paragraphs[3]),
		waterToLight:          parseMappings("waterToLight", paragraphs[4]),
		lightToTemperature:    parseMappings("lightToTemperature", paragraphs[5]),
		temperatureToHumidity: parseMappings("temperatureToHumidity", paragraphs[6]),
		humidityToLocation:    parseMappings("humidityToLocation", paragraphs[7]),
	}
}

func parseMappings(name string, paragraph []string) Mappings {
	m := make([]Mapping, 0)
	for _, line := range paragraph[1:] {
		parts := strings.Split(line, " ")

		m = append(m, Mapping{
			sourceStart:      ints.FromString(parts[1]),
			destinationStart: ints.FromString(parts[0]),
			rangeLength:      ints.FromString(parts[2]),
		})
	}
	return Mappings{
		name:     name,
		mappings: m,
	}
}

func parseSeedsPart1(paragraphs [][]string) []int {
	// parse "seeds: 79 14 55 13" into an array of ints
	seeds := make([]int, 0)
	// seeds is always the first paragraph and only has a single line
	seedLine := paragraphs[0][0][7:]
	seedLineParts := strings.Split(seedLine, " ")
	for _, seedPart := range seedLineParts {
		seeds = append(seeds, ints.FromString(seedPart))
	}
	return seeds
}

func determineSmallestLocationPart1(a Almanac) int {
	locations := make([]int, 0)
	for _, seed := range a.seeds {
		locations = append(locations, determineSeedLocation(seed, a))
	}

	return ints.Min(locations)
}

func determineSmallestLocationPart2(a Almanac) int {
	locations := make([]int, 0)
	for i := 0; i < len(a.seeds); i += 2 {
		seedStart := a.seeds[i]
		seedEnd := seedStart + a.seeds[i+1]
		for j := seedStart; j < seedEnd; j++ {
			locations = append(locations, determineSeedLocation(j, a))
		}
	}

	return ints.Min(locations)
}

func determineSeedLocation(seed int, a Almanac) int {
	soil := a.seedToSoil.determineDestination(seed)
	fertilizer := a.soilToFertilizer.determineDestination(soil)
	water := a.fertilizerToWater.determineDestination(fertilizer)
	light := a.waterToLight.determineDestination(water)
	temperature := a.lightToTemperature.determineDestination(light)
	humidity := a.temperatureToHumidity.determineDestination(temperature)
	location := a.humidityToLocation.determineDestination(humidity)
	return location
}
