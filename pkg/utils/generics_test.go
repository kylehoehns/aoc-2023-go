package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPtr(t *testing.T) {

	t.Run("Should return a pointer to the value", func(t *testing.T) {
		expected := "test"
		actual := Ptr("test")
		assert.Equal(t, expected, *actual)
	})

}
