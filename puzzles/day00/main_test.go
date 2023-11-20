package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 24000
		actual := elfWithMostCalories("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 45000
		actual := sumOfCaloriesCarriedByTopThreeElves("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
