func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	buckets := make([][]int, len(nums) + 1)

	for key, value := range m {
		buckets[value] = append(buckets[value], key)
	}

	for i := len(buckets) - 1; i >= 0 && k > 0; i-- {
		for j := range buckets[i] {
			res = append(res, buckets[i][j])
			k--
		}
	}

	return res
}
