package strings

func HasAllUniqueRunes(input string) bool {
	chars := make(map[rune]bool)
	allUnique := true
	for _, i := range input {
		_, ok := chars[i]
		if ok {
			allUnique = false
		}
		chars[i] = true
	}
	return allUnique
}
