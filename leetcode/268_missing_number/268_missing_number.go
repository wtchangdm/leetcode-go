package leetcode

func missingNumber(nums []int) int {
	expectedSum := (len(nums) * (len(nums) + 1)) / 2

	actualSum := 0

	for i := range nums {
		actualSum += nums[i]
	}

	return expectedSum - actualSum
}
