package leetcode

func getLowerCasedAlphaNumeric(r rune) rune {
	if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return r
	}

	if r >= 'A' && r <= 'Z' {
		return r + 32 // to lower case
	}

	return 0
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		leftChar := getLowerCasedAlphaNumeric(rune(s[left]))
		rightChar := getLowerCasedAlphaNumeric(rune(s[right]))

		if leftChar == 0 {
			left++
			continue
		}

		if rightChar == 0 {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}
