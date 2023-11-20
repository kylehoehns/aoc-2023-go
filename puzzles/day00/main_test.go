package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {

		expenses := []int{1721, 979, 366, 299, 675, 1456}
		actual := part1(expenses)
		expected := 514579
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {

		expenses := []int{1721, 979, 366, 299, 675, 1456}
		actual := part2(expenses)
		expected := 241861950
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}
