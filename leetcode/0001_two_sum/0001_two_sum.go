package leetcode

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	cache := map[int]int{}

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if complementIndex, exists := cache[target-curr]; exists {
			return []int{i, complementIndex}
		}

		cache[curr] = i
	}

	return []int{}
}
