package leetcode

import "testing"

func TestHammingWeight(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	answer := 3

	if hammingWeight(num) != answer {
		t.Error("error: the result is not", answer)
	}
}
