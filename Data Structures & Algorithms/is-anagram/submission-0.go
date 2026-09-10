func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[rune]int, len(s))
	m2 := make(map[rune]int, len(t))

	for _, r := range s {
		m1[r]++
	}

	for _, r := range t {
		m2[r]++
	}

	for key, _ := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}

	return true
}
