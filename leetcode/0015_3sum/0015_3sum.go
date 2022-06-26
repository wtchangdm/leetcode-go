package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1

		if i > 0 && nums[i] == nums[i-1] {
			// skip duplicates like [-3, -3, ...]
			continue
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return result
}
