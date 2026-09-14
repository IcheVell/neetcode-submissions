func isPalindrome(s string) bool {
	low := strings.ToLower(s)

	runes := []rune(low)

	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		fmt.Println(string(runes[i]), string(runes[j]))
		
		for ; i < j && !isAlpha(runes[i]) && !isNum(runes[i]); {
			i++
		}

		for ; j > i && !isAlpha(runes[j]) && !isNum(runes[j]) ; {
			j--
		}

		fmt.Println(string(runes[i]), string(runes[j]))

		if i >= j {
			break
		}

		if runes[i] != runes[j] {
			fmt.Println(string(runes[i]), string(runes[j]))
			return false
		}
	}

	return true
}

func isAlpha(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return true
	}

	return false
}

func isNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}
