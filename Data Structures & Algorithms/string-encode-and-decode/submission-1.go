type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		for _, r := range str {
			encodeRune := fmt.Sprintf("%x", r)
			sb.WriteString(encodeRune)
		}

		sb.WriteString(" ")
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	strs := make([]string, 0)
	
	var sb strings.Builder
	var sb2 strings.Builder

	runes := []rune(encoded)

	for i := 0; i < len(runes); i += 2 {
		if runes[i] == ' ' {
			strs = append(strs, sb.String())
			sb.Reset()
			i--
			continue
		}

		sb2.WriteRune(runes[i])
		sb2.WriteRune(runes[i + 1])

		str := sb2.String()
		sb2.Reset()
		num, _ := strconv.ParseInt(str, 16, 64)

		sb.WriteRune(rune(num))
	}

	return strs
}
