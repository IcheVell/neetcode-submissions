func twoSum(numbers []int, target int) []int {
	for left, right := 0, len(numbers) - 1; left < right; {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}

	return []int{}
}
