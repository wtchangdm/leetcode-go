package leetcode

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	// nums   [1, 2, 3, 4]
	// answer  1  1  2  6
	// answer  24 12 8  6

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfix
		postfix *= nums[i]
	}

	return answer
}
