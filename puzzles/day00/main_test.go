package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {

		expenses := []int{1721, 979, 366, 299, 675, 1456}
		expected := 514579
		actual := part1(expenses)
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {

		expenses := []int{1721, 979, 366, 299, 675, 1456}
		expected := 241861950
		actual := part2(expenses)
		assert.Equal(t, expected, actual)
	})

}
