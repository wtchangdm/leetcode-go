package leetcode

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	result := nums[0]

	for left <= right {
		if nums[left] <= nums[right] && nums[left] < result {
			result = nums[left]
			break
		}

		mid := left + (right-left)/2

		if nums[mid] < result {
			result = nums[mid]
			// 6 7 1 2 3 4 5
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
