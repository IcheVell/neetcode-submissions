func maxArea(heights []int) int {
	maxVolume := 0

	for left, right := 0, len(heights) - 1; left < right; {
		width := right - left

		var currVolume int

		if heights[left] <= heights[right] {
			currVolume = heights[left] * width
			left++
		} else {
			currVolume = heights[right] * width
			right--
		}

		if maxVolume < currVolume {
			maxVolume = currVolume
		}
	}

	return maxVolume
}