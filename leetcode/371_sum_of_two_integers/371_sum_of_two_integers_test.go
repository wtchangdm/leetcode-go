package leetcode

import (
	"testing"
)

func TestGetSum(t *testing.T) {

	for _, v := range []struct {
		input  int
		result int
	}{
		{input: getSum(5, 6), result: 11},
		{input: getSum(-5, 0), result: -5},
	} {
		if v.input != v.result {
			t.Errorf("error: %d != %d", v.input, v.result)
		}
	}
}
