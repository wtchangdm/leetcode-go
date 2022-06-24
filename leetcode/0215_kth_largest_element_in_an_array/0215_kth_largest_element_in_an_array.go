package leetcode

func quickSelect(nums []int, targetIndex, left, right int) int {
	p, pivot := left, right

	for i := left; i < right; i++ {
		if nums[i] <= nums[pivot] {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}

	nums[p], nums[pivot] = nums[pivot], nums[p]

	if p > targetIndex {
		return quickSelect(nums, targetIndex, left, p-1)
	}

	if p < targetIndex {
		return quickSelect(nums, targetIndex, p+1, right)
	}

	return nums[targetIndex]
}

func findKthLargest(nums []int, k int) int {
	targetIndex := len(nums) - k

	return quickSelect(nums, targetIndex, 0, len(nums)-1)
}
