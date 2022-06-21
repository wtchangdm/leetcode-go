package leetcode

func isValidPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	if isValidPalindrome(s) {
		return 1
	}

	return 2
}
