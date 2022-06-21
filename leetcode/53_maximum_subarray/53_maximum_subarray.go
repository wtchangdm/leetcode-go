package leetcode

import "math"

func maxSubArray(nums []int) int {
	globalMax := math.MinInt32
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += nums[i]

		if currentSum > globalMax {
			globalMax = currentSum
		}
	}

	return globalMax
}
