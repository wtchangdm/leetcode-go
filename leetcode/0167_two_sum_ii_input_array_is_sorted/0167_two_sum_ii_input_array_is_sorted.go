package leetcode

// time: O(n), space: O(1)
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{}
}

// time: O(n), space: O(n), but doesn't require the array to be sorted
func twoSumWithMap(numbers []int, target int) []int {
	cache := map[int]int{}

	for i := range numbers {
		complement := target - numbers[i]
		complementIndex, found := cache[complement]

		if found {
			return []int{complementIndex + 1, i + 1}
		}

		cache[numbers[i]] = i
	}

	return []int{}
}
