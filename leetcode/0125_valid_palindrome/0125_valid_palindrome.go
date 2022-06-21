package leetcode

import (
	"unicode"
)

func isAlphaNumeric(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		leftChar := unicode.ToLower(rune(s[left]))
		rightChar := unicode.ToLower(rune(s[right]))

		if !isAlphaNumeric(leftChar) {
			left++
			continue
		}

		if !isAlphaNumeric(rightChar) {
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
