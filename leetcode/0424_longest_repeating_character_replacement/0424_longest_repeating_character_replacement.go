package leetcode

func characterReplacement(s string, k int) int {
	maxLength := 0
	start := 0
	freq := map[byte]int{}
	maxFreq := 0

	for end := 0; end < len(s); end++ {
		freq[s[end]]++

		if freq[s[end]] > maxFreq {
			maxFreq = freq[s[end]]
		}

		length := end - start + 1

		if length-maxFreq > k {
			freq[s[start]]--
			start++
			length--
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
