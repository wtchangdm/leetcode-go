package leetcode

var openingParentheses = map[rune]struct{}{
	'(': {},
	'[': {},
	'{': {},
}

var closingParentheses = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValidParentheses(s string) bool {
	stack := []rune{}
	for _, r := range s {
		if _, isOpeningParenthesis := openingParentheses[r]; isOpeningParenthesis {
			stack = append(stack, r)
			continue
		}

		if _, isClosingParenthesis := closingParentheses[r]; isClosingParenthesis {

			if len(stack) == 0 {
				return false
			}

			poped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if poped != closingParentheses[r] {
				return false
			}
		}
	}

	return len(stack) == 0
}
