package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLength := 0
	cache := map[byte]int{}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, found := cache[s[j]]; found {
				i = max(cache[s[j]]+1, i)
			}

			maxLength = max(maxLength, j-i+1)
			cache[s[j]] = j
		}
	}

	return maxLength
}
