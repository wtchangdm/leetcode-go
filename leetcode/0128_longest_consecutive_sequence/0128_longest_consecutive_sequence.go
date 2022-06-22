package leetcode

func longestConsecutive(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}

	set := map[int]struct{}{}
	longestStreak := 0

	for i := range nums {
		set[nums[i]] = struct{}{}
	}

	for num := range set {
		current := num
		if _, hasPrevious := set[current-1]; hasPrevious {
			continue
		}

		currentStreak := 1
		_, nextExists := set[current+1]

		for nextExists {
			current++
			currentStreak++
			_, nextExists = set[current+1]
		}

		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}

	return longestStreak
}
