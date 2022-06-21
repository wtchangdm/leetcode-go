package leetcode

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	result := []int{}

	for i, num := range nums {
		complmentIndex, ok := m[num]

		if ok {
			result = append(result, complmentIndex, i)
			break
		}

		m[target-num] = i
	}

	return result
}
