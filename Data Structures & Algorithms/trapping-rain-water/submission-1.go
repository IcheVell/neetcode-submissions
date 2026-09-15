func trap(height []int) int {
	water := 0

	for left := 0; left < len(height) - 2; {
		if height[left] == 0 {
			left++
			continue
		}

		closeWallIdx := left + 1

		for right := left + 1; right < len(height); right++ {
			if height[right] >= height[left] {
				closeWallIdx = right
				break
			} else {
				if height[closeWallIdx] < height[right] {
					closeWallIdx = right
				}
			}
		}

		if closeWallIdx != 0 {
			minHeight := min(height[left], height[closeWallIdx])

			for i := left + 1; i < closeWallIdx; i++ {
				water += minHeight - height[i]
			}

			left = closeWallIdx
		} else {
			left++
		}
	}

	return water
}


func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}

	return num2
}
