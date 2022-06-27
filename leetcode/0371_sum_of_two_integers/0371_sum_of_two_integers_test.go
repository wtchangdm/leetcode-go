package leetcode

import (
	"testing"
)

func TestGetSum(t *testing.T) {

	for _, v := range []struct {
		input  int
		answer int
	}{
		{input: getSum(5, 6), answer: 11},
		{input: getSum(-5, 0), answer: -5},
	} {
		if v.input != v.answer {
			t.Errorf("error: %d != %d", v.input, v.answer)
		}
	}
}
