package leetcode

var climbingWays = map[int]int{
	0: 1,
	1: 1,
}

func climbStairsTopDown(n int) int {
	if _, ok := climbingWays[n]; !ok {
		climbingWays[n] = climbStairsTopDown(n-1) + climbStairsTopDown(n-2)
	}

	return climbingWays[n]
}

func climbStairsButtomUp(n int) int {
	dp := []int{1, 1}

	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[n]
}
