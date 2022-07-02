package leetcode

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	pivot := 0

	for left <= right {
		// seems normal, nums might be length N with N times rotations
		if nums[right] > nums[left] {
			return nums[pivot]
		}

		pivot = (left + right) / 2

		if nums[pivot] > nums[pivot+1] {
			pivot = pivot + 1
			break
		}

		// pivot is in the right side of pivot
		if nums[pivot] > nums[right] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return nums[pivot]
}
