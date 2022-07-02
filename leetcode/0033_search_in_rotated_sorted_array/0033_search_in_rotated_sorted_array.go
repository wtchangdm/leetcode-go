package leetcode

import "fmt"

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
		if nums[left] < nums[right] {
			break
		}

		pivot = (left + right) / 2

		if nums[pivot] > nums[pivot+1] {
			pivot = pivot + 1
			break
		}

		if nums[pivot] > nums[right] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	fmt.Println(nums, nums[pivot])

	for left, right := pivot, len(nums)-1+pivot; left <= right; {
		mid := (left + right) / 2
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
