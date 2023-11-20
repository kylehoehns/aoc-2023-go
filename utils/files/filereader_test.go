package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadLines(t *testing.T) {

	t.Run("Take a file name relative to the current file and return a string array of the lines", func(t *testing.T) {

		expected := []string{"this", "is", "a", "test", "file"}
		actual := ReadLines("./input.txt")

		assert.Equal(t, expected, actual)
	})

}

func TestRead(t *testing.T) {
	t.Run("Take a file name relative to the curent file and return a string representation of the contents", func(t *testing.T) {
		expected := "this\nis\na\ntest\nfile"
		actual := Read("./input.txt")

		assert.Equal(t, expected, actual)
	})
}
