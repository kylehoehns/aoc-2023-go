package slices

func SlidingWindow(slice []int, size int) [][]int {
	var result [][]int = make([][]int, 0)
	for i := 0; i <= len(slice)-size; i++ {
		window := slice[i : i+size]
		result = append(result, window)
	}
	return result
}
