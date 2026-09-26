func checkInclusion(s1 string, s2 string) bool {
	m := make(map[byte]int, 0)

	for i := range s1 {
		m[s1[i]]++
	}

	for left, right := 0, 0; right < len(s2); right++ {
		val, ok := m[s2[right]]

		if !ok {
			for left < right {
				m[s2[left]]++
				left++
			}

			left = right + 1
			continue
		}

		if val > 0 {
			m[s2[right]]--
		} else {
			for left <= right {
				if m[s2[right]] > 0 {
					m[s2[right]]--
					break
				}

				m[s2[left]]++
				left++
			}
		}

		if right - left + 1 == len(s1) {
			return true
		}
	}

	return false
}
