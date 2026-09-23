func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	currLength := 0

	runes := []rune(s)

	used := make(map[rune]struct{}, 0)

	for left, right := 0, 0; left < len(s); {
		if right >= len(s) {
			break
		}
		
		_, has := used[runes[right]]

		if !has {
			used[runes[right]] = struct{}{}
			currLength++
		} else {
			if maxLength < currLength {
				maxLength = currLength
			}

			for runes[left] != runes[right] {
				delete(used, runes[left])
				currLength--
				left++
			}

			if runes[left] == runes[right] {
				left++
			}

			used[runes[right]] = struct{}{}
		}

		right++
	}

	if maxLength < currLength {
		maxLength = currLength
	}

	return maxLength
}
