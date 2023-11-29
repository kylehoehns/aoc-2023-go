package ints

import "strconv"

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func FromString(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// FromStringSlice converts a slice of strings to a slice of ints
func FromStringSlice(input []string) []int {
	output := make([]int, 0)
	for _, str := range input {
		output = append(output, FromString(str))
	}
	return output
}
