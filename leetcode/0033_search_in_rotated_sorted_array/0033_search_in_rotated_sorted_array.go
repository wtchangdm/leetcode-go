package leetcode

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	pivot := 0
	length := len(nums)

	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left] < nums[right] && nums[left] <= nums[pivot] {
			pivot = left
			break
		}

		mid := left + (right-left)/2

		if nums[mid] < nums[pivot] {
			pivot = mid
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	for left, right := pivot, len(nums)-1+pivot; left <= right; {
		mid := left + (right-left)/2
		midValue := nums[mid%length]

		if midValue == target {
			return mid % length
		}

		if midValue > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
