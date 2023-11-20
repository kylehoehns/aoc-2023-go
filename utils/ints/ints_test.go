package ints

import "testing"

func TestSumList(t *testing.T) {

	t.Run("Should total all values in a list of ints", func(t *testing.T) {
		expected := 6
		actual := Sum([]int{1, 2, 3})
		if expected != actual {
			t.Fail()
			t.Logf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestFromString(t *testing.T) {

	t.Run("Should convert string to int", func(t *testing.T) {
		expected := 2113
		actual := FromString("2113")
		if expected != actual {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

	t.Run("Should panic if provide a string that cannot be turned into an int", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected code to panic")
			}
		}()

		FromString("test")
	})
}
