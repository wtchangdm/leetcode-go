package leetcode

func characterReplacement(s string, k int) int {
	maxLength := 0
	start := 0
	freq := map[byte]int{}
	maxFreq := 0

	for end := 0; end < len(s); end++ {
		char := s[end]
		freq[char]++

		if freq[char] > maxFreq {
			maxFreq = freq[char]
		}

		substrLength := end - start + 1
		charsToReplace := substrLength - maxFreq

		if charsToReplace > k {
			freq[s[start]]--
			start++
			substrLength--
		}

		if substrLength > maxLength {
			maxLength = substrLength
		}
	}

	return maxLength
}
