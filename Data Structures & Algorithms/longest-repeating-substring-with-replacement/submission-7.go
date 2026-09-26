func characterReplacement(s string, k int) int {
	maxLength, maxFreq := 0, 0

	m := make(map[byte]int, 0)

	for left, right := 0, 0; right < len(s); right++ {
		m[s[right]]++

		if maxFreq < m[s[right]] {
			maxFreq = m[s[right]]
		}

		for (right - left + 1) - maxFreq > k {
			m[s[left]]--
			left++
		}

		if right - left + 1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}	
