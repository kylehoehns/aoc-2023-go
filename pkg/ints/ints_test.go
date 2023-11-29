package ints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumList(t *testing.T) {

	t.Run("Should total all values in a list of ints", func(t *testing.T) {
		expected := 6
		actual := Sum([]int{1, 2, 3})
		assert.Equal(t, expected, actual)
	})

}

func TestFromString(t *testing.T) {

	t.Run("Should convert string to int", func(t *testing.T) {
		expected := 2113
		actual := FromString("2113")
		assert.Equal(t, expected, actual)
	})

	t.Run("Should panic if provide a string that cannot be turned into an int", func(t *testing.T) {
		assert.Panics(t, func() { FromString("test") })
	})
}

func TestFromSliceString(t *testing.T) {

	t.Run("Should convert slice string to slice int", func(t *testing.T) {
		expected := []int{1, 2, 3}
		actual := FromStringSlice([]string{"1", "2", "3"})
		assert.Equal(t, expected, actual)
	})

}

func TestAbs(t *testing.T) {

	t.Run("Should return the absolute value of a negative int", func(t *testing.T) {
		expected := 3
		actual := Abs(-3)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return the absolute value of a positive int", func(t *testing.T) {
		expected := 3
		actual := Abs(3)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return the absolute value of zero", func(t *testing.T) {
		expected := 0
		actual := Abs(0)
		assert.Equal(t, expected, actual)
	})
}
