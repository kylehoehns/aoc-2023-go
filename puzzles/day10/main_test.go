package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1 Simple", func(t *testing.T) {
		expected := 4
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 1 Harder", func(t *testing.T) {
		expected := 8
		actual := part1("test-input-2.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 0
		actual := part2("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
