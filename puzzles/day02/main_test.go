package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 8
		actual := SumOfPossibleGameIds("test-input.txt", 12, 13, 14)
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 2286
		actual := SumOfMinimumPower("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
