package strings

func HasAllUniqueRunes(input string) bool {
	chars := make(map[rune]bool)
	for _, i := range input {
		if _, ok := chars[i]; ok {
			return false
		}
		chars[i] = true
	}
	return true
}
