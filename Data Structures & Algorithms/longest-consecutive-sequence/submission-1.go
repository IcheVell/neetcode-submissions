func longestConsecutive(nums []int) int {
	maxLen := 0

	m := make(map[int]struct{})

	seq := make([]int, len(m))
	
	for _, num := range nums {
		m[num] = struct{}{}
	}

	for key, _ := range m {
		_, ok := m[key - 1]
		if !ok {
			seq = append(seq, key)
		}
	}

	for _, startSeq := range seq {
		currLen := 1
		for i := startSeq + 1; ; i++ {
			if _, ok := m[i]; ok {
				currLen++
			} else {
				if currLen > maxLen {
					maxLen = currLen
				}

				break
			}
		}
	}

	return maxLen 
}
