package leetcode

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := map[byte]int{}
	countT := map[byte]int{}

	for i := range s {
		countS[s[i]]++
		countT[t[i]]++
	}

	for i := range s {
		if countS[s[i]] != countT[s[i]] {
			return false
		}
	}

	return true
}
