func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := math.MaxInt32

	for _, price := range prices {
		if maxProfit < price - minBuy {
			maxProfit = price - minBuy
		}

		if minBuy > price {
			minBuy = price
		}
	}

	return maxProfit
}
