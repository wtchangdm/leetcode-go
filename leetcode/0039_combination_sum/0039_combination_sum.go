package leetcode

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var findSum func(subCandidates []int, i, total int)
	findSum = func(subCandidates []int, i, total int) {
		if total == target {
			clone := make([]int, len(subCandidates))
			copy(clone, subCandidates)
			result = append(result, clone)
			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		curr := candidates[i]
		findSum(subCandidates, i+1, total)
		subCandidates = append(subCandidates, curr)
		findSum(subCandidates, i, total+curr)
	}

	findSum([]int{}, 0, 0)

	return result
}
