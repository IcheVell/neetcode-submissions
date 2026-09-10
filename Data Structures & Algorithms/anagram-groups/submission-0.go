func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	maps := make([][]int, len(strs))

	for i, str := range strs {
		m := make([]int, 26, 26)	

		for _, r := range str {
			m[r - 97]++
		}

		maps[i] = m
	}

	set := make(map[int]struct{})

	for i := range strs {
		if _, ok := set[i]; ok {
			continue
		}

		gr := make([]int, 0)
		gr = append(gr, i)
		set[i] = struct{}{}

		for j := i + 1; j < len(strs); j++ {
			if _, ok := set[j]; ok {
				continue
			}

			if isAnagramm(maps[i], maps[j]) {
				gr = append(gr, j)
				set[j] = struct{}{}
			}
		}

		group := make([]string, 0, len(gr))

		for _, val := range gr {
			group = append(group, strs[val])
		}

		res = append(res, group)
	}
	
	return res
}

func isAnagramm(f1 []int, f2 []int) bool {
	if len(f1) != len(f2) {
		return false
	}

	for i := range f1 {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}
