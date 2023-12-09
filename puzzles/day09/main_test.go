package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 114
		actual := performMirageMaintenance("test-input.txt", false)
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 2
		actual := performMirageMaintenance("test-input.txt", true)
		assert.Equal(t, expected, actual)
	})

}
