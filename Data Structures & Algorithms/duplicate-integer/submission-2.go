func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		_, ok := set[num]
		if ok {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}
