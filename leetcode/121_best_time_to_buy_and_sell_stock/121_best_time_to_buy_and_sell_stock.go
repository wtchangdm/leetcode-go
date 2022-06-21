package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	buyAt, sellAt := 0, 1

	maxProfit := 0

	for sellAt < len(prices) {
		profit := prices[sellAt] - prices[buyAt]

		if profit > 0 {
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			buyAt = sellAt
		}

		sellAt++
	}

	return maxProfit
}
