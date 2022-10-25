package leetcode

var leftParentheses = map[rune]struct{}{
	'(': {},
	'[': {},
	'{': {},
}

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if _, isLeftParenthes := leftParentheses[char]; isLeftParenthes {
			stack = append(stack, char)
			continue
		}

		if _, isRightParenthes := pairs[char]; !isRightParenthes {
			continue
		}

		if len(stack) == 0 {
			return false
		}

		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pairs[char] != curr {
			return false
		}
	}

	return len(stack) == 0
}
