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

func FromStringSlice(strs []string) []int {
	vals := make([]int, 0)
	for _, str := range strs {
		vals = append(vals, FromString(str))
	}
	return vals
}

func SumStringSlice(strs []string) int {
	return Sum(FromStringSlice(strs))
}
