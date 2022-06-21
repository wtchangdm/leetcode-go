package leetcode

import "testing"

func TestHammingWeight(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	result := 3

	if hammingWeight(num) != result {
		t.Error("error: the answer is not", result)
	}
}
