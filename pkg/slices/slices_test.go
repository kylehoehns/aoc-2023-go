package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow(t *testing.T) {

	t.Run("Should return a sliding window of the ints", func(t *testing.T) {
		expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
		actual := SlidingWindow([]int{1, 2, 3, 4, 5}, 3)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return an empty 2-d slice when size isn't big enough", func(t *testing.T) {
		expected := [][]int{}
		actual := SlidingWindow([]int{1, 2, 3}, 5)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return an empty 2-d slice when input is empty", func(t *testing.T) {
		expected := [][]int{}
		actual := SlidingWindow([]int{}, 1)
		assert.Equal(t, expected, actual)
	})

}
