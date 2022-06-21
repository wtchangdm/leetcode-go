package leetcode

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	for _, v := range []struct {
		input  int
		result bool
	}{
		{input: -5, result: false},
		{input: 1, result: true},
		{input: 7, result: false},
		{input: 8, result: true},
	} {
		if isPowerOfTwo(v.input) != v.result {
			t.Errorf("error: result of %v != %v", v.input, v.result)
		}
	}
}
