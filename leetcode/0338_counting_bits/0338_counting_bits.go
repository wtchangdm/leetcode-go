package leetcode

func countBits(n int) []int {
	dp := []int{0}
	offset := 1

	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}

		dp = append(dp, 1+dp[i-offset])
	}

	return dp
}
