package leetcode

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		if num%2 != 0 {
			count++
		}

		num /= 2
	}

	return count
}
