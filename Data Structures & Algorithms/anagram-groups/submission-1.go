func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		var f [26]int
		for _, r := range str {
			f[r - 'a']++
		}

		m[f] = append(m[f], str)
	}

	res := make([][]string, 0, len(m))

	for _, val := range m {
		res = append(res, val)
	}

	return res
}
