func maxProfit(prices []int) int {
	maxProfit := 0

	for left, right := 0, 1; left + 1 < len(prices); {
		if right >= len(prices) {
			left++
			right = left + 1
			continue
		}

		currProfit := prices[right] - prices[left]

		if currProfit > maxProfit {
			maxProfit = currProfit
		}

		right++
	}

	return maxProfit
}
