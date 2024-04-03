func maxProfit(prices []int) int {
    maxProfit := 0
	newTemp := prices[0]
	for i := 1; i < len(prices); i++ {
		if newTemp < prices[i] {
			if (prices[i] - newTemp) > maxProfit {
				maxProfit = prices[i] - newTemp
			}
		} else {
			newTemp = prices[i]
		}
	}
	return maxProfit
}
